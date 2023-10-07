package database

import (
	"fmt"
	"github.com/javiertelioz/clean-architecture-go/config"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/database/model"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	instance *gorm.DB
	once     sync.Once
)

func Connect() *gorm.DB {
	once.Do(func() {
		database, _ := config.GetConfig[config.DatabaseConfig]("Database")
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
			database.Host,
			database.User,
			database.Password,
			database.Name,
			database.Port,
		)

		newLogger := getLogger()

		db, err := gorm.Open(postgres.New(
			postgres.Config{
				DSN:                  dsn,
				PreferSimpleProtocol: true,
			}), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: false,
			},
			Logger: newLogger,
		})

		if err != nil {
			panic(fmt.Sprintf("failed to connect to the database: %v", err))
		}

		runMigrations(db)

		instance = db
	})

	return instance
}

func runMigrations(db *gorm.DB) {
	err := db.AutoMigrate(&model.User{})

	if err != nil {
		panic(fmt.Sprintf("failed to migrate the database: %v", err))
	}

	fmt.Println("Migrations done successfully.")
}

func getLogger() logger.Interface {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	return newLogger
}

func CloseDB() {
	sqlDB, err := instance.DB()

	if err != nil {
		log.Println("Error closing database connection:", err)
		return
	}

	err = sqlDB.Close()
	if err != nil {
		log.Println("Error closing database connection:", err)
	}
}
