package api

import (
	"gateway/constants"
)

type Server struct {
}

func (s *Server) Start() {

}

func (s *Server) Run() {

}

func NewHttpServer() constants.Server {
	return &Server{}
}
