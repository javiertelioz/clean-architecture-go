package auth

import (
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/contracts/services"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/exceptions"
)

func GetAccessTokenUserUseCase(
	email string,
	password string,
	cryptoService services.CryptoService,
	jwtService services.JWTService,
	userService services.UserService,
	loggerService services.LoggerService,
) (string, error) {
	user, err := userService.GetUserByEmail(email)

	if err != nil {
		loggerService.Error(err.Error())
		return "", err
	}

	ok := cryptoService.Verify(password, user.Password)

	if ok != nil {
		return "", exceptions.AuthBadCredentials()
	}

	return jwtService.Generate(user)
}
