package api

import (
	"gateway/constants"
	"gateway/srv/api/handler/app"
	"gateway/srv/api/handler/broadcast"
	"gateway/srv/api/handler/config"
	"gateway/srv/api/handler/ws"
)

type Server struct {
}

func (s *Server) Start() {
	app.NewController()
	config.NewController()
	broadcast.NewController()
	ws.NewController()
}

func (s *Server) Run() {

}

func NewHttpServer() constants.Server {
	return &Server{}
}
