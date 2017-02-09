package mgdb

type Config struct {
	Host         string
	Port         string
	Username     string
	Password     string
	DBName       string
	MaxOpenConns int
}

func NewConfig(host, port, username, password, dbname string) *Config {
	return NewFullConfig(host, port, username, password, dbname, 3)
}

func NewSimpleConfig(username, password, dbname string) *Config {
	return NewFullConfig("", "", username, password, dbname, 3)
}

func NewFullConfig(host, port, username, password, dbname string, open int) *Config {
	if len(host) == 0 {
		host = "127.0.0.1"
	}
	if len(port) == 0 {
		port = "27017"
	}
	return &Config{Host: host, Port: port, Username: username, Password: password, DBName: dbname, MaxOpenConns: open}
}
