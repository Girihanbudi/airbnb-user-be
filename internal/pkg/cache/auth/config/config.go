package config

type Config struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Db       int    `mapstructure:"db"`
}
