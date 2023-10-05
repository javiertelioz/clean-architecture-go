package infrastructure

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/javiertelioz/clean-architecture-go/config"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/middleware"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/routes"
)

func Server() {
	server := initServer()

	serverConfig, _ := config.GetConfig[config.ServerConfig]("Server")
	addr := fmt.Sprintf("%s:%s", serverConfig.Host, serverConfig.Port)

	fmt.Printf("🚀 Starting application on: http://%s/\n", addr)

	err := server.Run(addr)

	if err != nil {
		panic(err)
	}
}

func initServer() *gin.Engine {
	setGinMode()
	router := gin.Default()

	initMiddleware(router)

	router.ForwardedByClientIP = true
	err := router.SetTrustedProxies([]string{"127.0.0.1"})

	if err != nil {
		return nil
	}

	initRouters(router)

	return router
}

func setGinMode() {
	mode, _ := config.GetConfig[string]("GinMode")
	gin.SetMode(mode)
}

func initMiddleware(router *gin.Engine) gin.IRoutes {
	return router.Use(
		middleware.TranslationMiddleware(),
		middleware.CORSMiddleware(),
	)
}

func initRouters(router *gin.Engine) {
	routes.SetupApplicationRoutes(router)

	v1 := router.Group("/api/v1")
	{
		routes.SetupHelloRoutes(v1)
		routes.SetupUserController(v1)
	}

	routes.SetupSwaggerRoutes(router)
}
