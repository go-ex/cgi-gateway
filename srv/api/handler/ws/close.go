package ws

import (
	"gateway/utils/hub"
	"gateway/utils/msg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (receiver *ws) close(c *gin.Context) {
	con := c.PostForm("fd")
	fd := hub.ConnectId(con)
	hub.GetHub().Close(fd)

	c.JSON(http.StatusOK, msg.Ok())
}
