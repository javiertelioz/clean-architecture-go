package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/serializers"
	"net/http"
)

type ApplicationController struct {
	appName string
}

func NewApplicationController(appName string) *ApplicationController {
	return &ApplicationController{
		appName: appName,
	}
}

// ApplicationInformationHandler godoc
//
//	@Summary	Retrieve application information
//	@Schemes
//	@Description	Retrieve application information
//	@Tags			Application
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	serializers.ApplicationSerializer
//	@Security		bearerAuth
//	@Router			/ [get]
func (c *ApplicationController) ApplicationInformationHandler(context *gin.Context) {
	message := fmt.Sprintf("Welcome to %s", c.appName)
	payload := serializers.NewApplicationSerializer(message)

	context.JSON(http.StatusOK, payload)
}
