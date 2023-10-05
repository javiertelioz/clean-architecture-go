package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/services"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/controllers"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/database"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/logger"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/repository"
)

func SetupUserController(router *gin.RouterGroup) {
	db := database.Connect()

	loggerService := logger.NewLogger()
	userRepository := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepository, loggerService)
	controller := controllers.NewUserController(userService, loggerService)

	router.GET("/users", controller.GetUsersHandler)
	router.GET("/users/:id", controller.GetUserByIdHandler)
	router.POST("/users", controller.CreateUserHandler)
	router.PUT("/users/:id", controller.UpdateUserHandler)
	router.DELETE("/users/:id", controller.DeleteUserHandler)
}
