package router

import "net/http"

type Handler func(w http.ResponseWriter, r *http.Request)

type Router struct {
	Routes []*Route
}

type Route struct {
	Path    string
	Handler Handler
}

func New() *Router {
	return &Router{}
}

func (r *Router) RegisterRoute(path string, handler Handler) {
	route := &Route{Path: path, Handler: handler}
	r.Routes = append(r.Routes, route)
}
