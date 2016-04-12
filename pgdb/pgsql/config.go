package pgsql

type Config struct {
	Host         string
	Port         string
	Username     string
	Password     string
	DBName       string
	SSL          bool
	MaxIdleConns int
	MaxOpenConns int
}

func NewConfig(host, port, username, password, dbname string, ssl bool) *Config {
	return NewFullConfig(host, port, username, password, dbname, ssl, 3, 3)
}

func NewFullConfig(host, port, username, password, dbname string, ssl bool, idle, open int) *Config {
	return &Config{Host: host, Port: port, Username: username, Password: password, DBName: dbname, SSL: ssl, MaxIdleConns: idle, MaxOpenConns: open}
}

func (s *Config) OpenSSl() *Config {
	s.SSL = true
	return s
}

func (s *Config) CloseSSl() *Config {
	s.SSL = false
	return s
}
