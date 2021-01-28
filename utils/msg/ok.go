package msg

import "github.com/gin-gonic/gin"

func Ok() gin.H {
	return gin.H{
		"code": 0,
		"msg":  "ok",
	}
}

func Success(data interface{}) gin.H {
	return gin.H{
		"code": 0,
		"msg":  "ok",
		"data": data,
	}
}
