package pgsql

import "errors"

var (
	ErrNotInited error = errors.New("pgsql: Sql is not initialized!")
	ErrNotConfig       = errors.New("pgsql: Sql is not register!")
)

var (
	ErrTxLock   error = errors.New("pgsql: error lock sql tx")
	ErrTxUnLock       = errors.New("pgsql: error unlock sql tx")
)
