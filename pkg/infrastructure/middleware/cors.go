package middleware

import (
	"net/http"
	"time"

	"github.com/go-chi/cors"
	"github.com/javiertelioz/clean-architecture-go/config"
)

func CORSMiddleware() func(next http.Handler) http.Handler {
	corsConfig, _ := config.GetConfig[config.CorsConfig]("Cors")
	return cors.Handler(cors.Options{
		AllowedOrigins:   corsConfig.AllowedOrigins, // []string{"https://foo.com"}
		AllowedMethods:   corsConfig.AllowMethods,   // []string{"GET", "POST", "PUT", "DELETE"}
		AllowedHeaders:   corsConfig.AllowHeaders,   // []string{"Content-Type", "Authorization"}
		ExposedHeaders:   corsConfig.ExposeHeaders,  // []string{"X-Total-Count"}
		AllowCredentials: corsConfig.AllowCredentials,
		MaxAge:           int((time.Duration(corsConfig.MaxAge) * time.Hour).Seconds()), // En segundos
	})
}
