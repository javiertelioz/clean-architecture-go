package services

type CryptoService interface {
	Hash(password string) (string, error)
	Verify(password, hashedPassword string) error
}
