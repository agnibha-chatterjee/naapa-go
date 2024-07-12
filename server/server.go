package server

import (
	"naapa-go/server/router"
	"net/http"
)

type Server struct {
	routers []*router.Router
	port    string
}

func Init() *Server {
	newServer := &Server{port: ":5173"}
	return newServer
}

func (s *Server) SetPort(port string) {
	s.port = port
}

func (s *Server) AddRouter(router *router.Router) {
	s.routers = append(s.routers, router)
}

func (s *Server) Listen() {
	if len(s.routers) < 1 {
		panic("No routers were attached to the server!")
	}

	mux := http.NewServeMux()

	for _, router := range s.routers {

		if len(router.Routes) < 1 {
			panic("No routes were found in " + router.Name + " router")
		}

		for _, route := range router.Routes {
			mux.HandleFunc(route.Path, http.HandlerFunc(route.Handler))
		}
	}

	http.ListenAndServe(s.port, mux)
}
