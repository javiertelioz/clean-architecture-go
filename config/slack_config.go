package config

type SlackConfig struct {
	AccessToken string `mapstructure:"access_token" validate:"required"`
	Channel     string `mapstructure:"channel" validate:"required"`
}
