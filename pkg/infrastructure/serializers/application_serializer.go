package serializers

import (
	"time"
)

// ApplicationSerializer godoc
// @Description Application information
// ApplicationSerializer represents a serialized application
// swagger:model ApplicationSerializer
type ApplicationSerializer struct {
	Message string `json:"message" example:"Clean Architecture GO"`
	Version string `json:"version" example:"1.0.0"`
	Date    string `json:"date" example:"2023-09-17 22:32:15.572201"`
}

func NewApplicationSerializer(message string) *ApplicationSerializer {
	return &ApplicationSerializer{
		Message: message,
		Date:    time.Now().String(),
	}
}
