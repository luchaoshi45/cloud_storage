package file

import (
	"github.com/gin-gonic/gin"
	"log"
)

func (f *File) UploadSuccess(c *gin.Context) {
	// 向客户端写入响应
	_, err := c.Writer.WriteString("uploadSuccess")
	if err != nil {
		log.Println("c.Writer.WriteString(\"uploadSuccess\") ", err.Error())
	}
}
