package app

import (
	"gateway/constants"
	"gateway/utils/msg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (receiver *app) index(c *gin.Context) {
	c.JSON(http.StatusOK, msg.Success(constants.GetAll()))
}
