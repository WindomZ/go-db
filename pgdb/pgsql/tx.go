package pgsql

import "github.com/jmoiron/sqlx"

type Tx struct {
	sqlx.Tx
	isolation bool
}

func (db *DB) BeginX(tx *sqlx.Tx) (*Tx, error) {
	if tx == nil {
		var err error
		if tx, err = db.Beginx(); err != nil {
			return nil, err
		}
		return &Tx{Tx: *tx, isolation: false}, nil
	}
	return &Tx{Tx: *tx, isolation: true}, nil
}

func (tx *Tx) Rollback() error {
	if tx.isolation {
		return nil
	}
	return tx.Tx.Rollback()
}

func (tx *Tx) Commit() error {
	if tx.isolation {
		return nil
	}
	return tx.Tx.Commit()
}
