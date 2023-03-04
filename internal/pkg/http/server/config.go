package server

type Config struct {
	Host string `mapstructure:"SERVER_HOST"`
	Port string `mapstructure:"SERVER_PORT"`
}
