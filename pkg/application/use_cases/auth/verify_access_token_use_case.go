package auth

import (
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/contracts/services"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/entity"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/exceptions"
)

func VerifyAccessTokenUserUseCase(
	token string,
	jwtService services.JWTService,
	loggerService services.LoggerService,
) (*entity.Token, error) {

	tokenInfo, err := jwtService.Verify(token)

	if err != nil {
		loggerService.Error(err.Error())
		return nil, exceptions.AuthInvalidToken()
	}

	return tokenInfo, nil
}
