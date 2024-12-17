package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/javiertelioz/clean-architecture-go/pkg/application/use_cases/auth"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/logger"
	infrastructureService "github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/services"
)

const (
	// authorizationHeaderKey is the key for authorization header in the request.
	authorizationHeaderKey = "Authorization"
	// authorizationType is the accepted authorization type.
	authorizationType = "bearer"
	// authorizationPayloadKey is the key for authorization payload in the context.
	authorizationPayloadKey = "authorization_payload"
)

// AuthorizeJWT is a middleware for verifying JWT tokens.
func AuthorizeJWT() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authorizationHeader := r.Header.Get(authorizationHeaderKey)

			if len(authorizationHeader) == 0 {
				http.Error(w, `{"code": 0, "message": "authorization header is not provided"}`, http.StatusBadRequest)
				return
			}

			fields := strings.Fields(authorizationHeader)
			if len(fields) != 2 {
				http.Error(w, `{"code": 0, "message": "authorization header format is invalid"}`, http.StatusBadRequest)
				return
			}

			currentAuthorizationType := strings.ToLower(fields[0])
			if currentAuthorizationType != authorizationType {
				http.Error(w, `{"code": 0, "message": "authorization type is not supported"}`, http.StatusUnauthorized)
				return
			}

			accessToken := fields[1]
			loggerService := logger.NewLogger()
			jwtService, _ := infrastructureService.NewJWTService("YELLOW SUBMARINE, BLACK WIZARDRY", loggerService)
			payload, err := auth.VerifyAccessTokenUserUseCase(accessToken, jwtService, loggerService)
			if err != nil {
				http.Error(w, `{"code": 0, "message": "`+err.Error()+`"}`, http.StatusUnauthorized)
				return
			}

			// Add the payload to the request context
			ctx := r.Context()
			ctx = context.WithValue(ctx, authorizationPayloadKey, payload)
			r = r.WithContext(ctx)

			// Call the next handler
			next.ServeHTTP(w, r)
		})
	}
}
