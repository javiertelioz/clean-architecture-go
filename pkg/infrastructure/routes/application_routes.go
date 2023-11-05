package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/javiertelioz/clean-architecture-go/config"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/controllers"
)

func SetupApplicationRoutes(route *gin.Engine) {
	appConfig, e := config.GetConfig[string]("AppName")

	fmt.Println(e)
	controller := controllers.NewApplicationController(appConfig)

	route.GET("/", controller.ApplicationInformationHandler)
}
