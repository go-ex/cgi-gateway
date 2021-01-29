package app

import "gateway/constants"

type app struct {
}

func NewController() {
	controller := &app{}

	constants.Router.GET("/app/add", controller.add)
	constants.Router.GET("/app/get", controller.get)
	constants.Router.GET("/app", controller.index)
}
