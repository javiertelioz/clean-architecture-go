package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/controllers"
)

type HelloRoutes struct {
	router     chi.Router
	controller *controllers.HelloController
}

func NewHelloRoutes(controller *controllers.HelloController) *HelloRoutes {
	router := chi.NewRouter()
	return &HelloRoutes{
		router:     router,
		controller: controller,
	}
}

func (h *HelloRoutes) Mount() http.Handler {
	h.router.Get("/{name}", h.controller.HelloHandler)

	return h.router
}
