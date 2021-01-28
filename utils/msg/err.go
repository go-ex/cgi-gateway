package msg

import "github.com/gin-gonic/gin"

func ErrorMsg(code int, msg string) gin.H {
	return gin.H{
		"code": code,
		"msg":  msg,
	}
}
