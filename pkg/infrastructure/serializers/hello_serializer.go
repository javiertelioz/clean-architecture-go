package serializers

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
