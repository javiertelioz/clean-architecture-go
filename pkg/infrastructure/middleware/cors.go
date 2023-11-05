package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/javiertelioz/clean-architecture-go/config"
)

func CORSMiddleware() gin.HandlerFunc {
	corsConfig, _ := config.GetConfig[config.CorsConfig]("Cors")
	c := cors.Config{
		AllowOrigins:     corsConfig.AllowedOrigins, // []string{"https://foo.com"},
		AllowMethods:     corsConfig.AllowMethods,
		AllowHeaders:     corsConfig.AllowHeaders,
		ExposeHeaders:    corsConfig.ExposeHeaders,
		AllowCredentials: corsConfig.AllowCredentials,
		MaxAge:           time.Duration(corsConfig.MaxAge) * time.Hour,
	}

	return cors.New(c)
}
