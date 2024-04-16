package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	userUseCases "github.com/javiertelioz/clean-architecture-go/pkg/application/use_cases/user"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/contracts/services"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/dto"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/response"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/serializers"
)

type Services struct {
	UserService   services.UserService
	CryptoService services.CryptoService
	LoggerService services.LoggerService
}

type UserController struct {
	services *Services
}

func NewUserController(services *Services) *UserController {
	return &UserController{
		services: services,
	}
}

// GetUserByIdHandler godoc
//
//	@Summary		Get user account by ID
//	@Description	get string by ID
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id				path		int													true	"User ID"	default(1)
//	@Param			Accept-Language	header		string												false	"Language"	default(en-US)
//	@Success		200				{object}	response.Response{data=serializers.UserSerializer}	"desc"
//	@Failure		404				{object}	response.Response									"desc"
//	@Router			/api/v1/users/{id} [get]
func (c *UserController) GetUserByIdHandler(context *gin.Context) {
	id := context.Param("id")

	user, err := userUseCases.GetUserByIdUseCase(id, c.services.UserService, c.services.LoggerService)
	if err != nil {
		response.ErrorResponse(context, http.StatusNotFound, err.Error())
		return
	}

	payload := serializers.NewUserSerializer(user)
	response.SuccessResponse(context, http.StatusOK, payload)
}

// GetUsersHandler godoc
//
//	@Summary		List all users
//	@Description	Retrieves a list of all registered users
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			Accept-Language	header		string						false	"Language"	default(en-US)
//	@Success		200				{array}		serializers.UserSerializer	"desc"
//	@Failure		401				{object}	response.Error				"desc"
//	@Security		bearerAuth
//	@Router			/api/v1/users [get]
func (c *UserController) GetUsersHandler(context *gin.Context) {
	users, err := userUseCases.GetUsersUseCase(c.services.UserService, c.services.LoggerService)

	if err != nil {
		response.ErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	payload := serializers.NewUserListSerializer(users)

	response.SuccessResponse(context, http.StatusOK, payload)
}

// CreateUserHandler godoc
//
//	@Summary		Create a new user
//	@Description	Register a new user based on provided data
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			Accept-Language	header		string						false	"Language"	default(en-US)
//	@Param			User			body		serializers.UserSerializer	true	"User data to be created"
//	@Success		201				{object}	serializers.UserSerializer	"desc"
//	@Failure		400				{object}	response.Response			"desc"
//	@Failure		500				{object}	response.Response			"desc"
//	@Router			/api/v1/users [post]
func (c *UserController) CreateUserHandler(context *gin.Context) {
	var createUserDTO dto.CreateUserDTO

	if err := context.ShouldBindJSON(&createUserDTO); err != nil {
		response.ErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	userEntity := createUserDTO.ToEntity()
	user, err := userUseCases.CreateUserUseCase(
		userEntity,
		c.services.CryptoService,
		c.services.UserService,
		c.services.LoggerService,
	)
	if err != nil {
		response.ErrorResponse(context, http.StatusConflict, err.Error())
		return
	}

	payload := serializers.NewUserSerializer(user)
	response.SuccessResponse(context, http.StatusCreated, payload)
}

// UpdateUserHandler godoc
//
//	@Summary		Update a user
//	@Description	Modify an existing user based on provided data
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id				path		int							true	"User ID"	default(1)
//	@Param			User			body		serializers.UserSerializer	true	"User data to be updated"
//	@Param			Accept-Language	header		string						false	"Language"	default(en-US)
//	@Success		200				{object}	serializers.UserSerializer	"desc"
//	@Failure		400				{object}	response.Response			"desc"
//	@Failure		500				{object}	response.Response			"desc"
//	@Router			/api/v1/users/{id} [put]
func (c *UserController) UpdateUserHandler(context *gin.Context) {
	var updateUserDto dto.UpdateUserDTO

	id := context.Param("id")

	if err := context.ShouldBindJSON(&updateUserDto); err != nil {
		response.ErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	intID, err := strconv.Atoi(id)
	if err != nil {
		response.ErrorResponse(context, http.StatusBadRequest, err.Error())
		return
	}

	userEntity := updateUserDto.ToEntity()
	userEntity.ID = uint(intID)

	user, err := userUseCases.UpdateUserUseCase(
		userEntity,
		c.services.UserService,
		c.services.LoggerService,
	)
	if err != nil {
		response.ErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(context, http.StatusOK, user)
}

// DeleteUserHandler godoc
//
//	@Summary		Delete a user
//	@Description	Remove a user based on provided ID
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id				path		int					true	"User ID"	default(1)
//	@Param			Accept-Language	header		string				false	"Language"	default(en-US)
//	@Success		200				{object}	response.Response	"desc"
//	@Failure		500				{object}	response.Response	"desc"
//	@Router			/api/v1/users/{id} [delete]
func (c *UserController) DeleteUserHandler(context *gin.Context) {
	id := context.Param("id")

	err := userUseCases.DeleteUserUseCase(
		id,
		c.services.UserService,
		c.services.LoggerService,
	)
	if err != nil {
		response.ErrorResponse(context, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(context, http.StatusOK, nil)
}
