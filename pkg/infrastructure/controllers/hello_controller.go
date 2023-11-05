package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/javiertelioz/clean-architecture-go/pkg/application/use_cases/hello"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/serializers"
)

type HelloController struct{}

func NewHelloController() *HelloController {
	return &HelloController{}
}

// HelloHandler godoc
//
//	@Summary		Say Hello
//	@Description	Say Hello
//	@Tags			Hello
//	@Accept			json
//	@Produce		json
//	@Param			name			path		string	true	"Name"		default(Joe)
//	@Param			Accept-Language	header		string	false	"Language"	default(en-US)
//	@Success		200				{object}	serializers.HelloSerializer
//	@Router			/api/v1/hello/{name} [get]
func (c *HelloController) HelloHandler(context *gin.Context) {
	name := context.Params.ByName("name")
	message := hello.SayHelloUseCase(name)
	payload := serializers.NewHelloSerializer(message)

	context.JSON(http.StatusOK, payload)
}
