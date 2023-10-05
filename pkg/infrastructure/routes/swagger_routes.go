package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/javiertelioz/clean-architecture-go/docs"
)

func SetupSwaggerRoutes(route *gin.Engine) {
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
