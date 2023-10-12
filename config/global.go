package config

type Configuration struct {
	GinMode  string         `mapstructure:"gin_mode" validate:"required,oneof=debug release test"`
	AppEnv   string         `mapstructure:"app_env" validate:"required"`
	AppName  string         `mapstructure:"app_name" validate:"required"`
	Server   ServerConfig   `mapstructure:"server" validate:"required"`
	Database DatabaseConfig `mapstructure:"database" validate:"required"`
	Cors     CorsConfig     `mapstructure:"cors" validate:"required"`
	Crypto   CryptoConfig   `mapstructure:"crypto" validate:"required"`
	Jwt      JwtConfig      `mapstructure:"jwt" validate:"required"`
	Slack    SlackConfig    `mapstructure:"slack" validate:"required"`
}
