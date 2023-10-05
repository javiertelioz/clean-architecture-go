package config

type JwtConfig struct {
	Secret     string `mapstructure:"secret" validate:"required"`
	Expiration string `mapstructure:"expiration" validate:"required"`
}
