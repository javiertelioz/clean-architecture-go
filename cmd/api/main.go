package main

import (
	"github.com/javiertelioz/clean-architecture-go/config"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure"
)

//	@title			Swagger Clean Architecture Go
//	@version		1.0
//	@description	This is a sample. You can find out more about Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/)
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@docs.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host						localhost:8080
//	@BasePath					/
//	@securityDefinitions.apikey	bearerAuth
//	@in							header
//	@name						Authorization
//
//	@accept						json
//	@produce					json
//
//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/

//	@Schemes	http https
func main() {
	config.LoadConfig()

	infrastructure.Server()
}
