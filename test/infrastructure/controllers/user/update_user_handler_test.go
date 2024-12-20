package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/suite"

	"github.com/javiertelioz/clean-architecture-go/pkg/domain/entity"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/controllers"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/dto"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/response"
	"github.com/javiertelioz/clean-architecture-go/test/mocks/service"
)

type UpdateUserHandlerTestSuite struct {
	suite.Suite
	route             *chi.Mux
	controller        *controllers.UserController
	mockUserService   *service.MockUserService
	mockLoggerService *service.MockLoggerService
	mockCryptoService *service.MockCryptoService
	request           *http.Request
	response          *httptest.ResponseRecorder
	payload           string
	userId            string
	user              *entity.User
	updateUserDto     dto.UpdateUserDTO
}

func TestUpdateUserHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(UpdateUserHandlerTestSuite))
}

func (suite *UpdateUserHandlerTestSuite) SetupTest() {

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
		Phone:    "+987654322",
		Password: "password123",
	}

	suite.updateUserDto = dto.UpdateUserDTO{
		Name:    "Jane",
		Surname: "jr",
		Email:   "jane@example.com",
	}
}

func (suite *UpdateUserHandlerTestSuite) givenUserId(id string) {
	suite.userId = id
}

func (suite *UpdateUserHandlerTestSuite) givenUpdateUserPayload(payload string) {
	suite.payload = payload
}

func (suite *UpdateUserHandlerTestSuite) givenUserServiceReturnsSuccess() {
	expectedUser := suite.updateUserDto.ToEntity()
	expectedUser.ID = 1
	suite.mockUserService.On("UpdateUser", expectedUser).Return(suite.user, nil)
}

func (suite *UpdateUserHandlerTestSuite) givenUserServiceReturnsError() {
	expectedUser := suite.updateUserDto.ToEntity()
	expectedUser.ID = 1
	suite.mockUserService.On("UpdateUser", expectedUser).Return(nil, fmt.Errorf("error updating user"))
}

func (suite *UpdateUserHandlerTestSuite) givenInvalidUpdateUserPayload(payload string) {
	suite.request, _ = http.NewRequest(http.MethodPut, fmt.Sprintf("/api/v1/users/%s", suite.userId), bytes.NewBufferString(payload))
	suite.request.Header.Set("Accept-Language", "es-MX")
	suite.response = httptest.NewRecorder()
}

func (suite *UpdateUserHandlerTestSuite) whenCallUpdateUserHandler() {
	suite.request, _ = http.NewRequest(http.MethodPut, fmt.Sprintf("/api/v1/users/%s", suite.userId), bytes.NewBufferString(suite.payload))
	suite.request.Header.Set("Accept-Language", "es-MX")
	suite.response = httptest.NewRecorder()

	suite.route.Put("/api/v1/users/{id}", suite.controller.UpdateUserHandler)
	suite.route.ServeHTTP(suite.response, suite.request)
}

func (suite *UpdateUserHandlerTestSuite) thenReturnSuccessResponse() {
	var responseBody response.Response
	err := json.Unmarshal(suite.response.Body.Bytes(), &responseBody)

	suite.NoError(err)
	suite.Equal(http.StatusOK, suite.response.Code)

	suite.mockUserService.AssertExpectations(suite.T())
}

func (suite *UpdateUserHandlerTestSuite) thenReturnBadRequestResponse() {
	var responseBody response.Response
	err := json.Unmarshal(suite.response.Body.Bytes(), &responseBody)

	suite.NoError(err)
	suite.Equal(http.StatusBadRequest, suite.response.Code)
}

func (suite *UpdateUserHandlerTestSuite) thenReturnErrorResponse() {
	var responseBody response.Response
	err := json.Unmarshal(suite.response.Body.Bytes(), &responseBody)

	suite.NoError(err)
	suite.Equal(http.StatusInternalServerError, suite.response.Code)

	suite.mockUserService.AssertExpectations(suite.T())
}

func (suite *UpdateUserHandlerTestSuite) thenReturnErrorBadRequest() {
	suite.Equal(http.StatusBadRequest, suite.response.Code)
}

func (suite *UpdateUserHandlerTestSuite) TestUpdateUserHandlerSuccess() {
	// Given
	suite.givenUserId("1")
	suite.givenUserServiceReturnsSuccess()
	expectedUser := suite.updateUserDto.ToEntity()
	expectedUser.ID = 1
	payload, _ := json.Marshal(suite.updateUserDto)
	suite.givenUpdateUserPayload(string(payload))

	// When
	suite.whenCallUpdateUserHandler()

	// Then
	suite.thenReturnSuccessResponse()
}

func (suite *UpdateUserHandlerTestSuite) TestUpdateUserHandlerWithInvalidId() {
	// Given
	suite.givenUserId("a")
	suite.givenUserServiceReturnsError()
	expectedUser := suite.updateUserDto.ToEntity()
	expectedUser.ID = 1
	payload, _ := json.Marshal(suite.updateUserDto)
	suite.givenUpdateUserPayload(string(payload))

	// When
	suite.whenCallUpdateUserHandler()

	// Then
	suite.thenReturnErrorBadRequest()
}

func (suite *UpdateUserHandlerTestSuite) TestUpdateUserHandlerError() {
	// Given
	suite.givenUserId("1")
	suite.givenUserServiceReturnsError()
	expectedUser := suite.updateUserDto.ToEntity()
	expectedUser.ID = 1
	payload, _ := json.Marshal(suite.updateUserDto)
	suite.givenUpdateUserPayload(string(payload))

	// When
	suite.whenCallUpdateUserHandler()

	// Then
	suite.thenReturnErrorResponse()
}

func (suite *UpdateUserHandlerTestSuite) TestUpdateUserHandlerWithInvalidPayload() {
	// Given
	suite.givenUserId("1")
	suite.givenUpdateUserPayload(`{"invalid_json_here": }`)

	// When
	suite.whenCallUpdateUserHandler()

	// Then
	suite.thenReturnBadRequestResponse()
}
