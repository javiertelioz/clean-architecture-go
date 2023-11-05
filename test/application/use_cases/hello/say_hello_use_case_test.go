package hello

import (
	"fmt"
	"testing"

	"github.com/javiertelioz/clean-architecture-go/pkg/application/use_cases/hello"
	"github.com/stretchr/testify/suite"
)

type SayHelloUseCaseTestSuite struct {
	suite.Suite
	name   string
	result string
}

func TestSayHelloUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(SayHelloUseCaseTestSuite))
}

func (suite *SayHelloUseCaseTestSuite) SetupTest() {
	suite.result = "Hello Joe"
}

func (suite *SayHelloUseCaseTestSuite) givenName(name string) {
	suite.name = name
}

func (suite *SayHelloUseCaseTestSuite) whenSayHelloUseCaseIsCalled() {
	suite.result = hello.SayHelloUseCase(suite.name)
}

func (suite *SayHelloUseCaseTestSuite) thenSayHello() {
	suite.Equal(suite.result, fmt.Sprintf("Hello %s", suite.name))
}

func (suite *SayHelloUseCaseTestSuite) TestSayHelloUseCase() {
	// Given
	suite.givenName("Joe")

	// When
	suite.whenSayHelloUseCaseIsCalled()

	// Then
	suite.thenSayHello()
}
