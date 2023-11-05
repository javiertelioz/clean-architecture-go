package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/controllers"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/graphql"
)

func SetupGraphQLRoutes(route *gin.Engine) {
	schema := graphql.NewSchema()
	controller := controllers.NewGraphQLController(schema)

	route.GET("/sandbox", controller.SandboxHandler)
	route.POST("/graphql", controller.GraphQLHandler)
	route.GET("/graphql", controller.GraphQLHandler)
}
