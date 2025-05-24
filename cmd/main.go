package main

import (
	"github.com/asm2212/go-gin-starter/config"
	"github.com/asm2212/go-gin-starter/internal/api"
	"github.com/asm2212/go-gin-starter/internal/db"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()
	dbPool := db.NewPostgresDB(cfg)
	defer dbPool.Close()

	r := gin.Default()
	handler := api.NewExampleHandler()

	r.GET("/hello", handler.HelloWorld)

	r.Run(":" + cfg.AppPort)
}
