package api

import (
	"gateway/constants"
	"log"
)

type Server struct {
}

func (s *Server) Start() {
	log.Println("api at 0.0.0.0/admin")
}

func (s *Server) Run() {

}

func NewHttpServer() constants.Server {
	return &Server{}
}
