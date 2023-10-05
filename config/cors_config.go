package config

type CorsConfig struct {
	AllowedOrigins   []string `mapstructure:"allowed_origins" validate:"required"`
	AllowMethods     []string `mapstructure:"allow_methods" validate:"required"`
	AllowHeaders     []string `mapstructure:"allow_headers" validate:"required"`
	ExposeHeaders    []string `mapstructure:"expose_headers" validate:"required"`
	AllowCredentials bool     `mapstructure:"allow_credentials"`
	MaxAge           int64    `mapstructure:"max_age" validate:"gte=0"`
}
