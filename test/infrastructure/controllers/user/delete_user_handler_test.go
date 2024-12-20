package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/suite"

	"github.com/javiertelioz/clean-architecture-go/pkg/domain/exceptions"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/controllers"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/response"
	"github.com/javiertelioz/clean-architecture-go/test/mocks/service"
)

type DeleteUserHandlerTestSuite struct {
	suite.Suite
	route             *chi.Mux
	controller        *controllers.UserController
	mockUserService   *service.MockUserService
	mockLoggerService *service.MockLoggerService
	mockCryptoService *service.MockCryptoService
	request           *http.Request
	response          *httptest.ResponseRecorder
	userId            string
}

func TestDeleteUserHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(DeleteUserHandlerTestSuite))
}

func (suite *DeleteUserHandlerTestSuite) SetupTest() {
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
	suite.request, _ = http.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/users/%s", suite.userId), http.NoBody)
	suite.request.Header.Set("Accept-Language", "es-MX")
	suite.response = httptest.NewRecorder()

	suite.route.Delete("/api/v1/users/{id}", suite.controller.DeleteUserHandler)
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
	// Given
	suite.givenUserId("1")
	suite.givenUserServiceReturnsSuccess()

	// When
	suite.whenCallDeleteUserHandler()

	// Then
	suite.thenReturnSuccessResponse()
}

func (suite *DeleteUserHandlerTestSuite) TestDeleteUserHandlerError() {
	// Given
	suite.givenUserId("10")
	suite.givenUserServiceReturnsError(exceptions.UserNotFound())

	// When
	suite.whenCallDeleteUserHandler()

	// Then
	suite.thenReturnErrorResponse()
}
