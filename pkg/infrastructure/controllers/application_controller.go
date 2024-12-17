package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/serializers"
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
//	@Router			/application [get]
func (c *ApplicationController) ApplicationInformationHandler(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("Welcome to %s", c.appName)
	payload := serializers.NewApplicationSerializer(message)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
