package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/localizer"
)

const i18nContextKey = "i18n"

// TranslationMiddleware es un middleware que configura el localizador basado en el encabezado "Accept-Language".
func TranslationMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			lang := r.Header.Get("Accept-Language")
			i18n, ok := localizer.Get(lang)

			log.Printf("Use language %s, %v", lang, ok)

			ctx := context.WithValue(r.Context(), i18nContextKey, i18n)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}

// GetI18n
func GetI18n(r *http.Request) interface{} {
	return r.Context().Value(i18nContextKey)
}
