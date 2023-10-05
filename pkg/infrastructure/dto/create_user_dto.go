package dto

import (
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/entity"
)

type CreateUserDTO struct {
	Name     string `json:"name" binding:"required"`
	LastName string `json:"lastName" binding:"required"`
	Surname  string `json:"surname"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Phone    string `json:"phone" binding:"required"`
}

func (dto *CreateUserDTO) ToEntity() *entity.User {
	return &entity.User{
		Name:     dto.Name,
		LastName: dto.LastName,
		Surname:  dto.Surname,
		Email:    dto.Email,
		Password: dto.Password,
		Phone:    dto.Phone,
	}
}
