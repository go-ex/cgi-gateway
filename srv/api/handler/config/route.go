package config

import "gateway/constants"

type config struct {
}

func NewController() {
	controller := &config{}

	constants.Router.GET("/config", controller.index)
}
