package controllers

import (
	"bytes"
	"github.com/graphql-go/graphql"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/controllers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockGraphQLHandler struct {
	mock.Mock
}

func (m *MockGraphQLHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.Called(w, r)
}

type GraphQLHandlerTestSuite struct {
	suite.Suite
	route              *gin.Engine
	controller         *controllers.GraphQLController
	request            *http.Request
	response           *httptest.ResponseRecorder
	mockGraphQLHandler *MockGraphQLHandler
	mockSchema         graphql.Schema
}

func TestGraphQLHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(GraphQLHandlerTestSuite))
}

func (suite *GraphQLHandlerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	suite.route = gin.Default()

	suite.mockGraphQLHandler = new(MockGraphQLHandler)
	suite.mockSchema = graphql.Schema{}
	suite.controller = controllers.NewGraphQLController(suite.mockSchema)
}

func (suite *GraphQLHandlerTestSuite) whenCallGraphQLHandler() {
	suite.request, _ = http.NewRequest(http.MethodPost, "/graphql", bytes.NewBufferString(""))
	suite.response = httptest.NewRecorder()

	suite.route.POST("/graphql", suite.controller.GraphQLHandler)
	suite.route.ServeHTTP(suite.response, suite.request)
}

func (suite *GraphQLHandlerTestSuite) thenHandlerIsInvoked() {
	suite.mockGraphQLHandler.AssertNumberOfCalls(suite.T(), "ServeHTTP", 1)
}

func (suite *GraphQLHandlerTestSuite) thenReturnHTML() {
	suite.Equal(http.StatusOK, suite.response.Code)
	suite.Contains(suite.response.Body.String(), "Must provide an operation.")
}

func (suite *GraphQLHandlerTestSuite) TestGraphQLHandlerInvocation() {
	// When
	suite.whenCallGraphQLHandler()

	// Then
	suite.thenReturnHTML()
}
