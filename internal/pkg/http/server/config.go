package server

type Config struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}
