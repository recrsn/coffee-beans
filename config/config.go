package config

type Server struct {
	ListenAddress string `mapstructure:"listen_address"`
	ListenPort    int    `mapstructure:"listen_port"`
	BaseURL       string `mapstructure:"base_url"`
}

type RepositoryType string

const (
	RepositoryMaven   RepositoryType = "maven"
	RepositoryPypi    RepositoryType = "pypi"
	RepositoryGeneric RepositoryType = "generic"
)

type Repository struct {
	Id          string
	Name        string
	Description string
	Type        RepositoryType
	Root        string
}

type Config struct {
	Server       Server
	Repositories []Repository
}
