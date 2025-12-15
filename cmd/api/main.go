package main

import (
	"task-manager/internal/config"
	"task-manager/internal/http/handlers/auth"
	"task-manager/internal/http/handlers/task"
	"task-manager/internal/http/middleware"
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

	userRepo := repository.NewUserRepository(db)
	authService := services.NewAuthService(userRepo, []byte(cfg.JWTSecret))
	authHandler := auth.NewHandler(authService)

	taskRepo := repository.NewTaskRepository(db)
	taskService := services.NewTaskService(taskRepo)
	taskHandler := task.NewHandler(taskService)

	r := gin.Default()

	auth := r.Group("/auth")
	{
		auth.GET("/users", authHandler.GetAll)
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	tasks := r.Group("/tasks")
	tasks.Use(middleware.AuthMiddleware([]byte(cfg.JWTSecret)))
	{
		tasks.GET("/", taskHandler.GetAll)
		tasks.POST("/", taskHandler.Create)
		tasks.PATCH("/:id", taskHandler.Update)
		tasks.DELETE("/:id", taskHandler.Delete)
	}

	r.Run()
}
