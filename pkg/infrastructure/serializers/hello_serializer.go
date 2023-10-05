package serializers

import (
	"github.com/gin-gonic/gin"
)

// HelloSerializer godoc
// @Description Hello information
// HelloSerializer represents a serialized hello
// swagger:model HelloSerializer
type HelloSerializer struct {
	Message string `json:"message" example:"Joe"`
}

func NewHelloSerializer(message string) *HelloSerializer {
	return &HelloSerializer{
		Message: message,
	}
}

func (s *HelloSerializer) Serialize() gin.H {
	return gin.H{
		"message": s.Message,
	}
}
