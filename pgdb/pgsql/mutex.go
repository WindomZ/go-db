package pgsql

import "github.com/WindomZ/go-develop-kit/mutex"

type DBMutex struct {
	txMutex mutex.Mutex
}

func (mux *DBMutex) txMustLock() {
	//mux.txMutex.MustLock()
}

func (mux *DBMutex) txSafeLock() {
	//mux.txMutex.SafeLock()
}

func (mux *DBMutex) txUnsafeLock() {
	//mux.txMutex.UnsafeLock()
}

func (mux *DBMutex) txMustUnlock() {
	//mux.txMutex.MustUnlock()
}

func (mux *DBMutex) txSafeUnlock() {
	//mux.txMutex.SafeUnlock()
}

func (mux *DBMutex) txUnsafeUnlock() {
	//mux.txMutex.UnsafeUnlock()
}
