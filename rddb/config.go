package rddb

type Config struct {
	Host     string
	Port     string
	Password string
	MaxIdle  int
}

func NewConfig(host, port, password string) *Config {
	return NewFullConfig(host, port, password, 3)
}

func NewConfigSimple(password string) *Config {
	return NewFullConfig("", "", password, 3)
}

func NewFullConfig(host, port, password string, idle int) *Config {
	if len(host) == 0 {
		host = "localhost"
	}
	if len(port) == 0 {
		port = "6379"
	}
	return &Config{Host: host, Port: port, Password: password, MaxIdle: idle}
}
