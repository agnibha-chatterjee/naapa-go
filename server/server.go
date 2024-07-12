package server

import (
	"net/http"
)

type Server struct {
	router *http.ServeMux
	port   string
}

type Handler func(w http.ResponseWriter, r *http.Request)

func Init() *Server {
	router := http.NewServeMux()
	newServer := &Server{router: router, port: ":5173"}
	return newServer
}

func (s *Server) SetPort(port string) {
	s.port = port
}

func (s *Server) Listen() {
	http.ListenAndServe(s.port, s.router)
}

func (s *Server) RegisterRoute(path string, handler Handler) {
	s.router.Handle(path, http.HandlerFunc(handler))
}
