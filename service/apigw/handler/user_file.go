package handler

import (
	"cloud_storage/service/account/proto"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// UserFiles : 查询批量的文件元信息
func UserFiles(c *gin.Context) {
	limitCnt, _ := strconv.Atoi(c.Request.FormValue("limit"))
	username := c.Request.FormValue("username")

	rpcResp, err := userCli.UserFiles(context.TODO(), &proto.ReqUserFile{
		Username: username,
		Limit:    int32(limitCnt),
	})

	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	if len(rpcResp.FileData) <= 0 {
		rpcResp.FileData = []byte("[]")
	}
	c.Data(http.StatusOK, "application/json", rpcResp.FileData)
}
