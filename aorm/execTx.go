package aorm

import (
	"context"
	"database/sql"
	"github.com/hi-iwi/AaGo/ae"
)

type Tx struct {
	Tx *sql.Tx
}

func (d *DB) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, *ae.Error) {
	tx, err := d.DB.BeginTx(ctx, opts)
	if err != nil {
		return nil, ae.NewSqlError(err)
	}
	t := Tx{Tx: tx}
	return &t, nil
}

func (t *Tx) Rollback() *ae.Error {
	return ae.NewSqlError(t.Tx.Rollback())
}

func (t *Tx) Commit() *ae.Error {
	return ae.NewSqlError(t.Tx.Commit())
}
func (t *Tx) Prepare(ctx context.Context, query string) (*sql.Stmt, *ae.Error) {
	stmt, err := t.Tx.PrepareContext(ctx, query)
	return stmt, ae.NewSqlError(err)
}

func (t *Tx) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, *ae.Error) {
	stmt, e := t.Prepare(ctx, query)
	if e != nil {
		return nil, e
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, args...)
	return res, ae.NewSqlError(err)
}

func (t *Tx) Insert(ctx context.Context, query string, args ...interface{}) (uint, *ae.Error) {
	res, e := t.Exec(ctx, query, args...)
	if e != nil {
		return 0, e
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, ae.NewSqlError(err)
	}
	return uint(id), nil
}

func (t *Tx) QueryRow(ctx context.Context, query string, args ...interface{}) (*sql.Row, *ae.Error) {
	stmt, e := t.Prepare(ctx, query)
	if e != nil {
		return nil, e
	}
	defer stmt.Close()
	row := stmt.QueryRowContext(ctx, args...)
	return row, nil
}

// 批量查询
/*
	stmt,_ := db.Prepare("select count(*) from tb where id=?")
	defer stmt.Close()
	for i:=0;i<1000;i++{
		stmt.QueryRowContext(ctx, i).&Scan()
	}
*/
func (t *Tx) BatchQueryRow(ctx context.Context, query string, margs ...[]interface{}) ([]*sql.Row, *ae.Error) {
	stmt, e := t.Prepare(ctx, query)
	if e != nil {
		return nil, e
	}
	defer stmt.Close()
	rows := make([]*sql.Row, len(margs))
	for i, args := range margs {
		rows[i] = stmt.QueryRowContext(ctx, args...)
	}
	return rows, nil
}

func (t *Tx) ScanRow(ctx context.Context, query string, dest ...interface{}) *ae.Error {
	row, e := t.QueryRow(ctx, query)
	if e != nil {
		return e
	}
	return ae.NewSqlError(row.Scan(dest...))
}

// do not forget to close *sql.Rows
// 不要忘了关闭 rows
func (t *Tx) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, *ae.Error) {
	stmt, e := t.Prepare(ctx, query)
	if e != nil {
		return nil, e
	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(ctx, args...)
	return rows, ae.NewSqlError(err)
}
