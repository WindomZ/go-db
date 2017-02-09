package mgdb

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

type DB struct {
	session  *mgo.Session
	dataBase string
	config   *Config
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
		return nil, ErrNoConfig
	} else if len(c.Host) == 0 {
		return nil, ErrNoConfig
	}
	db := &DB{dataBase: c.DBName, config: c}
	if _, err := db.getSession(); err != nil {
		return nil, err
	}
	return db, nil
}

func getDataSource(c *Config) string {
	return fmt.Sprintf("mongodb://%v:%v@%v:%v", c.Username, c.Password, c.Host, c.Port)
}

func (s *DB) getSession() (*mgo.Session, error) {
	if s.session == nil {
		var err error
		s.session, err = mgo.Dial(getDataSource(s.config))
		if err != nil {
			return nil, err
		}
		s.session.SetPoolLimit(s.config.MaxOpenConns)
	}
	return s.session.Clone(), nil
}

type Collection struct {
	mgo.Collection
}

func (s *DB) C(collection string, f func(*Collection) error) error {
	session, err := s.getSession()
	if err != nil {
		return err
	}
	defer session.Close()
	return f(&Collection{*session.DB(s.dataBase).C(collection)})
}
