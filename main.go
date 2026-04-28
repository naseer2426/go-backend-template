package main

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/naseer2426/go-backend-template/internal/api"
	"github.com/naseer2426/go-backend-template/internal/config"
	"github.com/naseer2426/go-backend-template/internal/db"
)

func main() {
	cfg := config.MustLoad()
	router := initRouter()

	if cfg.Database.URL != "" {
		if err := db.Init(cfg.Database.URL); err != nil {
			log.Fatalf("database init failed: %v", err)
		}
		if err := db.RunMigrations(db.GetDB()); err != nil {
			log.Fatalf("database migrations failed: %v", err)
		}
	}

	router.GET("/", api.HealthCheck)

	addr := fmt.Sprintf(":%s", cfg.Server.Port)
	if err := router.Run(addr); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func initRouter() *gin.Engine {
	router := gin.Default()
	// Allow CORS for all origins
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Accept", "X-Requested-With"},
		ExposeHeaders:   []string{"Content-Length"},
	}))

	return router
}
