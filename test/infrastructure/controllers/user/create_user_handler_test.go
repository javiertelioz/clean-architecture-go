package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/entity"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/exceptions"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/controllers"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/dto"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/response"
	"github.com/javiertelioz/clean-architecture-go/test/application/user_cases/user/mocks"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type CreateUserHandlerTestSuite struct {
	suite.Suite
	route             *gin.Engine
	controller        *controllers.UserController
	mockUserService   *mocks.MockUserService
	mockLoggerService *mocks.MockLoggerService
	mockCryptoService *mocks.MockCryptoService
	request           *http.Request
	response          *httptest.ResponseRecorder
	userDTO           dto.CreateUserDTO
	user              *entity.User
}

func TestCreateUserHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(CreateUserHandlerTestSuite))
}

func (suite *CreateUserHandlerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)

	suite.route = gin.Default()
	suite.mockUserService = new(mocks.MockUserService)
	suite.mockLoggerService = new(mocks.MockLoggerService)
	suite.mockCryptoService = new(mocks.MockCryptoService)
	suite.controller = controllers.NewUserController(suite.mockCryptoService, suite.mockUserService, suite.mockLoggerService)

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

func (suite *CreateUserHandlerTestSuite) givenUserServiceReturnsSuccess() {
	suite.mockUserService.On("CreateUser", suite.userDTO.ToEntity()).Return(suite.user, nil)
}

func (suite *CreateUserHandlerTestSuite) givenUserServiceReturnsError(err error) {
	suite.mockUserService.On("CreateUser", suite.userDTO.ToEntity()).Return(nil, err)
}

func (suite *CreateUserHandlerTestSuite) givenUserServiceByEmailReturnsSuccess() {
	suite.mockUserService.On("GetUserByEmail", suite.user.Email).Return(suite.user, nil)
}

func (suite *CreateUserHandlerTestSuite) givenUserServiceByEmailReturnsError() {
	suite.mockUserService.On("GetUserByEmail", suite.user.Email).Return(nil, exceptions.UserNotFound())
}

func (suite *CreateUserHandlerTestSuite) givenCryptoServiceReturnsSuccess() {
	suite.mockCryptoService.On("Hash", suite.user.Password).Return("password123", nil)
}

func (suite *CreateUserHandlerTestSuite) givenCryptoServiceReturnsError() {
	suite.mockCryptoService.On("Hash", suite.user.Password).Return(nil, errors.New("password_wrong"))
}

func (suite *CreateUserHandlerTestSuite) whenCallCreateUserHandler() {
	data, _ := json.Marshal(suite.userDTO)
	suite.request, _ = http.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(data))
	suite.request.Header.Set("Accept-Language", "es-MX")
	suite.response = httptest.NewRecorder()

	suite.route.POST("/api/v1/users", suite.controller.CreateUserHandler)
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
}

func (suite *CreateUserHandlerTestSuite) thenReturnErrorResponse() {
	var responseBody response.Response
	err := json.Unmarshal(suite.response.Body.Bytes(), &responseBody)

	suite.NoError(err)
	suite.Equal(http.StatusConflict, suite.response.Code)
	suite.Equal("USER_ALREADY_EXISTS", responseBody.Message)

	suite.mockUserService.AssertExpectations(suite.T())
}

func (suite *CreateUserHandlerTestSuite) TestCreateUserHandlerSuccess() {
	suite.givenCryptoServiceReturnsSuccess()
	suite.givenUserServiceByEmailReturnsError()
	suite.givenUserServiceReturnsSuccess()
	suite.whenCallCreateUserHandler()
	suite.thenReturnSuccessResponse()
}

func (suite *CreateUserHandlerTestSuite) TestCreateUserHandlerWithErrorResult() {
	suite.givenCryptoServiceReturnsSuccess()
	suite.givenUserServiceByEmailReturnsError()
	suite.givenUserServiceReturnsError(exceptions.UserAlreadyExists())
	suite.whenCallCreateUserHandler()
	suite.thenReturnErrorResponse()
}
