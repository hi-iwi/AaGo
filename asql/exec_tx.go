package asql

import (
	"context"
	"database/sql"
	"github.com/hi-iwi/AaGo/ae"
)

type Tx struct {
	Tx *sql.Tx
}

// 为了避免跟db同名导致失误，这里统一加后缀
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

// defer tx.Recover
func (t *Tx) Recover() func() {
	return func() {
		if p := recover(); p != nil {
			t.Tx.Rollback()
		}
	}
}

func (t *Tx) PrepareTx(ctx context.Context, query string) (*sql.Stmt, *ae.Error) {
	stmt, err := t.Tx.PrepareContext(ctx, query)
	return stmt, ae.NewSqlError(err)
}

func (t *Tx) ExecTx(ctx context.Context, query string, args ...interface{}) (*sql.Stmt, sql.Result, *ae.Error) {
	stmt, e := t.PrepareTx(ctx, query)
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

func (t *Tx) ExecuteTx(ctx context.Context, query string, args ...interface{}) *ae.Error {
	stmt, _, e := t.ExecTx(ctx, query, args...)
	if e == nil {
		stmt.Close()
	}
	return e
}

func (t *Tx) InsertTx(ctx context.Context, query string, args ...interface{}) (uint, *ae.Error) {
	stmt, res, e := t.ExecTx(ctx, query, args...)
	if e != nil {
		return 0, e
	}
	defer stmt.Close()
	// 由于事务是先执行，后回滚或提交，所以可以先获取插入的ID，后commit()
	id, err := res.LastInsertId()
	return uint(id), ae.NewSqlError(err)
}

func (t *Tx) UpdateTx(ctx context.Context, query string, args ...interface{}) (int64, *ae.Error) {
	stmt, res, e := t.ExecTx(ctx, query, args...)
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
		stmt.QueryRowContext(ctx, i).&ScanTx()
	}
*/
func (t *Tx) BatchQueryRowTx(ctx context.Context, query string, margs ...[]interface{}) (*sql.Stmt, []*sql.Row, *ae.Error) {
	stmt, e := t.PrepareTx(ctx, query)
	if e != nil {
		return stmt, nil, e
	}
	rows := make([]*sql.Row, len(margs))
	for i, args := range margs {
		rows[i] = stmt.QueryRowContext(ctx, args...)
	}
	return stmt, rows, nil
}

func (t *Tx) QueryRowTx(ctx context.Context, query string, args ...interface{}) (*sql.Stmt, *sql.Row, *ae.Error) {
	stmt, e := t.PrepareTx(ctx, query)
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

func (t *Tx) ScanArgsTx(ctx context.Context, query string, args []interface{}, dest ...interface{}) *ae.Error {
	stmt, row, e := t.QueryRowTx(ctx, query, args...)
	if e != nil {
		return e
	}
	defer stmt.Close()
	return ae.NewSqlError(row.Scan(dest...))
}

func (t *Tx) ScanRowTx(ctx context.Context, query string, dest ...interface{}) *ae.Error {
	stmt, row, e := t.QueryRowTx(ctx, query)
	if e != nil {
		return e
	}
	defer stmt.Close()
	return ae.NewSqlError(row.Scan(dest...))
}

func (t *Tx) ScanTx(ctx context.Context, query string, id uint64, dest ...interface{}) *ae.Error {
	stmt, row, e := t.QueryRowTx(ctx, query, id)
	if e != nil {
		return e
	}
	defer stmt.Close()
	return ae.NewSqlError(row.Scan(dest...))
}

func (t *Tx) ScanXTx(ctx context.Context, query string, id string, dest ...interface{}) *ae.Error {
	stmt, row, e := t.QueryRowTx(ctx, query, id)
	if e != nil {
		return e
	}
	defer stmt.Close()
	return ae.NewSqlError(row.Scan(dest...))
}

// do not forget to close *sql.Rows
// 不要忘了关闭 rows
func (t *Tx) QueryTx(ctx context.Context, query string, args ...interface{}) (*sql.Stmt, *sql.Rows, *ae.Error) {
	stmt, e := t.PrepareTx(ctx, query)
	if e != nil {
		return stmt, nil, e
	}
	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		stmt.Close()
		return nil, nil, ae.NewSqlError(err)
	}
	return stmt, rows, nil
}
