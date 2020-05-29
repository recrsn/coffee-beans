package config

type Config struct {
	server       Server
	repositories []Repository
	contentRoot  string
}

func (c Config) Server() Server {
	return c.server
}

func (c Config) Repositories() []Repository {
	return c.repositories
}

func (c Config) ContentRoot() string {
	return c.contentRoot
}
