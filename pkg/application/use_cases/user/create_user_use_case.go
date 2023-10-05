package user

import (
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/contracts/services"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/entity"
)

func CreateUserUseCase(
	user *entity.User,
	userService services.UserService,
	logger services.LoggerService,
) (*entity.User, error) {
	createdUser, err := userService.CreateUser(user)

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return createdUser, nil
}
