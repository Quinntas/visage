package api

type Args struct {
	host string
	port string
}

func (a *Args) Addr() string {
	return a.host + ":" + a.port
}

const (
	DEFAULT_HOST = "localhost"
	DEFAULT_PORT = "6969"
)

func NewArgs(host string, port string) Args {
	if host == "" {
		host = DEFAULT_HOST
	}

	if port == "" {
		port = DEFAULT_PORT
	}

	return Args{
		host: host,
		port: port,
	}
}
