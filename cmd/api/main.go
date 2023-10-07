package main

import (
	"github.com/javiertelioz/clean-architecture-go/config"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/database"
)

func main() {
	config.LoadConfig()

	database.Connect()
	defer database.CloseDB()

	infrastructure.Server()
}
