package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/controllers"
)

type AuthRoutes struct {
	router     chi.Router
	controller *controllers.AuthController
}

func NewAuthRoutes(controller *controllers.AuthController) *AuthRoutes {
	router := chi.NewRouter()

	return &AuthRoutes{
		router:     router,
		controller: controller,
	}
}

func (ar *AuthRoutes) Mount() http.Handler {
	ar.router.Post("/login", ar.controller.GetAccessTokenHandler)

	return ar.router
}
