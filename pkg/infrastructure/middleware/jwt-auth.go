package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/javiertelioz/clean-architecture-go/pkg/application/use_cases/auth"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/logger"
	infrastructureService "github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/services"
)

const (
	// authorizationHeaderKey is the key for authorization header in the request.
	authorizationHeaderKey = "authorization"
	// authorizationType is the accepted authorization type.
	authorizationType = "bearer"
	// authorizationPayloadKey is the key for authorization payload in the context.
	authorizationPayloadKey = "authorization_payload"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)

		isEmpty := len(authorizationHeader) == 0

		if isEmpty {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": 0, "message": err.Error()})
			return
		}

		fields := strings.Fields(authorizationHeader)
		isValid := len(fields) == 2
		if !isValid {
			err := errors.New("authorization header format is invalid")
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"code": 0, "message": err.Error()})
			return
		}

		currentAuthorizationType := strings.ToLower(fields[0])
		if currentAuthorizationType != authorizationType {
			err := errors.New("authorization type is not supported")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": 0, "message": err.Error()})
			return
		}

		accessToken := fields[1]
		loggerService := logger.NewLogger()
		jwtService, _ := infrastructureService.NewJWTService("YELLOW SUBMARINE, BLACK WIZARDRY", loggerService)
		payload, err := auth.VerifyAccessTokenUserUseCase(accessToken, jwtService, loggerService)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": 0, "message": err.Error()})
			return
		}

		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}
