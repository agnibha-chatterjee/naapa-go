package main

import (
	"naapa-go/server"
	"naapa-go/server/handlers"
	"naapa-go/server/router"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	s := server.Init()
	h := router.New("health-check")
	h.RegisterRoute("GET /status", handlers.HealthCheck)

	t := router.New("task-router")
	t.RegisterRoute("POST /execute-task", handlers.HandleExecuteTask)

	s.AddRouter(h)
	s.AddRouter(t)
	s.Listen()
}
