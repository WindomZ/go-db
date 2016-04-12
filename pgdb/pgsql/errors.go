package pgsql

import "errors"

var (
	ERR_NOT_INITED = errors.New("Sql is not initialized!")
	ERR_NOT_CONFIG = errors.New("Sql is not register!")
)
