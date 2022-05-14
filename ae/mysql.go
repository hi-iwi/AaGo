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
	pos := Caller(1)
	m := err.Error()

	switch err {
	case driver.ErrBadConn:
		NewError(500, pos+" sql bad conn: "+m)
	case driver.ErrSkip:
		// ErrSkip may be returned by some optional interfaces' methods to
		// indicate at runtime that the fast path is unavailable and the sql
		// package should continue as if the optional interface was not
		// implemented. ErrSkip is only supported where explicitly
		// documented.
		NewError(500, pos+" sql skip: "+m)
	case driver.ErrRemoveArgument:
		NewError(500, pos+" sql remove argument: "+m)
	case sql.ErrConnDone:
		// ErrConnDone is returned by any operation that is performed on a connection
		// that has already been returned to the connection pool.
		NewError(500, pos+" sql conn done: "+m)
	case sql.ErrTxDone:
		return NewError(500, pos+" sql tx done: "+m)
	case sql.ErrNoRows:
		return NotFound
	}

	dupExp := regexp.MustCompile(`Error\s\d+:\s+Duplicate\s+entry\s+'([^']*)'\s+for\s+key\s+'([^']*)'`)
	dupMatches := dupExp.FindAllStringSubmatch(m, -1)
	if dupMatches != nil && len(dupMatches) > 0 && len(dupMatches[0]) == 3 {
		m = pos + " duplicate entry `" + dupMatches[0][1] + "` for parameter `" + dupMatches[0][2] + "`"
		return NewError(400, m)
	}

	return NewError(500, pos+" sql error: "+m)
}
