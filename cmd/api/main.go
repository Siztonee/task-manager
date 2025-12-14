package main

import (
	"task-manager/internal/config"
	"task-manager/internal/http/handlers/task"
	"task-manager/internal/repository"
	"task-manager/internal/services"
	"task-manager/internal/storage/postgres"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	cfg := config.Load()

	db := postgres.InitDB(cfg)

	taskRepo := repository.NewTaskRepository(db)
	taskService := services.NewTaskService(taskRepo)
	taskHandler := task.NewHandler(taskService)

	r := gin.Default()
	tasks := r.Group("/tasks")
	{
		tasks.GET("/", taskHandler.GetAll)
		tasks.POST("/", taskHandler.Create)
	}

	r.Run()
}
