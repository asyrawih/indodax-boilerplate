package server

import (
	"log"

	"bitbucket.org/btcid/startrack/internal/config"
)

type Server struct {
	config *config.Config
}

func MustNewServer(ctx *config.Config) *Server {
	return &Server{
		config: ctx,
	}
}

func (s *Server) Start() {

}

func (s *Server) Stop() {
	log.Println("Server Has Stop")
}
