package database

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/database/model"
)

var (
	instance *gorm.DB
	once     sync.Once
)

// DatabaseConfig define la configuración necesaria para la conexión.
type DatabaseConfig struct {
	Host             string
	User             string
	Password         string
	Name             string
	Port             int
	EnableMigrations bool
	LogLevel         logger.LogLevel
}

// Connect inicializa y devuelve la conexión a la base de datos.
func Connect(cfg DatabaseConfig) (*gorm.DB, error) {
	var err error

	once.Do(func() {
		// DSN
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
			cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port,
		)

		// Logger
		newLogger := getLogger(cfg.LogLevel)

		// Conexión
		instance, err = gorm.Open(postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true,
		}), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: false,
			},
			Logger: newLogger,
		})

		if err != nil {
			err = fmt.Errorf("failed to connect to the database: %w", err)
			return
		}

		// Ejecutar migraciones si está habilitado
		if cfg.EnableMigrations {
			if err = runMigrations(instance); err != nil {
				return
			}
		}

		log.Println("Database connection initialized successfully.")
	})

	return instance, err
}

// runMigrations ejecuta las migraciones de la base de datos.
func runMigrations(db *gorm.DB) error {
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		return fmt.Errorf("failed to migrate the database: %w", err)
	}
	log.Println("Migrations executed successfully.")
	return nil
}

// getLogger crea un logger de GORM.
func getLogger(logLevel logger.LogLevel) logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logLevel,
			Colorful:      true,
		},
	)
}

// CloseDB cierra la conexión a la base de datos.
func CloseDB() {
	if instance == nil {
		log.Println("No active database connection to close.")
		return
	}

	sqlDB, err := instance.DB()
	if err != nil {
		log.Println("Error obtaining database connection:", err)
		return
	}

	if err = sqlDB.Close(); err != nil {
		log.Println("Error closing database connection:", err)
	} else {
		log.Println("Database connection closed successfully.")
	}
}
