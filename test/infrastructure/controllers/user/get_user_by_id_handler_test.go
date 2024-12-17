package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/suite"

	"github.com/javiertelioz/clean-architecture-go/pkg/domain/entity"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/exceptions"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/controllers"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/response"
	"github.com/javiertelioz/clean-architecture-go/test/mocks/service"
)

type GetUserByIdHandlerTestSuite struct {
	suite.Suite
	route             *chi.Mux
	controller        *controllers.UserController
	mockUserService   *service.MockUserService
	mockLoggerService *service.MockLoggerService
	mockCryptoService *service.MockCryptoService
	request           *http.Request
	response          *httptest.ResponseRecorder
	userId            string
	user              *entity.User
}

func TestGetUserByIdHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(GetUserByIdHandlerTestSuite))
}

func (suite *GetUserByIdHandlerTestSuite) SetupTest() {
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
	suite.user = &entity.User{
		ID:       1,
		LastName: "Doe",
		Name:     "John",
		Email:    "john@example.com",
		Phone:    "+123456789",
		Password: "password123",
	}
}

func (suite *GetUserByIdHandlerTestSuite) givenUserId(id string) {
	suite.userId = id
}

func (suite *GetUserByIdHandlerTestSuite) givenUserServiceReturnsSuccess() {
	suite.mockUserService.On("GetUser", suite.userId).Return(suite.user, nil)
}

func (suite *GetUserByIdHandlerTestSuite) givenUserServiceReturnsError() {
	suite.mockUserService.On("GetUser", suite.userId).Return(nil, exceptions.UserNotFound())
}

func (suite *GetUserByIdHandlerTestSuite) whenCallGetUserByIdHandler() {
	suite.request, _ = http.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/user/%s", suite.userId), http.NoBody)
	suite.request.Header.Set("Accept-Language", "es-MX")
	suite.response = httptest.NewRecorder()

	suite.route.Get("/api/v1/user/{id}", suite.controller.GetUserByIdHandler)
	suite.route.ServeHTTP(suite.response, suite.request)
}

func (suite *GetUserByIdHandlerTestSuite) thenReturnSuccessResponse() {
	var responseBody response.Response
	err := json.Unmarshal(suite.response.Body.Bytes(), &responseBody)

	suite.NoError(err)
	suite.Equal(http.StatusOK, suite.response.Code)
	suite.Equal("Operation was successful", responseBody.Message)

	suite.mockUserService.AssertExpectations(suite.T())
}

func (suite *GetUserByIdHandlerTestSuite) thenReturnNoFoundResponse() {
	var responseBody response.Response
	err := json.Unmarshal(suite.response.Body.Bytes(), &responseBody)

	suite.NoError(err)
	suite.Equal(http.StatusNotFound, suite.response.Code)
	suite.Equal("USER_NOT_FOUND", responseBody.Message)

	suite.mockUserService.AssertExpectations(suite.T())
}

func (suite *GetUserByIdHandlerTestSuite) TestGetUserByIdHandlerSuccess() {
	// Given
	suite.givenUserId("1")
	suite.givenUserServiceReturnsSuccess()

	// When
	suite.whenCallGetUserByIdHandler()

	// Then
	suite.thenReturnSuccessResponse()
}

func (suite *GetUserByIdHandlerTestSuite) TestGetUserByIdHandlerNotFound() {
	// Given
	suite.givenUserId("10")
	suite.givenUserServiceReturnsError()

	// When
	suite.whenCallGetUserByIdHandler()

	// Then
	suite.thenReturnNoFoundResponse()
}
