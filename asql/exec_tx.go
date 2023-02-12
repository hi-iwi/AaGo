package asql

import (
	"context"
	"database/sql"
	"github.com/hi-iwi/AaGo/ae"
)

type Tx struct {
	Tx *sql.Tx
}

func (d *DB) Begin(ctx context.Context, opts *sql.TxOptions) (*Tx, *ae.Error) {
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

// defer tx.Recover
func (t *Tx) Recover() func() {
	return func() {
		if p := recover(); p != nil {
			t.Tx.Rollback()
		}
	}
}

func (t *Tx) Prepare(ctx context.Context, query string) (*sql.Stmt, *ae.Error) {
	stmt, err := t.Tx.PrepareContext(ctx, query)
	return stmt, ae.NewSqlError(err)
}

func (t *Tx) Exec(ctx context.Context, query string, args ...interface{}) (*sql.Stmt, sql.Result, *ae.Error) {
	stmt, e := t.Prepare(ctx, query)
	if e != nil {
		return stmt, nil, e
	}
	res, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		stmt.Close()
		return nil, res, ae.NewSqlError(err)
	}
	return stmt, res, nil
}

func (t *Tx) Execute(ctx context.Context, query string, args ...interface{}) *ae.Error {
	stmt, _, e := t.Exec(ctx, query, args...)
	if e == nil {
		stmt.Close()
	}
	return e
}

func (t *Tx) Insert(ctx context.Context, query string, args ...interface{}) (uint, *ae.Error) {
	stmt, res, e := t.Exec(ctx, query, args...)
	if e != nil {
		return 0, e
	}
	defer stmt.Close()
	// 由于事务是先执行，后回滚或提交，所以可以先获取插入的ID，后commit()
	id, err := res.LastInsertId()
	return uint(id), ae.NewSqlError(err)
}

func (t *Tx) Update(ctx context.Context, query string, args ...interface{}) (int64, *ae.Error) {
	stmt, res, e := t.Exec(ctx, query, args...)
	if e != nil {
		return 0, e
	}
	defer stmt.Close()
	// 由于事务是先执行，后回滚或提交，所以可以先获取更新结果，后commit()
	id, err := res.RowsAffected()
	return id, ae.NewSqlError(err)
}

// 批量查询
/*
	stmt,_ := db.Prepare("select count(*) from tb where id=?")
	defer stmt.Close()
	for i:=0;i<1000;i++{
		stmt.QueryRowContext(ctx, i).&Scan()
	}
*/
func (t *Tx) BatchQueryRow(ctx context.Context, query string, margs ...[]interface{}) (*sql.Stmt, []*sql.Row, *ae.Error) {
	stmt, e := t.Prepare(ctx, query)
	if e != nil {
		return stmt, nil, e
	}
	rows := make([]*sql.Row, len(margs))
	for i, args := range margs {
		rows[i] = stmt.QueryRowContext(ctx, args...)
	}
	return stmt, rows, nil
}

func (t *Tx) QueryRow(ctx context.Context, query string, args ...interface{}) (*sql.Stmt, *sql.Row, *ae.Error) {
	stmt, e := t.Prepare(ctx, query)
	if e != nil {
		return stmt, nil, e
	}
	row := stmt.QueryRowContext(ctx, args...)
	if row.Err() != nil {
		stmt.Close()
		return nil, nil, ae.NewSqlError(row.Err())
	}
	return stmt, row, nil
}

func (t *Tx) ScanArgs(ctx context.Context, query string, args []interface{}, dest ...interface{}) *ae.Error {
	stmt, row, e := t.QueryRow(ctx, query, args...)
	if e != nil {
		return e
	}
	defer stmt.Close()
	return ae.NewSqlError(row.Scan(dest...))
}

func (t *Tx) ScanRow(ctx context.Context, query string, dest ...interface{}) *ae.Error {
	stmt, row, e := t.QueryRow(ctx, query)
	if e != nil {
		return e
	}
	defer stmt.Close()
	return ae.NewSqlError(row.Scan(dest...))
}

func (t *Tx) Scan(ctx context.Context, query string, id uint64, dest ...interface{}) *ae.Error {
	stmt, row, e := t.QueryRow(ctx, query, id)
	if e != nil {
		return e
	}
	defer stmt.Close()
	return ae.NewSqlError(row.Scan(dest...))
}

func (t *Tx) ScanX(ctx context.Context, query string, id string, dest ...interface{}) *ae.Error {
	stmt, row, e := t.QueryRow(ctx, query, id)
	if e != nil {
		return e
	}
	defer stmt.Close()
	return ae.NewSqlError(row.Scan(dest...))
}

// do not forget to close *sql.Rows
// 不要忘了关闭 rows
func (t *Tx) Query(ctx context.Context, query string, args ...interface{}) (*sql.Stmt, *sql.Rows, *ae.Error) {
	stmt, e := t.Prepare(ctx, query)
	if e != nil {
		return stmt, nil, e
	}
	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		stmt.Close()
		if err == sql.ErrNoRows {
			return nil, nil, ae.NoRows
		}
		return nil, nil, ae.NewSqlError(err)
	}
	return stmt, rows, nil
}
