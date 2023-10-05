package config

type ServerConfig struct {
	Host string `mapstructure:"host" validate:"required,hostname"`
	Port string `mapstructure:"port" validate:"required,numeric,gt=0,lte=65535"`
}
