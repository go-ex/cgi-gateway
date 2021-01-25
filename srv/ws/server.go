package ws

import (
	"gateway/constants"
	"net/http"
)

type Server struct {
	hub *Hub
}

func (s *Server) Start() {
	s.hub = newHub()
	go s.hub.run()
}

func (s *Server) Run() {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(s.hub, w, r)
	})
}

func NewWebsocketServer() constants.Server {
	return &Server{}
}
