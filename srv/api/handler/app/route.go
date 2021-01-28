package app

import "gateway/constants"

type app struct {
}

func init() {
	app := &app{}

	constants.HttpServer.GET("/app/add", app.add)
	constants.HttpServer.GET("/app/get", app.get)
}
