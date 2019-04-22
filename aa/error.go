package aa

import (
	"database/sql"
	"database/sql/driver"
	"fmt"

	"github.com/luexu/AaGo/aa"
	"github.com/luexu/AaGo/dict"
)

type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewError(code int, msg ...interface{}) *Error {
	m := ""
	if len(msg) == 0 {
		m = dict.Code2Msg(code)
	} else {
		if s, ok := msg[0].(string); ok {
			m = s
		} else if e, ok := msg[0].(error); ok {
			m = e.Error()
		}
	}

	return &Error{
		Code: code,
		Msg:  m,
	}
}

func NewSqlError(err error) *Error {
	if err == nil {
		return nil
	}
	switch err {
	case driver.ErrBadConn:
		aa.NewError(500, fmt.Sprintf("sql bad conn: ", err))
	case driver.ErrSkip:
		// ErrSkip may be returned by some optional interfaces' methods to
		// indicate at runtime that the fast path is unavailable and the sql
		// package should continue as if the optional interface was not
		// implemented. ErrSkip is only supported where explicitly
		// documented.
		aa.NewError(500, fmt.Sprintf("sql skip: ", err))
	case driver.ErrRemoveArgument:
		aa.NewError(500, fmt.Sprintf("sql remove argument: ", err))
	case sql.ErrConnDone:
		// ErrConnDone is returned by any operation that is performed on a connection
		// that has already been returned to the connection pool.
		aa.NewError(500, fmt.Sprintf("sql conn done: ", err))
	case sql.ErrTxDone:
		return aa.NewError(500, fmt.Sprintf("sql tx done: ", err))
	case sql.ErrNoRows:
		return aa.NewError(404, fmt.Sprintf("sql query not found: ", err))
	}
	return aa.NewError(500, err)
}

func (e *Error) Error() string {
	return fmt.Sprintf("code %d, msg %s", e.Code, e.Msg)
}
