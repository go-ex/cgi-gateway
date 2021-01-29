package broadcast

import (
	"gateway/utils/hub"
	"github.com/gin-gonic/gin"
	"log"
)

// 广播所有
func (receiver *broadcast) index(c *gin.Context) {
	msg := c.PostForm("msg")

	if msg != "" {
		hub.GetHub().SendAll([]byte(msg))
	} else {
		log.Println(msg)
	}
}
