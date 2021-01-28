package app

import (
	"gateway/constants"
	"gateway/utils/msg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (receiver *app) add(c *gin.Context) {
	appStrID := c.Query("app_id")
	agent := c.Query("agent")

	if appStrID == "" {
		c.JSON(http.StatusOK, msg.ErrorMsg(1, "无app_id参数"))

		return
	}

	if agent == "" {
		c.JSON(http.StatusOK, msg.ErrorMsg(1, "无agent参数"))
		return
	}

	appID, _ := strconv.Atoi(appStrID)

	constants.AddAppAgent(appID, agent)

	c.JSON(http.StatusOK, msg.Ok())
}
