package config

import (
	"github.com/spf13/viper"
)

const (
	contentRootKey         = "content_root"
	listenAddressConfigKey = "server.listen_address"
	listenPortConfigKey    = "server.listen_port"
	repositoriesConfigKey  = "repositories"
)

func Load() (Config, error) {
	viper.SetDefault(listenAddressConfigKey, "")
	viper.SetDefault(listenPortConfigKey, "8080")
	viper.SetDefault(contentRootKey, "data")

	viper.SetConfigFile("coffee-beans.yaml")
	viper.AddConfigPath("/etc")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		return Config{}, err
	}

	return Config{
		server:       readServerConfig(),
		repositories: readRepositoryConfig(),
		contentRoot:  viper.GetString(contentRootKey),
	}, nil
}

func readRepositoryConfig() []Repository {
	repositoryIds := viper.GetStringSlice(repositoriesConfigKey)

	repositories := make([]Repository, 0, len(repositoryIds))

	for _, id := range repositoryIds {
		repository := Repository{id}
		repositories = append(repositories, repository)
	}

	return repositories
}

func readServerConfig() Server {
	return Server{
		listenAddress: viper.GetString(listenAddressConfigKey),
		listenPort:    viper.GetInt(listenPortConfigKey),
	}
}
