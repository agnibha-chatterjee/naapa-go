package server

import (
	"naapa-go/server/router"
	"net/http"
)

type Server struct {
	router []*router.Router
	port   string
}

func Init() *Server {
	newServer := &Server{port: ":5173"}
	return newServer
}

func (s *Server) SetPort(port string) {
	s.port = port
}

func (s *Server) AddRouter(router *router.Router) {
	s.router = append(s.router, router)
}

func (s *Server) Listen() {
	if s.router == nil {
		panic("The server needs to have at least 1 router!")
	}

	mux := http.NewServeMux()

	for _, router := range s.router {
		for _, route := range router.Routes {
			mux.HandleFunc(route.Path, http.HandlerFunc(route.Handler))
		}
	}

	http.ListenAndServe(s.port, mux)
}
