package config

type Server struct {
	listenAddress string
	listenPort    int
}

func (s Server) ListenAddress() string {
	return s.listenAddress
}

func (s Server) ListenPort() int {
	return s.listenPort
}
