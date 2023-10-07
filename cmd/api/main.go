package main

import (
	"github.com/javiertelioz/clean-architecture-go/config"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure"
)

func main() {
	config.LoadConfig()

	infrastructure.Server()
}
