package ae

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
)

func NewSqlError(err error) *Error {
	if err == nil {
		return nil
	}
	switch err {
	case driver.ErrBadConn:
		NewError(500, fmt.Sprintf("sql bad conn: %s", err))
	case driver.ErrSkip:
		// ErrSkip may be returned by some optional interfaces' methods to
		// indicate at runtime that the fast path is unavailable and the sql
		// package should continue as if the optional interface was not
		// implemented. ErrSkip is only supported where explicitly
		// documented.
		NewError(500, fmt.Sprintf("sql skip: %s", err))
	case driver.ErrRemoveArgument:
		NewError(500, fmt.Sprintf("sql remove argument: %s", err))
	case sql.ErrConnDone:
		// ErrConnDone is returned by any operation that is performed on a connection
		// that has already been returned to the connection pool.
		NewError(500, fmt.Sprintf("sql conn done: %s", err))
	case sql.ErrTxDone:
		return NewError(500, fmt.Sprintf("sql tx done: %s", err))
	case sql.ErrNoRows:
		return NewError(404, fmt.Sprintf("sql query not found: %s", err))
	}
	return NewError(500, fmt.Sprintf("sql: %s", err))
}
