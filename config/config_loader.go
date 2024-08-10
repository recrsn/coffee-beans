package config

import (
	"log"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

const (
	listenAddressConfigKey = "server.listen_address"
	listenPortConfigKey    = "server.listen_port"
	baseURLConfigKey       = "server.base_url"
)

func Load() (Config, error) {
	viper.SetDefault(listenAddressConfigKey, "")
	viper.SetDefault(listenPortConfigKey, "8080")
	viper.SetDefault(baseURLConfigKey, "/beans")

	viper.SetConfigFile("coffee-beans.yaml")
	viper.AddConfigPath("../etc")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		return Config{}, errors.Wrapf(err, "error loading config")
	}

	var c Config

	err = viper.Unmarshal(&c)

	log.Printf("Loaded config: %+v\n", c)

	if err != nil {
		return Config{}, errors.Wrapf(err, "error loading config")
	}

	return c, err
}
