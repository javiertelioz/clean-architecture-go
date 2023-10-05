package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/controllers"
)

func SetupHelloRoutes(router *gin.RouterGroup) {
	controller := controllers.NewHelloController()

	router.GET("/hello/:name", controller.HelloHandler)
}
