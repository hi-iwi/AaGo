package aorm

import (
	"context"
	"database/sql"

	"github.com/hi-iwi/AaGo/ae"
)

type DB struct {
	Schema string
	close  bool
	DB     *sql.DB
	err    error
}

func AliveDriver(schema string, db *sql.DB, close bool, err error) *DB {
	return &DB{
		Schema: schema,
		close:  close,
		DB:     db,
		err:    err,
	}
}

func Driver(schema string, db *sql.DB, err error) *DB {
	return &DB{
		close: true,
		DB:    db,
		err:   err,
	}
}

func (d *DB) Close() {
	if d.close {
		d.DB.Close()
	}
}

func (d *DB) Insert(ctx context.Context, query string, args ...interface{}) (uint, *ae.Error) {
	res, e := d.ExecContext(ctx, query, args...)
	if e != nil {
		return 0, e
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, ae.NewSqlError(err)
	}
	return uint(id), nil
}

func (d *DB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, *ae.Error) {
	var res sql.Result
	if d.err != nil {
		return res, ae.NewSqlError(d.err)
	}
	stmt, err := d.DB.PrepareContext(ctx, query)
	if err != nil {
		return res, ae.NewSqlError(err)
	}
	defer stmt.Close()
	res, err = stmt.ExecContext(ctx, args...)

	return res, ae.NewSqlError(err)
}

func (d *DB) Exec(query string, args ...interface{}) (sql.Result, *ae.Error) {
	return d.ExecContext(context.Background(), query, args...)
}

func (d *DB) QueryRowContext(ctx context.Context, query string, args ...interface{}) (*sql.Row, *ae.Error) {
	if d.err != nil {
		return nil, ae.NewSqlError(d.err)
	}
	return d.DB.QueryRowContext(ctx, query, args...), nil
}

func (d *DB) QueryRow(query string, args ...interface{}) (*sql.Row, *ae.Error) {
	return d.DB.QueryRowContext(context.Background(), query, args...), nil
}

func (d *DB) ScanRowContext(ctx context.Context, query string, dest ...interface{}) *ae.Error {
	row, e := d.QueryRowContext(ctx, query)
	if e != nil {
		return e
	}
	return ae.NewSqlError(row.Scan(dest...))
}

func (d *DB) ScanRow(query string, dest ...interface{}) *ae.Error {
	return d.ScanRowContext(context.Background(), query, dest...)
}

func (d *DB) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, *ae.Error) {
	if d.err != nil {
		return nil, ae.NewSqlError(d.err)
	}
	rows, err := d.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, ae.NewSqlError(err)
	}
	return rows, nil
}

func (d *DB) Query(query string, args ...interface{}) (*sql.Rows, *ae.Error) {
	return d.QueryContext(context.Background(), query, args...)
}

func (d *DB) CloseRows(rows *sql.Rows) *ae.Error {
	var err error
	if err = rows.Close(); err != nil {
		return ae.NewSqlError(err)
	}
	if err = rows.Err(); err != nil {
		return ae.NewSqlError(err)
	}
	return nil
}
