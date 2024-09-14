package config

import "github.com/spf13/viper"

type Config struct {
	AppUrl    string `mapstructure:"APP_URL"`
	SmptHost  string `mapstructure:"SMPT_HOST"`
	SmptPort  string `mapstructure:"SMPT_PORT"`
	FromEmail string `mapstructure:"EMAIL_FROM"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
