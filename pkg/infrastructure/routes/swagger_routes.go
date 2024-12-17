package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	swagger "github.com/swaggo/http-swagger/v2"

	_ "github.com/javiertelioz/clean-architecture-go/docs"
)

type SwaggerRoutes struct {
	router chi.Router
}

func NewSwaggerRoutes() *SwaggerRoutes {
	router := chi.NewRouter()
	return &SwaggerRoutes{
		router: router,
	}
}

func (s *SwaggerRoutes) Mount() http.Handler {
	s.router.Get("/*", swagger.Handler(
		swagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	return s.router
}
