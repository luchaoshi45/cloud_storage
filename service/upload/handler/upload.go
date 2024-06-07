package handler

import (
	"cloud_storage/configurator"
	"cloud_storage/db/mysql"
	"cloud_storage/file"
	"cloud_storage/rabbitmq"
	"crypto/sha1"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func ShowUploadPage(c *gin.Context) {
	//c.Redirect(http.StatusFound, "/static/view/index.html")
	c.File("static/view/index.html")
}

func ReceiveFile(c *gin.Context) {
	errCode := 0
	defer func() {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		if errCode < 0 {
			c.JSON(http.StatusOK, gin.H{
				"code": errCode,
				"msg":  "上传失败",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": errCode,
				"msg":  "上传成功",
			})
		}
	}()

	// 秒传
	if FastUpload(c) {
		return
	}

	// 接收文件
	f, head, err := c.Request.FormFile("file")
	if err != nil {
		log.Println("r.FormFile(\"file\") ", err.Error())
		errCode = -1
		return
	}
	defer f.Close()

	// 本地创建文件
	newFile, err := os.Create("tmp/" + head.Filename)
	if err != nil {
		log.Println("os.Create(\"/tmp/\" + head.Filename) ", err.Error())
		errCode = -2
		return
	}

	// copy 文件
	fileSize, err := io.Copy(newFile, f)
	if err != nil {
		log.Println("io.Copy(newFile, file) ", err.Error())
		errCode = -3
		return
	}

	// 计算 fileSha1
	newFile.Seek(0, 0)
	fileSha1 := file.FileSha1(newFile)

	// 计算 Oss
	newFile.Seek(0, 0)
	ossPath := "oss/" + fileSha1
	//err = oss.Bucket().PutObject(ossPath, newFile)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	w.Write([]byte("Upload failed!"))
	//	return
	//}

	data := rabbitmq.TransferData{
		FileHash:     fileSha1,
		CurLocation:  "tmp/" + head.Filename,
		DestLocation: ossPath}
	pubData, _ := json.Marshal(data)
	cfg := configurator.GetRabbitMQConfig()
	pubSuc := rabbitmq.Publish(
		cfg.GetAttr("Exchange"),
		cfg.GetAttr("RoutingKey"),
		pubData)
	if !pubSuc {
		log.Println("rabbitmq.Publish ", err.Error())
		errCode = -4
		return
	}

	// 更新数据库
	tFile := mysql.NewFile()
	tFile.SetAttrs(map[string]interface{}{
		"UploadAt": time.Now().Format("2006-01-02 15:04:05"),
		"Name":     head.Filename,
		"Dir":      "tmp/",
		"Size":     fileSize,
		"Sha1":     fileSha1,
	})
	success := tFile.Insert()

	if success == false {
		errCode = -6
		return
	}
	// 上传UserFile
	userFile := mysql.NewUserFile()
	username := c.PostForm("username")
	success = userFile.Insert(username, fileSha1, head.Filename, fileSize)
	if !success {
		errCode = -6
	} else {
		errCode = 0
	}
}

// FastUpload 秒传接口
func FastUpload(c *gin.Context) bool {
	username := c.PostForm("username")

	// 1.2 filehash
	// 接收文件
	tf, head, err := c.Request.FormFile("file")
	if err != nil {
		log.Println("r.FormFile(\"file\") ", err.Error())
		return false
	}
	defer tf.Close()

	// 创建SHA-1哈希对象
	h := sha1.New()
	// 将文件内容传入哈希对象
	filesize, err := io.Copy(h, tf)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Io copy err",
		})
		return false
	}
	// 计算哈希值
	hashInBytes := h.Sum(nil)
	// 将哈希值转换成16进制字符串
	filehash := hex.EncodeToString(hashInBytes)

	// 1.3 filename
	filename := head.Filename

	// 2. 从文件表中查询相同hash的文件记录
	fileMeta, err := mysql.NewFile().Query(filehash)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Query err",
		})
		return false
	}

	// 3. 查不到记录则返回秒传失败
	if fileMeta == nil {
		//resp := handler.RespMsg{
		//	Code: -1,
		//	Msg:  "秒传失败，请访问普通上传接口",
		//}
		//w.Write(resp.JSONBytes())
		return false
	}

	// 4. 上传过则将文件信息写入用户文件表， 返回成功
	userFile := mysql.NewUserFile()
	suc := userFile.Insert(username, filehash, filename, int64(filesize))
	if suc {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "FastUpload success",
		})
		return true
	}
	//resp := handler.RespMsg{
	//	Code: -2,
	//	Msg:  "秒传失败，请稍后重试",
	//}
	//w.Write(resp.JSONBytes())
	return false
}
