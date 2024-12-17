package controllers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"

	"github.com/javiertelioz/clean-architecture-go/pkg/application/use_cases/hello"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/contracts/services"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/serializers"
)

type HelloController struct {
	sayHelloUseCase hello.SayHelloUseCase
	loggerService   services.LoggerService
}

func NewHelloController(
	sayHelloUseCase hello.SayHelloUseCase,
	loggerService services.LoggerService,
) *HelloController {
	return &HelloController{
		sayHelloUseCase: sayHelloUseCase,
		loggerService:   loggerService,
	}
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
func (c *HelloController) HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	message := c.sayHelloUseCase.Execute(name)
	payload := serializers.NewHelloSerializer(message)

	c.loggerService.Trace("sdfadsf")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
