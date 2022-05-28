package asql

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
		/*
			https://pkg.go.dev/database/sql#Open
			The returned DB is safe for concurrent use by multiple goroutines and maintains its own pool of idle connections. Thus, the Open function should be called just once. It is rarely necessary to close a DB.
			Go MySQL 自带连接池，不不建议关闭 DB。
		*/
		d.DB.Close()
	}
}

// prepare 性能会更好，但需要支持 mysqli
// 不要忘记 stmt.Close() 释放连接池资源
// Prepared statements take up server resources and should be closed after use.
func (d *DB) Prepare(ctx context.Context, query string) (*sql.Stmt, *ae.Error) {
	if d.err != nil {
		return nil, ae.NewSqlError(d.err)
	}
	stmt, err := d.DB.PrepareContext(ctx, query)
	return stmt, ae.NewSqlError(err)
}

func (d *DB) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, *ae.Error) {
	stmt, e := d.Prepare(ctx, query)
	if e != nil {
		return nil, e
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, args...)
	return res, ae.NewSqlError(err)
}

func (d *DB) Insert(ctx context.Context, query string, args ...interface{}) (uint, *ae.Error) {
	res, e := d.Exec(ctx, query, args...)
	if e != nil {
		return 0, e
	}
	// 由于事务是先执行，后回滚或提交，所以可以先获取插入的ID，后commit()
	id, err := res.LastInsertId()
	return uint(id), ae.NewSqlError(err)
}

func (d *DB) Update(ctx context.Context, query string, args ...interface{}) (int64, *ae.Error) {
	res, e := d.Exec(ctx, query, args...)
	if e != nil {
		return 0, e
	}
	// 由于事务是先执行，后回滚或提交，所以可以先获取更新结果，后commit()
	id, err := res.RowsAffected()
	return id, ae.NewSqlError(err)
}

func (d *DB) QueryRow(ctx context.Context, query string, args ...interface{}) (*sql.Row, *ae.Error) {
	stmt, e := d.Prepare(ctx, query)
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
func (d *DB) BatchQueryRow(ctx context.Context, query string, margs ...[]interface{}) ([]*sql.Row, *ae.Error) {
	stmt, e := d.Prepare(ctx, query)
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

func (d *DB) ScanRow(ctx context.Context, query string, dest ...interface{}) *ae.Error {
	row, e := d.QueryRow(ctx, query)
	if e != nil {
		return e
	}
	return ae.NewSqlError(row.Scan(dest...))
}

func (d *DB) Scan(ctx context.Context, query string, id interface{}, dest ...interface{}) *ae.Error {
	row, e := d.QueryRow(ctx, query, id)
	if e != nil {
		return e
	}
	return ae.NewSqlError(row.Scan(dest...))
}

// do not forget to close *sql.Rows
// 不要忘了关闭 rows
func (d *DB) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, *ae.Error) {
	stmt, e := d.Prepare(ctx, query)
	if e != nil {
		return nil, e
	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(ctx, args...)
	return rows, ae.NewSqlError(err)
}
