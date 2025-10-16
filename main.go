package main

import (
	"go-gin-app/config"
	"go-gin-app/db"
	"go-gin-app/handlers"
	"go-gin-app/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	if err := db.InitDB(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	r := gin.Default()

	api := r.Group("/api/v1")
	{
		api.POST("/register", handlers.Register)
		api.POST("/login", handlers.Login)

		authRequired := api.Group("/")
		authRequired.Use(middleware.AuthMiddleware())
		{
			authRequired.GET("/profile", handlers.GetProfile)
			authRequired.POST("/distillation-data", handlers.UploadDistillationData)
		}
	}

	r.Run()
}
