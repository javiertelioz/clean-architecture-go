package routes

import (
	"github.com/gin-gonic/gin"

	_ "github.com/javiertelioz/clean-architecture-go/docs"
	swaggo "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
)

func SetupSwaggerRoutes(route *gin.Engine) {
	route.GET("/swagger/*any", swagger.WrapHandler(swaggo.Handler))
}
