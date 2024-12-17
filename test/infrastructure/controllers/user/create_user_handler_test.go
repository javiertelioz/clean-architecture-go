package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/suite"

	"github.com/javiertelioz/clean-architecture-go/pkg/domain/entity"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/exceptions"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/controllers"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/dto"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/response"
	"github.com/javiertelioz/clean-architecture-go/test/mocks/service"
)

type CreateUserHandlerTestSuite struct {
	suite.Suite
	route             *chi.Mux
	controller        *controllers.UserController
	mockUserService   *service.MockUserService
	mockLoggerService *service.MockLoggerService
	mockCryptoService *service.MockCryptoService
	request           *http.Request
	response          *httptest.ResponseRecorder
	userDTO           dto.CreateUserDTO
	user              *entity.User
}

func TestCreateUserHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(CreateUserHandlerTestSuite))
}

func (suite *CreateUserHandlerTestSuite) SetupTest() {

	suite.route = chi.NewRouter()
	suite.mockUserService = new(service.MockUserService)
	suite.mockLoggerService = new(service.MockLoggerService)
	suite.mockCryptoService = new(service.MockCryptoService)

	services := &controllers.Services{
		CryptoService: suite.mockCryptoService,
		UserService:   suite.mockUserService,
		LoggerService: suite.mockLoggerService,
	}

	suite.controller = controllers.NewUserController(services)

	suite.userDTO = dto.CreateUserDTO{
		LastName: "Doe",
		Name:     "John",
		Email:    "john@example.com",
		Phone:    "+123456789",
		Password: "password123",
	}
	suite.user = &entity.User{
		ID:       1,
		LastName: "Doe",
		Name:     "John",
		Email:    "john@example.com",
		Phone:    "+123456789",
		Password: "password123",
	}
}

func (suite *CreateUserHandlerTestSuite) givenUserServiceReturns(user *entity.User, err error) {
	suite.mockUserService.On("CreateUser", suite.userDTO.ToEntity()).Return(user, err)
}

func (suite *CreateUserHandlerTestSuite) givenUserServiceByEmailReturns(user *entity.User, err error) {
	suite.mockUserService.On("GetUserByEmail", suite.user.Email).Return(user, err)
}

func (suite *CreateUserHandlerTestSuite) givenCryptoServiceReturnsHashedPassword(password string, err error) {
	suite.mockCryptoService.On("Hash", suite.user.Password).Return(password, err)
}

func (suite *CreateUserHandlerTestSuite) whenCallCreateUserHandler() {
	data, err := json.Marshal(suite.userDTO)
	suite.NoError(err)

	suite.request, err = http.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(data))
	suite.NoError(err)

	suite.request.Header.Set("Accept-Language", "es-MX")
	suite.response = httptest.NewRecorder()

	suite.route.Post("/api/v1/users", suite.controller.CreateUserHandler)
	suite.route.ServeHTTP(suite.response, suite.request)
}

func (suite *CreateUserHandlerTestSuite) thenReturnSuccessResponse() {
	var responseBody response.Response
	err := json.Unmarshal(suite.response.Body.Bytes(), &responseBody)

	suite.NoError(err)
	suite.Equal(http.StatusCreated, suite.response.Code)

	data, ok := responseBody.Data.(map[string]interface{})
	suite.True(ok, "Data should be a map")
	suite.Equal(suite.user.Name, data["name"])

	suite.mockUserService.AssertExpectations(suite.T())
	suite.mockCryptoService.AssertExpectations(suite.T())
}

func (suite *CreateUserHandlerTestSuite) thenReturnErrorResponse() {
	var responseBody response.Response
	err := json.Unmarshal(suite.response.Body.Bytes(), &responseBody)

	suite.NoError(err)
	suite.Equal(http.StatusConflict, suite.response.Code)
	suite.Equal("USER_PASSWORD_WRONG", responseBody.Message)
}

func (suite *CreateUserHandlerTestSuite) TestCreateUserHandlerSuccess() {
	// Given
	suite.givenCryptoServiceReturnsHashedPassword("password123", nil)
	suite.givenUserServiceByEmailReturns(nil, exceptions.UserNotFound())
	suite.givenUserServiceReturns(suite.user, nil)

	// When
	suite.whenCallCreateUserHandler()

	// Then
	suite.thenReturnSuccessResponse()
}

func (suite *CreateUserHandlerTestSuite) TestCreateUserHandlerWithErrorResult() {
	// Given
	suite.givenCryptoServiceReturnsHashedPassword("password123", errors.New("USER_PASSWORD_WRONG"))
	suite.givenUserServiceByEmailReturns(nil, exceptions.UserNotFound())
	suite.givenUserServiceReturns(nil, exceptions.UserPasswordWrong())

	// When
	suite.whenCallCreateUserHandler()

	// Then
	suite.thenReturnErrorResponse()
}
