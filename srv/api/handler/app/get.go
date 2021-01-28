package app

import (
	"gateway/constants"
	"gateway/utils/msg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (receiver *app) get(c *gin.Context) {
	appStrID := c.Query("app_id")

	if appStrID == "" {
		c.JSON(http.StatusOK, msg.ErrorMsg(1, "无app_id参数"))

		return
	}

	appID, _ := strconv.Atoi(appStrID)
	got := constants.GetAppAgent(appID)

	c.JSON(http.StatusOK, msg.Success(got))
}
