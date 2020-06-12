package config

type Server struct {
	ListenAddress string `mapstructure:"listen_address"`
	ListenPort    int    `mapstructure:"listen_port"`
}

type Repository struct {
	Id      string
	Backend string
	Root    string
}

type Config struct {
	Server       Server
	Repositories []Repository
}
