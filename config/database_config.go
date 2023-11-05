package config

type DatabaseConfig struct {
	Name     string `validate:"required"`
	User     string `validate:"required"`
	Password string `validate:"required"`
	Host     string `validate:"required,hostname"`
	Port     int    `validate:"required,gt=0,lte=65535"`
	Schema   string
}
