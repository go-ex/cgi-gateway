package ws

import "gateway/constants"

type ws struct {
}

func NewController() {
	controller := &ws{}

	constants.Router.POST("/ws/close", controller.close)
	constants.Router.POST("/ws/send", controller.send)
}
