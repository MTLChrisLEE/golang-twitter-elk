package config

import (
	"github.com/spf13/viper"
)

type TwitterConfig struct {
	API_KEY string `mapstructure:"API_KEY"`
	SECRET  string `mapstructure:"SECRET"`
}

type Config struct {
	Twitter TwitterConfig `mapstructure:"twitterConfig"`
}

func LoadConfig() (Config, error) {
	vp := viper.New()
	vp.SetConfigFile("./config/config.json")

	var config Config

	err := vp.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	err = vp.Unmarshal(&config)

	if err != nil {
		return Config{}, err
	}

	return config, nil
}
