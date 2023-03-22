package config

type Config struct {
	ClientId     string   `mapstructure:"clientid"`
	ClientSecret string   `mapstructure:"clientsecret"`
	RedirectUrl  string   `mapstructure:"redirecturl"`
	UserInfoApi  string   `mapstructure:"userinfoapi"`
	Scopes       []string `mapstructure:"scopes"`
}
