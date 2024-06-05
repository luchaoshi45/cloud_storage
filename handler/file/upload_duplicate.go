package file

import (
	"github.com/gin-gonic/gin"
	"log"
)

func (f *File) UploadDuplicate(c *gin.Context) {
	_, err := c.Writer.WriteString("uploadDuplicate")
	if err != nil {
		log.Println("c.Writer.WriteString(\"uploadDuplicate\") ", err.Error())
	}
}
