package main

import (
	"naapa-go/server"
	"naapa-go/server/handlers"
	"naapa-go/server/router"
)

func main() {
	s := server.Init()
	healthCheckRouter := router.New()
	healthCheckRouter.RegisterRoute("GET /status", handlers.HealthCheck)
	s.AddRouter(healthCheckRouter)

	s.Listen()
}
