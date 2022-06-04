package util

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	log    *Logger
	router *mux.Router
}

func NewServer() *Server {
	return &Server{
		log: NewLogger("Server"),
	}
}

func (s *Server) Start(port string, handler http.Handler) {
	s.log.Info("server is listening on %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), handler))
}
