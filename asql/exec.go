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

// 批处理 prepare 性能会更好，但需要支持 mysqli；非批处理，不要使用 prepare，会造成多余开销
// 不要忘记 stmt.Close() 释放连接池资源
// Prepared statements take up server resources and should be closed after use.
func (d *DB) Prepare(ctx context.Context, query string) (*sql.Stmt, *ae.Error) {
	if d.err != nil {
		return nil, ae.NewSqlE(d.err, query)
	}
	stmt, err := d.DB.PrepareContext(ctx, query)
	return stmt, ae.NewSqlE(err, query)
}

/*
stmt close 必须要等到相关都执行完（包括  res.LastInsertId()  ,  row.Scan()
*/
func (d *DB) Execute(ctx context.Context, query string, args ...any) (sql.Result, *ae.Error) {
	res, err := d.DB.ExecContext(ctx, query, args...)
	return res, ae.NewSqlE(err, query, args...)
}

func (d *DB) Exec(ctx context.Context, query string, args ...any) *ae.Error {
	_, e := d.Execute(ctx, query, args...)
	return e
}
func (d *DB) Insert(ctx context.Context, query string, args ...any) (uint, *ae.Error) {
	res, e := d.Execute(ctx, query, args...)
	if e != nil {
		return 0, e
	}
	// 由于事务是先执行，后回滚或提交，所以可以先获取插入的ID，后commit()
	id, err := res.LastInsertId()
	return uint(id), ae.NewSqlE(err, query, args...)
}

func (d *DB) Update(ctx context.Context, query string, args ...any) (int64, *ae.Error) {
	res, e := d.Execute(ctx, query, args...)
	if e != nil {
		return 0, e
	}
	// 由于事务是先执行，后回滚或提交，所以可以先获取更新结果，后commit()
	id, err := res.RowsAffected()
	return id, ae.NewSqlE(err, query, args...)
}

// 批量查询
/*
	stmt,_ := db.Prepare("select count(*) from tb where id=?")
	defer stmt.Close()
	for i:=0;i<1000;i++{
		stmt.QueryRowContext(ctx, i).&Scan()
	}
*/
//func (d *DB) BatchQueryRow(ctx context.Context, query string, margs ...[]any) ([]*sql.Row, *ae.Error) {
//	stmt, e := d.Prepare(ctx, query)
//	if e != nil {
//		return nil, e
//	}
//	defer stmt.Close()
//	rows := make([]*sql.Row, len(margs))
//	for i, args := range margs {
//		rows[i] = stmt.QueryRowContext(ctx, args...)
//	}
//	return rows, nil
//}

func (d *DB) QueryRow(ctx context.Context, query string, args ...any) (*sql.Row, *ae.Error) {
	row := d.DB.QueryRowContext(ctx, query, args...)
	return row, ae.NewSqlE(row.Err(), query, args...)
}

func (d *DB) ScanArgs(ctx context.Context, query string, args []any, dest ...any) *ae.Error {
	row, e := d.QueryRow(ctx, query, args...)
	if e != nil {
		return e
	}
	return ae.NewSqlE(row.Scan(dest...), query, args...)
}
func (d *DB) ScanRow(ctx context.Context, query string, dest ...any) *ae.Error {
	row, e := d.QueryRow(ctx, query)
	if e != nil {
		return e
	}
	return ae.NewSqlE(row.Scan(dest...), query)
}

func (d *DB) Scan(ctx context.Context, query string, id uint64, dest ...any) *ae.Error {
	row, e := d.QueryRow(ctx, query, id)
	if e != nil {
		return e
	}
	return ae.NewSqlE(row.Scan(dest...), query, id)
}
func (d *DB) ScanX(ctx context.Context, query string, id string, dest ...any) *ae.Error {
	row, e := d.QueryRow(ctx, query, id)
	if e != nil {
		return e
	}
	return ae.NewSqlE(row.Scan(dest...), query, id)
}

// do not forget to close *sql.Rows
// 不要忘了关闭 rows
// 只有 QueryRow 找不到才会返回 ae.NotFound；Query 即使不存在，也是 nil
func (d *DB) Query(ctx context.Context, query string, args ...any) (*sql.Rows, *ae.Error) {
	rows, err := d.DB.QueryContext(ctx, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ae.NoRows
		}
		return nil, ae.NewSqlE(err, query, args...)
	}
	return rows, nil
}
