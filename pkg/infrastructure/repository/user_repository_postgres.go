package repository

import (
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/contracts/repository"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/entity"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/database/model"
	"gorm.io/gorm"
)

type UserRepositoryPostgres struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &UserRepositoryPostgres{
		db: db,
	}
}

func (r *UserRepositoryPostgres) Create(entityUser *entity.User) (*entity.User, error) {
	user := model.User{
		Name:     entityUser.Name,
		LastName: entityUser.LastName,
		Surname:  entityUser.Surname,
		Phone:    entityUser.Phone,
		Email:    entityUser.Email,
		Password: entityUser.Password,
	}

	err := r.db.Create(&user).Error

	if err != nil {
		return nil, err
	}

	return &entity.User{
		ID:       user.ID,
		Name:     user.Name,
		LastName: user.LastName,
		Surname:  user.Surname,
		Phone:    user.Phone,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}

func (r *UserRepositoryPostgres) FindByID(id string) (*entity.User, error) {
	user := &model.User{}
	err := r.db.First(user, id).Error

	if err != nil {
		return nil, err
	}

	return &entity.User{
		ID:       user.ID,
		Name:     user.Name,
		LastName: user.LastName,
		Surname:  user.Surname,
		Phone:    user.Phone,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}

func (r *UserRepositoryPostgres) FindAll() ([]*entity.User, error) {
	modelUsers := make([]*model.User, 0)
	tx := r.db.Limit(10).Find(&modelUsers)

	if tx.Error != nil {
		return nil, tx.Error
	}

	entityUsers := make([]*entity.User, len(modelUsers))
	for i, modelUser := range modelUsers {
		entityUsers[i] = &entity.User{
			ID:       modelUser.ID,
			Name:     modelUser.Name,
			LastName: modelUser.LastName,
			Surname:  modelUser.Surname,
			Phone:    modelUser.Phone,
			Email:    modelUser.Email,
			Password: modelUser.Password,
		}
	}

	return entityUsers, nil
}

func (r *UserRepositoryPostgres) Update(entityUser *entity.User) (*entity.User, error) {
	user := &model.User{
		Name:     entityUser.Name,
		LastName: entityUser.LastName,
		Surname:  entityUser.Surname,
		Phone:    entityUser.Phone,
		Email:    entityUser.Email,
		Password: entityUser.Password,
	}

	err := r.db.Where("id = ?", entityUser.ID).Updates(user).Error
	if err != nil {
		return nil, err
	}

	return &entity.User{
		ID:       user.ID,
		Name:     user.Name,
		LastName: user.LastName,
		Surname:  user.Surname,
		Phone:    user.Phone,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}

func (r *UserRepositoryPostgres) Delete(id int64) error {
	return r.db.Delete(&model.User{}, id).Error
}

func (r *UserRepositoryPostgres) FindByEmail(email string) (*entity.User, error) {
	user := &model.User{}
	err := r.db.Where("email = ?", email).First(user).Error
	if err != nil {
		return nil, err
	}
	return &entity.User{
		ID:       user.ID,
		Name:     user.Name,
		LastName: user.LastName,
		Surname:  user.Surname,
		Phone:    user.Phone,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}
