package ae

import (
	"database/sql"
	"database/sql/driver"
	"regexp"
)

func NewSqlError(err error) *Error {
	if err == nil {
		return nil
	}
	msg, pos := CallerMsg(err.Error(), 1)

	switch err {
	case driver.ErrBadConn:
		NewError(500, pos+" sql bad conn: "+msg)
	case driver.ErrSkip:
		// ErrSkip may be returned by some optional interfaces' methods to
		// indicate at runtime that the fast path is unavailable and the sql
		// package should continue as if the optional interface was not
		// implemented. ErrSkip is only supported where explicitly
		// documented.
		NewError(500, pos+" sql skip: "+msg)
	case driver.ErrRemoveArgument:
		NewError(500, pos+" sql remove argument: "+msg)
	case sql.ErrConnDone:
		// ErrConnDone is returned by any operation that is performed on a connection
		// that has already been returned to the connection pool.
		NewError(500, pos+" sql conn done: "+msg)
	case sql.ErrTxDone:
		return NewError(500, pos+" sql tx done: "+msg)
	case sql.ErrNoRows:
		return NotFound // 通过在 asql层，对数组转换为 ae.NoRows
	}

	dupExp := regexp.MustCompile(`Duplicate\s+entry\s+'([^']*)'\s+for\s+key\s+'([^']*)'`)
	dupMatches := dupExp.FindAllStringSubmatch(msg, -1)
	if dupMatches != nil && len(dupMatches) > 0 && len(dupMatches[0]) == 3 {
		// dupMatches[0][1]
		return NewError(409, "key conflict")
	}

	return NewError(500, pos+" sql error: "+msg)
}
