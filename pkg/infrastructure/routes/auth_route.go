package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/javiertelioz/clean-architecture-go/config"
	domainService "github.com/javiertelioz/clean-architecture-go/pkg/domain/services"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/controllers"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/database"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/logger"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/repository"
	infrastructureService "github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/services"
)

func SetupAuthRoutes(router *gin.RouterGroup) {
	db := database.Connect()

	salt, _ := config.GetConfig[int]("Crypto.salt")

	loggerService := logger.NewLogger()
	userRepository := repository.NewUserRepository(db)
	cryptoService := infrastructureService.NewBcryptService(salt)

	jwtService, _ := infrastructureService.NewJWTService("YELLOW SUBMARINE, BLACK WIZARDRY", loggerService)
	userService := domainService.NewUserService(userRepository, loggerService)
	controller := controllers.NewAuthController(cryptoService, jwtService, userService, loggerService)

	router.POST("/auth/login", controller.GetAccessTokenHandler)
}
