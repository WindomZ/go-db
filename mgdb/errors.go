package mgdb

import (
	"errors"
	"gopkg.in/mgo.v2"
)

var (
	ErrNoInited error = errors.New("Mongo is not initialized!")
	ErrNoConfig       = errors.New("Mongo is not register!")
	ErrNotFound       = mgo.ErrNotFound
	ErrCursor         = mgo.ErrCursor
	ErrExist          = errors.New("exist")
)
