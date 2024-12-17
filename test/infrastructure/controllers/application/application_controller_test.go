package application

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/suite"

	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/controllers"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/serializers"
)

type ApplicationControllerTestSuite struct {
	suite.Suite
	route      *chi.Mux
	request    *http.Request
	response   *httptest.ResponseRecorder
	controller *controllers.ApplicationController
	body       *serializers.ApplicationSerializer
	error      error
	appName    string
}

func TestApplicationControllerTestSuite(t *testing.T) {
	suite.Run(t, new(ApplicationControllerTestSuite))
}

func (suite *ApplicationControllerTestSuite) SetupTest() {
	suite.appName = "TestApp"
	suite.route = chi.NewRouter()
	suite.controller = controllers.NewApplicationController(suite.appName)
	suite.body = &serializers.ApplicationSerializer{}
}

func (suite *ApplicationControllerTestSuite) whenCallApplicationHandler() {
	suite.route.Get("/", suite.controller.ApplicationInformationHandler)

	suite.request, suite.error = http.NewRequest(
		http.MethodGet,
		"/",
		bytes.NewBuffer(nil),
	)
	suite.response = httptest.NewRecorder()

	suite.route.ServeHTTP(suite.response, suite.request)

	suite.error = json.Unmarshal(suite.response.Body.Bytes(), suite.body)
}

func (suite *ApplicationControllerTestSuite) thenReturnSuccessResponse() {
	suite.NoError(suite.error)
	suite.Equal(http.StatusOK, suite.response.Code)
	suite.Equal("", suite.body.Version)
	suite.Equal("Welcome to TestApp", suite.body.Message)
}

func (suite *ApplicationControllerTestSuite) TestApplicationHandler() {
	suite.whenCallApplicationHandler()
	suite.thenReturnSuccessResponse()
}
