package pgsql

import "github.com/jmoiron/sqlx"

type Tx struct {
	sqlx.Tx
	db        *DB
	isolation bool
}

func (db *DB) BeginX(tx *sqlx.Tx) (*Tx, error) {
	if tx == nil {
		db.txMutex.Lock()
		var err error
		if tx, err = db.Beginx(); err != nil {
			db.txMutex.Unlock()
			return nil, err
		}
		return &Tx{Tx: *tx, db: db, isolation: false}, nil
	}
	return &Tx{Tx: *tx, isolation: true}, nil
}

func (db *DB) WaitTxLock() {
	db.txMutex.Lock()
	db.txMutex.Unlock()
}

func (tx *Tx) GetTx() *sqlx.Tx {
	return &tx.Tx
}

func (tx *Tx) Rollback() error {
	if tx.isolation {
		return nil
	} else if tx.db != nil {
		tx.db.txMutex.Unlock()
	}
	return tx.Tx.Rollback()
}

func (tx *Tx) Commit() error {
	if tx.isolation {
		return nil
	} else if tx.db != nil {
		tx.db.txMutex.Unlock()
	}
	return tx.Tx.Commit()
}
