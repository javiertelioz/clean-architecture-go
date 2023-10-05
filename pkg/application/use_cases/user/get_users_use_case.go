package user

import (
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/contracts/services"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/entity"
)

func GetUsersUseCase(
	userService services.UserService,
	logger services.LoggerService,
) ([]*entity.User, error) {

	users, err := userService.GetUsers()

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return users, nil
}
