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

//	@title						Swagger Clean Architecture Go
//	@version					1.0
//	@description				This is a sample. You can find out more about Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/)
//	@termsOfService				http://swagger.io/terms/
//	@contact.name				API Support
//	@contact.url				http://www.swagger.io/support
//	@contact.email				support@docs.io
//
//	@license.name				Apache 2.0
//	@license.url				http://www.apache.org/licenses/LICENSE-2.0.html
//
//	@host						localhost:8080
//	@BasePath					/
//	@Schemes					http https
//
//	@securityDefinitions.apikey	bearerAuth
//	@in							header
//	@name						Authorization
//	@description				Type "Bearer" followed by a space and the access token.
//
//	@accept						json
//	@produce					json
//
//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/
func initRouters(router *gin.Engine) {
	routes.SetupApplicationRoutes(router)

	routes.SetupGraphQLRoutes(router)
	routes.SetupSwaggerRoutes(router)

	v1 := router.Group("/api/v1")
	{
		routes.SetupHelloRoutes(v1)
		routes.SetupUserController(v1)
		routes.SetupAuthRoutes(v1)
	}
}
