package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/controllers"
)

type GraphQLRoutes struct {
	router     chi.Router
	controller *controllers.GraphQLController
}

func NewGraphQLRoutes(controller *controllers.GraphQLController) *GraphQLRoutes {
	router := chi.NewRouter()
	return &GraphQLRoutes{
		router:     router,
		controller: controller,
	}
}

func (g *GraphQLRoutes) Mount() http.Handler {
	g.router.Get("/sandbox", g.controller.SandboxHandler)
	g.router.Get("/", g.controller.GraphQLHandler)
	g.router.Post("/", g.controller.GraphQLHandler)

	return g.router
}
