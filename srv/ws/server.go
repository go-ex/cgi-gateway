package ws

import (
	"gateway/constants"
	"github.com/gin-gonic/gin"
)

type Server struct {
	hub *Hub
}

func (s *Server) Start() {
	s.hub = newHub()
	go s.hub.run()
}

func (s *Server) Run() {
	constants.HttpServer.GET("/ws", func(c *gin.Context) {
		serveWs(s.hub, c.Writer, c.Request)
	})
}

func NewWebsocketServer() constants.Server {
	return &Server{}
}
