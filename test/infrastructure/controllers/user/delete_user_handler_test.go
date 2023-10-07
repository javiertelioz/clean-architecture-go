package user

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/exceptions"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/controllers"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/response"
	"github.com/javiertelioz/clean-architecture-go/test/application/user_cases/user/mocks"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type DeleteUserHandlerTestSuite struct {
	suite.Suite
	route             *gin.Engine
	controller        *controllers.UserController
	mockUserService   *mocks.MockUserService
	mockLoggerService *mocks.MockLoggerService
	request           *http.Request
	response          *httptest.ResponseRecorder
	userId            string
}

func TestDeleteUserHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(DeleteUserHandlerTestSuite))
}

func (suite *DeleteUserHandlerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)

	suite.route = gin.Default()
	suite.mockUserService = new(mocks.MockUserService)
	suite.mockLoggerService = new(mocks.MockLoggerService)
	suite.controller = controllers.NewUserController(suite.mockUserService, suite.mockLoggerService)
}

func (suite *DeleteUserHandlerTestSuite) givenUserId(id string) {
	suite.userId = id
}

func (suite *DeleteUserHandlerTestSuite) givenUserServiceReturnsSuccess() {
	suite.mockUserService.On("DeleteUser", suite.userId).Return(nil)
}

func (suite *DeleteUserHandlerTestSuite) givenUserServiceReturnsError(error error) {

	suite.mockUserService.On("DeleteUser", suite.userId).Return(error)
}

func (suite *DeleteUserHandlerTestSuite) whenCallDeleteUserHandler() {
	suite.request, _ = http.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/users/%s", suite.userId), nil)
	suite.request.Header.Set("Accept-Language", "es-MX")
	suite.response = httptest.NewRecorder()

	suite.route.DELETE("/api/v1/users/:id", suite.controller.DeleteUserHandler)
	suite.route.ServeHTTP(suite.response, suite.request)
}

func (suite *DeleteUserHandlerTestSuite) thenReturnSuccessResponse() {
	var responseBody response.Response
	err := json.Unmarshal(suite.response.Body.Bytes(), &responseBody)

	suite.NoError(err)
	suite.Equal(http.StatusOK, suite.response.Code)

	suite.mockUserService.AssertExpectations(suite.T())
}

func (suite *DeleteUserHandlerTestSuite) thenReturnErrorResponse() {
	var responseBody response.Response
	err := json.Unmarshal(suite.response.Body.Bytes(), &responseBody)

	suite.NoError(err)
	suite.Equal(http.StatusInternalServerError, suite.response.Code)

	suite.mockUserService.AssertExpectations(suite.T())
}

func (suite *DeleteUserHandlerTestSuite) TestDeleteUserHandlerSuccess() {
	suite.givenUserId("1")
	suite.givenUserServiceReturnsSuccess()
	suite.whenCallDeleteUserHandler()
	suite.thenReturnSuccessResponse()
}

func (suite *DeleteUserHandlerTestSuite) TestDeleteUserHandlerError() {
	suite.givenUserId("10")
	suite.givenUserServiceReturnsError(exceptions.UserNotFound())
	suite.whenCallDeleteUserHandler()
	suite.thenReturnErrorResponse()
}