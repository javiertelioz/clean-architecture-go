package hello

import (
	"bytes"
	"fmt"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/controllers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type HelloControllerTestSuite struct {
	suite.Suite
	route      *gin.Engine
	request    *http.Request
	response   *httptest.ResponseRecorder
	controller *controllers.HelloController
	name       string
	error      error
}

func TestHelloControllerTestSuite(t *testing.T) {
	suite.Run(t, new(HelloControllerTestSuite))
}

func (suite *HelloControllerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)

	suite.route = gin.Default()
	suite.controller = controllers.NewHelloController()
}

func (suite *HelloControllerTestSuite) givenName(name string) {
	suite.name = name
}

func (suite *HelloControllerTestSuite) whenHelloHandlerIsCalled() {
	suite.route.GET("/api/v1/hello/:name", suite.controller.HelloHandler)
	suite.request, suite.error = http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("/api/v1/hello/%s", suite.name),
		bytes.NewBuffer(nil))
	suite.request.Header.Set("Accept-Language", "es-MX")
	suite.response = httptest.NewRecorder()
	suite.route.ServeHTTP(suite.response, suite.request)
}

func (suite *HelloControllerTestSuite) thenReturnSuccessResponse() {
	suite.NoError(suite.error)
	suite.Equal(http.StatusOK, suite.response.Code)
	suite.Contains(suite.response.Body.String(), fmt.Sprintf("Hello %s", suite.name))
}

func (suite *HelloControllerTestSuite) thenReturnNoFoundResponse() {
	suite.NoError(suite.error)
	suite.Equal(http.StatusNotFound, suite.response.Code)
}

func (suite *HelloControllerTestSuite) TestHelloHandlerWithParam() {
	// Given
	suite.givenName("Joe")

	// When
	suite.whenHelloHandlerIsCalled()

	// Then
	suite.thenReturnSuccessResponse()
}

func (suite *HelloControllerTestSuite) TestHelloHandlerWithoutParam() {
	// Given
	suite.givenName("")

	// When
	suite.whenHelloHandlerIsCalled()

	// Then
	suite.thenReturnNoFoundResponse()
}
