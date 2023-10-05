package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/javiertelioz/clean-architecture-go/config"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/controllers"
)

func SetupApplicationRoutes(route *gin.Engine) {
	appConfig, _ := config.GetConfig[string]("AppName")
	controller := controllers.NewApplicationController(appConfig)

	route.GET("/", controller.ApplicationInformationHandler)
}
