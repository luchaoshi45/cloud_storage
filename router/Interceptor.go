package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HttpInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.FormValue("token")

		// 验证登录token是否有效
		if !IsTokenValid(token) {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"msg":  "token无效",
				"code": -1,
			})
			return
		}

		c.Next()
	}

}

// IsTokenValid : token是否有效
func IsTokenValid(token string) bool {
	if len(token) != 40 {
		return false
	}
	// TODO: 判断token的时效性，是否过期
	// TODO: 从数据库表tbl_user_token查询username对应的token信息
	// TODO: 对比两个token是否一致
	return true
}
