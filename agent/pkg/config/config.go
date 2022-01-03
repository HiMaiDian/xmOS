package config

const (
	CONF_URL = "./conf/conf.ini"
)

type server struct {
	Port   string
	Server string
}

type client struct {
	Name         string
	Manager      string
	Email        string
	Description  string
	Token        string
	CollectTime  string
	PushBaseTime string
}

var Server server

var Client client

func Init() {
	// mapingTo("server", &Server)
	mapingTo("client", &Client)
}
