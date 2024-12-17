package resolve

import (
	"errors"
	"strconv"

	"github.com/graphql-go/graphql"
	userUseCases "github.com/javiertelioz/clean-architecture-go/pkg/application/use_cases/user"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/contracts/services"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/entity"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/serializers"
)

type UserResolver struct {
	cryptoService services.CryptoService
	userService   services.UserService
	loggerService services.LoggerService
}

func NewUserResolver(
	cryptoService services.CryptoService,
	userService services.UserService,
	loggerService services.LoggerService,
) *UserResolver {
	return &UserResolver{
		cryptoService: cryptoService,
		userService:   userService,
		loggerService: loggerService,
	}
}

func (r *UserResolver) CreateUser(p graphql.ResolveParams) (interface{}, error) {
	userInput, _ := p.Args["user"].(map[string]interface{})
	user := &entity.User{
		Name:     userInput["name"].(string),
		LastName: userInput["lastname"].(string),
		Surname:  userInput["surname"].(string),
		Email:    userInput["email"].(string),
		Password: userInput["password"].(string),
		Phone:    userInput["phone"].(string),
	}

	user, err := userUseCases.CreateUserUseCase(user, r.cryptoService, r.userService, r.loggerService)
	if err != nil {
		return nil, err
	}

	payload := serializers.NewUserSerializer(user)

	return payload, nil
}

func (r *UserResolver) GetUserById(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(string)

	if !ok {
		return nil, errors.New("id not provided")
	}

	user, err := userUseCases.GetUserByIdUseCase(id, r.userService, r.loggerService)
	if err != nil {
		return nil, err
	}

	payload := serializers.NewUserSerializer(user)

	return payload, nil
}

func (r *UserResolver) UpdateUser(p graphql.ResolveParams) (interface{}, error) {
	userInput, _ := p.Args["user"].(map[string]interface{})
	id, _ := p.Args["id"].(string)
	userId, _ := strconv.Atoi(id)

	user := &entity.User{
		ID:       uint(userId),
		Name:     userInput["name"].(string),
		LastName: userInput["lastname"].(string),
		Surname:  userInput["surname"].(string),
		Email:    userInput["email"].(string),
		Password: userInput["password"].(string),
		Phone:    userInput["phone"].(string),
	}

	user, err := userUseCases.UpdateUserUseCase(user, r.userService, r.loggerService)
	if err != nil {
		return nil, err
	}

	payload := serializers.NewUserSerializer(user)

	return payload, nil
}

func (r *UserResolver) DeleteUserById(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(string)

	if !ok {
		return nil, errors.New("id not provided")
	}

	err := userUseCases.DeleteUserUseCase(id, r.userService, r.loggerService)
	if err != nil {
		return nil, err
	}

	return struct{ ID string }{ID: id}, nil
}
