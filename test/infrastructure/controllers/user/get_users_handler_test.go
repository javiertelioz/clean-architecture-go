package user

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/entity"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/controllers"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/response"
	"github.com/javiertelioz/clean-architecture-go/test/application/user_cases/user/mocks"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type GetUsersHandlerTestSuite struct {
	suite.Suite
	route             *gin.Engine
	controller        *controllers.UserController
	mockUserService   *mocks.MockUserService
	mockLoggerService *mocks.MockLoggerService
	mockCryptoService *mocks.MockCryptoService
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
	suite.mockUserService = new(mocks.MockUserService)
	suite.mockLoggerService = new(mocks.MockLoggerService)
	suite.mockCryptoService = new(mocks.MockCryptoService)
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
	suite.givenUserServiceReturnsSuccess()
	suite.whenCallGetUsersHandler()
	suite.thenReturnSuccessResponse()
}
func (suite *GetUsersHandlerTestSuite) TestGetUserHandlerResponseError() {
	suite.givenUserServiceReturnsError(errors.New("internal Server Error"))
	suite.whenCallGetUsersHandler()
	suite.thenReturnErrorResponse()
}
