package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/controllers"
)

type UserRoutes struct {
	router     chi.Router
	controller *controllers.UserController
}

func NewUserRoutes(controller *controllers.UserController) *UserRoutes {
	router := chi.NewRouter()
	return &UserRoutes{
		router:     router,
		controller: controller,
	}
}

func (ur *UserRoutes) Mount() http.Handler {
	ur.router.Get("/", ur.controller.GetUsersHandler)
	ur.router.Post("/", ur.controller.CreateUserHandler)
	ur.router.Get("/{id}", ur.controller.GetUserByIdHandler)
	ur.router.Put("/{id}", ur.controller.UpdateUserHandler)
	ur.router.Delete("/{id}", ur.controller.DeleteUserHandler)

	return ur.router
}
