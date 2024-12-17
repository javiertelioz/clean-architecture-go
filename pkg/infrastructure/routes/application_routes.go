package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/controllers"
)

type ApplicationRoutes struct {
	router     chi.Router
	controller *controllers.ApplicationController
}

func NewApplicationRoutes(controller *controllers.ApplicationController) *ApplicationRoutes {
	router := chi.NewRouter()
	return &ApplicationRoutes{
		router:     router,
		controller: controller,
	}
}

func (a *ApplicationRoutes) Mount() http.Handler {
	a.router.Get("/", a.controller.ApplicationInformationHandler)

	return a.router
}
