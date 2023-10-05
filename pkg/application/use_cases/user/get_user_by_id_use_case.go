package user

import (
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/contracts/services"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/entity"
)

func GetUserByIdUseCase(
	id string,
	userService services.UserService,
	logger services.LoggerService,
) (*entity.User, error) {

	user, err := userService.GetUser(id)

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	// logger.Info(fmt.Sprintf("User information: %v", user,))
	return user, nil
}
