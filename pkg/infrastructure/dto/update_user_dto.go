package dto

import (
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/entity"
)

type UpdateUserDTO struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

func (dto *UpdateUserDTO) ToEntity() *entity.User {
	return &entity.User{
		Name:     dto.Name,
		LastName: dto.LastName,
		Surname:  dto.Surname,
		Email:    dto.Email,
		Password: dto.Password,
		Phone:    dto.Phone,
	}
}
