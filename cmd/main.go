package main

import (
	"github.com/asm2212/go-gin-starter/config"
	"github.com/asm2212/go-gin-starter/internal/api"
	"github.com/asm2212/go-gin-starter/internal/db"
	"github.com/asm2212/go-gin-starter/internal/middleware"
	"github.com/asm2212/go-gin-starter/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	dbPool := db.NewPostgresDB(cfg)
	defer dbPool.Close()

	userRepo := db.NewUserRepository(dbPool)
	userService := service.NewUserService(userRepo)
	userHandler := api.NewUserHandler(userService)

	r := gin.Default()

	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)

	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/profile", userHandler.Profile)

	r.Run(":" + cfg.AppPort)
}
