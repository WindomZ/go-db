package mgdb

import "errors"

var (
	ERR_NOT_INITED = errors.New("Mongo is not initialized!")
	ERR_NOT_CONFIG = errors.New("Mongo is not register!")
)
