package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	Port          string `mapstructure:"PORT"`
	AuthSvcUrl    string `mapstructure:"AUTH_SVC_URL"`
	ProductSvcUrl string `mapstructure:"PRODUCT_SVC_URL"`
	OrderSvcUrl   string `mapstructure:"ORDER_SVC_URL"`
}

func Load() (c Config, err error) {
	viper.AddConfigPath("./pkg/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	err = viper.ReadInConfig()
	if err != nil {
		log.Error().Msg("error reading config file")
		return Config{}, err
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		log.Error().Msg("error unmarshaling config vars to struct")
		return Config{}, err
	}

	return
}
