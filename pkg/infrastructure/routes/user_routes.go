package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/javiertelioz/clean-architecture-go/config"
	domainServices "github.com/javiertelioz/clean-architecture-go/pkg/domain/services"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/controllers"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/database"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/logger"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/middleware"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/repository"
	infrastructureServices "github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/services"
)

func SetupUserController(router *gin.RouterGroup) {
	db := database.Connect()
	salt, _ := config.GetConfig[int]("Crypto.salt")

	loggerService := logger.NewLogger()
	userRepository := repository.NewUserRepository(db)
	cryptoService := infrastructureServices.NewBcryptService(salt)
	userService := domainServices.NewUserService(userRepository, loggerService)

	services := &controllers.Services{
		CryptoService: cryptoService,
		UserService:   userService,
		LoggerService: loggerService,
	}

	controller := controllers.NewUserController(services)

	router.GET("/users", middleware.AuthorizeJWT(), controller.GetUsersHandler)
	router.GET("/users/:id", controller.GetUserByIdHandler)
	router.POST("/users", controller.CreateUserHandler)
	router.PUT("/users/:id", controller.UpdateUserHandler)
	router.DELETE("/users/:id", controller.DeleteUserHandler)
}
