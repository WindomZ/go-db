package rddb

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

var (
	ErrNil           = redis.ErrNil
	ErrPoolExhausted = redis.ErrPoolExhausted
)

type DB struct {
	pool   *redis.Pool
	config *Config
}

func NewDB(c *Config) *DB {
	db, err := NewDataBase(c)
	if err != nil {
		panic(err)
	}
	return db
}

func NewDataBase(c *Config) (*DB, error) {
	if c == nil {
		return nil, ERR_NOT_CONFIG
	} else if len(c.Host) == 0 {
		return nil, ERR_NOT_CONFIG
	} else if c.MaxIdle <= 0 {
		c.MaxIdle = 3
	}
	pool := &redis.Pool{
		MaxIdle:     c.MaxIdle,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			rc, err := redis.Dial("tcp", fmt.Sprintf("%v:%v", c.Host, c.Port))
			if err != nil {
				return nil, err
			} else if len(c.Password) != 0 {
				if _, err := rc.Do("AUTH", c.Password); err != nil {
					rc.Close()
					return nil, err
				}
			}
			return rc, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return &DB{pool: pool, config: c}, nil
}

func (s *DB) G() redis.Conn {
	return s.pool.Get()
}

func (s *DB) Close() error {
	return s.pool.Close()
}

func (s *DB) Err() error {
	c := s.G()
	defer c.Close()
	return c.Err()
}

func (s *DB) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	c := s.G()
	defer c.Close()
	return c.Do(commandName, args...)
}

func (s *DB) Send(commandName string, args ...interface{}) error {
	c := s.G()
	defer c.Close()
	return c.Send(commandName, args...)
}

func (s *DB) Flush() error {
	c := s.G()
	defer c.Close()
	return c.Flush()
}

func (s *DB) Receive() (reply interface{}, err error) {
	c := s.G()
	defer c.Close()
	return c.Receive()
}
