package app

import (
	"github.com/bolanebyla/Go-ToDo/internal/config"
	"github.com/bolanebyla/Go-ToDo/internal/database"
	"github.com/bolanebyla/Go-ToDo/internal/services"
	"github.com/bolanebyla/Go-ToDo/internal/transport"
)

func Run(conf *config.Config) {

	taskRepo := database.CreateTaskRepo()

	taskService := services.CreateTaskService(taskRepo)

	taskHandler := transport.CreateTaskHandler(taskService)
	router := transport.CreateRouter(taskHandler)

	apiServer := transport.CreateApiServer(router, conf.Host, conf.Port)

	apiServer.StartApiServer()
}
