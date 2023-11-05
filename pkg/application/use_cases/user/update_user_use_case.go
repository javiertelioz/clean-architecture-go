package user

import (
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/contracts/services"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/entity"
)

func UpdateUserUseCase(
	user *entity.User,
	userService services.UserService,
	logger services.LoggerService,
) (*entity.User, error) {
	updateUser, err := userService.UpdateUser(user)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return updateUser, nil
}
