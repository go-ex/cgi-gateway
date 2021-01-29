package config

import (
	"gateway/constants"
	"gateway/utils/msg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (receiver *config) index(c *gin.Context) {
	c.JSON(http.StatusOK, msg.Success(constants.Config))
}
