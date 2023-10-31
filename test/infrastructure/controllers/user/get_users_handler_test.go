package user

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/entity"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/controllers"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/response"
	"github.com/javiertelioz/clean-architecture-go/test/mocks/service"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type GetUsersHandlerTestSuite struct {
	suite.Suite
	route             *gin.Engine
	controller        *controllers.UserController
	mockUserService   *service.MockUserService
	mockLoggerService *service.MockLoggerService
	mockCryptoService *service.MockCryptoService
	request           *http.Request
	response          *httptest.ResponseRecorder
	users             []*entity.User
}

func TestGetUsersHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(GetUsersHandlerTestSuite))
}

func (suite *GetUsersHandlerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)

	suite.route = gin.Default()
	suite.mockUserService = new(service.MockUserService)
	suite.mockLoggerService = new(service.MockLoggerService)
	suite.mockCryptoService = new(service.MockCryptoService)
	suite.controller = controllers.NewUserController(suite.mockCryptoService, suite.mockUserService, suite.mockLoggerService)
	suite.users = []*entity.User{
		{
			ID:       1,
			LastName: "Doe",
			Name:     "John",
		},
		{
			ID:       2,
			LastName: "Smith",
			Name:     "Anna",
		},
	}
}

func (suite *GetUsersHandlerTestSuite) givenUserServiceReturnsSuccess() {
	suite.mockUserService.On("GetUsers").Return(suite.users, nil)
}

func (suite *GetUsersHandlerTestSuite) givenUserServiceReturnsError(err error) {
	suite.mockUserService.On("GetUsers").Return(nil, err)
}

func (suite *GetUsersHandlerTestSuite) whenCallGetUsersHandler() {
	suite.request, _ = http.NewRequest(http.MethodGet, "/api/v1/users", nil)
	suite.request.Header.Set("Accept-Language", "es-MX")
	suite.response = httptest.NewRecorder()

	suite.route.GET("/api/v1/users", suite.controller.GetUsersHandler)
	suite.route.ServeHTTP(suite.response, suite.request)
}

func (suite *GetUsersHandlerTestSuite) thenReturnSuccessResponse() {
	var responseBody response.Response
	err := json.Unmarshal(suite.response.Body.Bytes(), &responseBody)

	suite.NoError(err)
	suite.Equal(http.StatusOK, suite.response.Code)
	suite.Equal("Operation was successful", responseBody.Message)

	data, ok := responseBody.Data.([]interface{})
	suite.True(ok, "Data should be an array")
	suite.Equal(len(suite.users), len(data))

	suite.mockUserService.AssertExpectations(suite.T())
}

func (suite *GetUsersHandlerTestSuite) thenReturnErrorResponse() {
	suite.Equal(http.StatusInternalServerError, suite.response.Code)
	suite.mockUserService.AssertExpectations(suite.T())
}

func (suite *GetUsersHandlerTestSuite) TestGetUsersHandlerAndResponseSuccess() {
	// Given
	suite.givenUserServiceReturnsSuccess()

	// When
	suite.whenCallGetUsersHandler()

	// Then
	suite.thenReturnSuccessResponse()
}
func (suite *GetUsersHandlerTestSuite) TestGetUserHandlerResponseError() {
	// Given
	suite.givenUserServiceReturnsError(errors.New("internal Server Error"))

	// When
	suite.whenCallGetUsersHandler()

	// Then
	suite.thenReturnErrorResponse()
}
