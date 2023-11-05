package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/localizer"
)

func TranslationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		lang := c.GetHeader("Accept-Language")
		i18n, ok := localizer.Get(lang)

		log.Printf("Use language %s, %v", lang, ok)

		c.Set("i18n", i18n)

		c.Next()
	}
}
