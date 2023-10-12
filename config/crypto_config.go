package config

type CryptoConfig struct {
	Salt int `mapstructure:"salt" validate:"required,numeric,gt=1,lte=15"`
}
