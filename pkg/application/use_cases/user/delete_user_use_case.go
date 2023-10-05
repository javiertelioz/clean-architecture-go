package user

import (
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/contracts/services"
)

func DeleteUserUseCase(
	id string,
	userService services.UserService,
	logger services.LoggerService,
) error {
	err := userService.DeleteUser(id)

	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}
