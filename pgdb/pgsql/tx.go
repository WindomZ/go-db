package pgsql

import "github.com/jmoiron/sqlx"

type Tx struct {
	sqlx.Tx
	db        *DB
	isolation bool
}

func (db *DB) BeginX(tx *sqlx.Tx) (*Tx, error) {
	if tx == nil {
		db.txMustLock()
		var err error
		if tx, err = db.Beginx(); err != nil {
			db.txMustUnlock()
			return nil, err
		}
		return &Tx{Tx: *tx, db: db, isolation: false}, nil
	}
	db.txSafeLock()
	return &Tx{Tx: *tx, isolation: true}, nil
}

func (tx *Tx) txSafeUnlock() {
	if tx != nil && tx.db != nil {
		tx.db.txSafeUnlock()
	}
}

func (tx *Tx) txUnsafeUnlock() {
	if tx != nil && tx.db != nil {
		tx.db.txUnsafeUnlock()
	}
}

func (tx *Tx) GetTx() *sqlx.Tx {
	return &tx.Tx
}

func (tx *Tx) Rollback() error {
	if tx.isolation {
		tx.txUnsafeUnlock()
		return nil
	}
	tx.txSafeUnlock()
	return tx.Tx.Rollback()
}

func (tx *Tx) Commit() error {
	if tx.isolation {
		tx.txUnsafeUnlock()
		return nil
	}
	tx.txSafeUnlock()
	return tx.Tx.Commit()
}
