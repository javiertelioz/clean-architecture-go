package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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
//	@Summary		Get user by ID
//	@Description	Retrieve user information by ID
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string						true	"User ID"
//	@Success		200	{object}	serializers.UserSerializer	"User Data"
//	@Failure		404	{object}	response.Response			"User Not Found"
//	@Router			/api/v1/users/{id} [get]
func (c *UserController) GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	user, err := userUseCases.GetUserByIdUseCase(id, c.services.UserService, c.services.LoggerService)
	if err != nil {
		response.ErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	payload := serializers.NewUserSerializer(user)
	response.SuccessResponse(w, http.StatusOK, payload)
}

// GetUsersHandler godoc
//
//	@Summary		List all users
//	@Description	Retrieve a list of all registered users
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		serializers.UserSerializer	"List of Users"
//	@Failure		500	{object}	response.Response			"Internal Server Error"
//	@Router			/api/v1/users [get]
func (c *UserController) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := userUseCases.GetUsersUseCase(c.services.UserService, c.services.LoggerService)

	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	payload := serializers.NewUserListSerializer(users)
	response.SuccessResponse(w, http.StatusOK, payload)
}

// CreateUserHandler godoc
//
//	@Summary		Create a new user
//	@Description	Register a new user with provided information
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.CreateUserDTO			true	"User Data"
//	@Success		201		{object}	serializers.UserSerializer	"Created User"
//	@Failure		400		{object}	response.Response			"Bad Request"
//	@Failure		409		{object}	response.Response			"Conflict"
//	@Router			/api/v1/users [post]
func (c *UserController) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var createUserDTO dto.CreateUserDTO

	if err := json.NewDecoder(r.Body).Decode(&createUserDTO); err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, err.Error())
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
		response.ErrorResponse(w, http.StatusConflict, err.Error())
		return
	}

	payload := serializers.NewUserSerializer(user)
	response.SuccessResponse(w, http.StatusCreated, payload)
}

// UpdateUserHandler godoc
//
//	@Summary		Update a user by ID
//	@Description	Modify an existing user's information
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string						true	"User ID"
//	@Param			request	body		dto.UpdateUserDTO			true	"User Data to Update"
//	@Success		200		{object}	serializers.UserSerializer	"Updated User"
//	@Failure		400		{object}	response.Response			"Bad Request"
//	@Failure		500		{object}	response.Response			"Internal Server Error"
//	@Router			/api/v1/users/{id} [put]
func (c *UserController) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var updateUserDTO dto.UpdateUserDTO
	id := chi.URLParam(r, "id")

	if err := json.NewDecoder(r.Body).Decode(&updateUserDTO); err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	intID, err := strconv.Atoi(id)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	userEntity := updateUserDTO.ToEntity()
	userEntity.ID = uint(intID)

	user, err := userUseCases.UpdateUserUseCase(
		userEntity,
		c.services.UserService,
		c.services.LoggerService,
	)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(w, http.StatusOK, user)
}

// DeleteUserHandler godoc
//
//	@Summary		Delete a user by ID
//	@Description	Remove a user's information by ID
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string				true	"User ID"
//	@Success		200	{object}	response.Response	"Deleted Successfully"
//	@Failure		500	{object}	response.Response	"Internal Server Error"
//	@Router			/api/v1/users/{id} [delete]
func (c *UserController) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := userUseCases.DeleteUserUseCase(
		id,
		c.services.UserService,
		c.services.LoggerService,
	)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.SuccessResponse(w, http.StatusOK, nil)
}
