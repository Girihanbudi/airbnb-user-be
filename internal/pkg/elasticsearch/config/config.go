package config

type Config struct {
	Addresses []string `mapstructure:"addresses"`
	Username  string   `mapstructure:"username"`
	Password  string   `mapstructure:"password"`
	MainIndex string   `mapstructure:"mainindex"`
	Separator string   `mapstructure:"separator"`
}
