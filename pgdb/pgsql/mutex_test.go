package pgsql

import "testing"

var mux *DBMutex = new(DBMutex)

func TestDBMutexPlan1(t *testing.T) {
	mux.txMustLock()
	for i := 0; i < 110; i++ {
		mux.txUnsafeLock()
	}
	for i := 0; i < 100; i++ {
		mux.txUnsafeUnlock()
	}
	mux.txMustUnlock()
}

func TestDBMutexPlan2(t *testing.T) {
	mux.txMustLock()
	for i := 0; i < 100; i++ {
		mux.txUnsafeLock()
	}
	for i := 0; i < 110; i++ {
		mux.txUnsafeUnlock()
	}
	mux.txSafeUnlock()
}

func TestDBMutexPlan3(t *testing.T) {
	for i := 0; i < 100; i++ {
		mux.txUnsafeLock()
		mux.txUnsafeUnlock()
		mux.txMustLock()
		mux.txMustUnlock()
	}
	for i := 0; i < 100; i++ {
		mux.txMustLock()
		mux.txUnsafeUnlock()
		mux.txUnsafeLock()
		mux.txMustUnlock()
	}
	for i := 0; i < 100; i++ {
		mux.txMustLock()
		mux.txUnsafeLock()
		mux.txUnsafeLock()
		mux.txMustUnlock()
	}
}
