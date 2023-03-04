package gorm

type Config struct {
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	Name     string `mapstructure:"DB_NAME"`
	Host     string `mapstructure:"DB_HOST"`
	Port     string `mapstructure:"DB_PORT"`
	SslMode  string `mapstructure:"DB_SSLMODE"`
	Timezone string `mapstructure:"DB_TIMEZONE"`
}
