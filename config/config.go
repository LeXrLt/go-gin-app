package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Source string `mapstructure:"source"`
	} `mapstructure:"database"`
	JWTSecret string `mapstructure:"jwt_secret"`
}

var Cfg *Config

func LoadConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return viper.Unmarshal(&Cfg)
}
