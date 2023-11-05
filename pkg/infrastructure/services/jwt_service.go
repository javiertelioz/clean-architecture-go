package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/contracts/services"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/entity"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/exceptions"
	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
)

type PasetoService struct {
	paseto       *paseto.V2
	symmetricKey []byte
	logger       services.LoggerService
}

/*func getSecretKey() string {
	secret, _ := config.GetConfig[string]("Jwt.secret")

	return secret
}*/

func NewJWTService(
	symmetricKey string,
	logger services.LoggerService,
) (services.JWTService, error) {
	validSymmetricKey := len(symmetricKey) == chacha20poly1305.KeySize

	if !validSymmetricKey {
		return nil, errors.New("invalid token key size")
	}

	return &PasetoService{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
		logger:       logger,
	}, nil
}

func (ps *PasetoService) Generate(user *entity.User) (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	duration := time.Duration(24)
	payload := entity.Token{
		ID:        id,
		UserID:    uint64(user.ID),
		Role:      user.Role,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	token, err := ps.paseto.Encrypt(ps.symmetricKey, payload, nil)

	return token, err
}

func (ps *PasetoService) Verify(token string) (*entity.Token, error) {
	var payload entity.Token

	err := ps.paseto.Decrypt(token, ps.symmetricKey, &payload, nil)
	if err != nil {
		return nil, exceptions.AuthInvalidToken()
	}

	isExpired := time.Now().After(payload.ExpiredAt)

	if isExpired {
		return nil, exceptions.AuthExpiredToken()
	}

	return &payload, nil
}
