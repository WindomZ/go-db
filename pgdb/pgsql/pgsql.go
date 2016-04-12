package pgsql

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	sqlx.DB
}

func NewDB(c *Config) *DB {
	db, err := NewDataBase(c)
	if err != nil {
		panic(err)
	}
	return &DB{*db}
}

func NewDataBase(c *Config) (*sqlx.DB, error) {
	if c == nil {
		return nil, ERR_NOT_CONFIG
	} else if len(c.Username) == 0 {
		return nil, ERR_NOT_CONFIG
	}
	db, err := sqlx.Connect("postgres", getDataSource(c))
	if err != nil {
		return db, err
	}
	db.SetMaxIdleConns(c.MaxIdleConns)
	db.SetMaxOpenConns(c.MaxOpenConns)
	return db, err
}

func getDataSource(c *Config) string {
	uri := fmt.Sprintf("postgres://%v:%v@%v:%v/%v", c.Username, c.Password, c.Host, c.Port, c.DBName)
	if !c.SSL {
		uri += "?sslmode=disable"
	}
	return uri
}
