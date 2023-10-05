package serializers

import (
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/entity"
)

// UserSerializer godoc
// @Description User information
// UserSerializer represents a serialized user
// swagger:model UserSerializer
type UserSerializer struct {
	Id       uint   `json:"id" example:"1"`
	Name     string `json:"name" example:"John"`
	LastName string `json:"lastName" example:"Doe"`
	Surname  string `json:"surname" example:"Jr"`
	Email    string `json:"email" example:"john.doe@example.com"`
	Password string `json:"password" example:"password123"`
	Phone    string `json:"phone" example:"+1234567890"`
}

func NewUserSerializer(user *entity.User) *UserSerializer {
	return &UserSerializer{
		Id:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		LastName: user.LastName,
		Surname:  user.Surname,
		Password: user.Password,
		Phone:    user.Phone,
	}
}

func NewUserListSerializer(users []*entity.User) []*UserSerializer {
	userSerializers := make([]*UserSerializer, len(users))

	for i, user := range users {
		userSerializers[i] = NewUserSerializer(user)
	}

	return userSerializers
}
