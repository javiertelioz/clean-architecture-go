package application

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/controllers"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/serializers"
	"github.com/stretchr/testify/suite"
)

type ApplicationControllerTestSuite struct {
	suite.Suite
	route      *gin.Engine
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
	gin.SetMode(gin.TestMode)
	suite.appName = "TestApp"
	suite.route = gin.Default()
	suite.controller = controllers.NewApplicationController(suite.appName)
	suite.body = &serializers.ApplicationSerializer{}
}

func (suite *ApplicationControllerTestSuite) whenCallApplicationHandler() {
	suite.route.GET("/", suite.controller.ApplicationInformationHandler)
	suite.request, suite.error = http.NewRequest(
		http.MethodGet,
		"/",
		bytes.NewBuffer(nil))
	suite.response = httptest.NewRecorder()
	suite.route.ServeHTTP(suite.response, suite.request)

	suite.error = json.Unmarshal(suite.response.Body.Bytes(), &suite.body)
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
