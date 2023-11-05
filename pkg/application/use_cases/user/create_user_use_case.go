package user

import (
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/contracts/services"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/entity"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/exceptions"
)

func CreateUserUseCase(
	user *entity.User,
	cryptoService services.CryptoService,
	userService services.UserService,
	logger services.LoggerService,
) (*entity.User, error) {
	hashedPassword, err := cryptoService.Hash(user.Password)
	user.Password = hashedPassword

	exist, err := userService.GetUserByEmail(user.Email)

	if exist != nil {
		return nil, exceptions.UserAlreadyExists()
	}

	createdUser, err := userService.CreateUser(user)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return createdUser, nil
}
