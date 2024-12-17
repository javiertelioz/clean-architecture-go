package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/graphql-go/graphql"
	"github.com/stretchr/testify/suite"

	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/controllers"
)

type SandboxHandlerTestSuite struct {
	suite.Suite
	route      *chi.Mux
	controller *controllers.GraphQLController
	request    *http.Request
	response   *httptest.ResponseRecorder
	mockSchema graphql.Schema
}

func TestSandboxHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(SandboxHandlerTestSuite))
}

func (suite *SandboxHandlerTestSuite) SetupTest() {
	suite.route = chi.NewRouter()

	suite.mockSchema = graphql.Schema{}
	suite.controller = controllers.NewGraphQLController(suite.mockSchema)
}

func (suite *SandboxHandlerTestSuite) whenCallSandboxHandler() {
	suite.request, _ = http.NewRequest(http.MethodGet, "/sandbox", nil)
	suite.response = httptest.NewRecorder()

	suite.route.Get("/sandbox", suite.controller.SandboxHandler)
	suite.route.ServeHTTP(suite.response, suite.request)
}

func (suite *SandboxHandlerTestSuite) thenReturnSandboxHTML() {
	suite.Equal(http.StatusOK, suite.response.Code)
	suite.Contains(suite.response.Body.String(), "http://localhost:8080/graphql")
}

func (suite *SandboxHandlerTestSuite) TestSandboxHandler() {
	// When
	suite.whenCallSandboxHandler()

	// Then
	suite.thenReturnSandboxHTML()
}
