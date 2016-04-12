package rddb

import "errors"

var (
	ERR_NOT_INITED = errors.New("Redis is not initialized!")
	ERR_NOT_CONFIG = errors.New("Redis is not register!")
)
