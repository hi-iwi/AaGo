package aorm

import (
	"context"
	"database/sql"
)

func (db *DB) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return db.DB.BeginTx(ctx, opts)
}

func (db *DB) Begin() (*sql.Tx, error) {
	return db.DB.BeginTx(context.Background(), nil)
}
