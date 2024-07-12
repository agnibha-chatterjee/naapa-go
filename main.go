package main

import (
	"naapa-go/server"
	"naapa-go/server/handlers"
)

func main() {
	s := server.Init()
	s.RegisterRoute("GET /", handlers.HealthCheck)
	s.Listen()
}
