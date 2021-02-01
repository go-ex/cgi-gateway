package ws

import (
	"gateway/utils/hub"
	"gateway/utils/msg"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (receiver *ws) send(c *gin.Context) {
	con := c.PostForm("fd")
	data := c.PostForm("msg")

	if con == "" {
		log.Println("无连接fd")
		return
	}

	if data == "" {
		log.Println("无通知信息")
		return
	}

	hub.GetHub().SendAll([]byte(data))

	c.JSON(http.StatusOK, msg.Ok)
}
