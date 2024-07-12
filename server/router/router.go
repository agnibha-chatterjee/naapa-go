package router

import "net/http"

type Handler func(w http.ResponseWriter, r *http.Request)

type Router struct {
	Name   string
	Routes []*Route
}

type Route struct {
	Path    string
	Handler Handler
}

func New(name string) *Router {
	return &Router{Name: name}
}

func (r *Router) RegisterRoute(path string, handler Handler) {
	route := &Route{Path: path, Handler: handler}
	r.Routes = append(r.Routes, route)
}
