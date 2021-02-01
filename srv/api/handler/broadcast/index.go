package broadcast

import (
	"gateway/utils/hub"
	"gateway/utils/msg"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 广播所有
func (receiver *broadcast) index(c *gin.Context) {
	data := c.PostForm("msg")

	if data != "" {
		hub.GetHub().SendAll([]byte(data))
	} else {
		log.Println(data)
	}

	c.JSON(http.StatusOK, msg.Ok)
}
