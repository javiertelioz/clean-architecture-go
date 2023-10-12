package serializers

// TokenSerializer godoc
// @Description Token information
// TokenSerializer represents a serialized token
// swagger:model TokenSerializer
type TokenSerializer struct {
	Token string `json:"token" example:""`
}

func NewTokenSerializer(token string) *TokenSerializer {
	return &TokenSerializer{
		Token: token,
	}
}
