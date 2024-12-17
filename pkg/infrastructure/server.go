package infrastructure

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gorm.io/gorm/logger"

	"github.com/javiertelioz/clean-architecture-go/config"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/database"
)

func RunApplication() {

	dbConfig, _ := config.GetConfig[config.DatabaseConfig]("Database")
	db, err := database.Connect(
		database.DatabaseConfig{
			Host:             dbConfig.Host,
			User:             dbConfig.User,
			Password:         dbConfig.Password,
			Name:             dbConfig.Name,
			Port:             dbConfig.Port,
			EnableMigrations: true,
			LogLevel:         logger.Info,
		})

	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer database.CloseDB()

	dependencies := Bootstrap(db)

	serverConfig, _ := config.GetConfig[config.ServerConfig]("Server")
	addr := fmt.Sprintf("%s:%s", serverConfig.Host, serverConfig.Port)

	server := &http.Server{
		Addr:    addr,
		Handler: dependencies.Router,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("🚀 Starting application on: http://%s/\n", addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v\n", addr, err)
		}
	}()

	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited properly")
}
