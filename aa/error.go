package aa

import (
	"database/sql"
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
	if err == sql.ErrNoRows {
		return aa.NewError(404, fmt.Sprintf("sql query not found: ", err))
	} else if err != nil {
		return aa.NewError(500, fmt.Sprintf("sql query error: ", err))
	}
	return nil
}

func (e *Error) Error() string {
	return fmt.Sprintf("code %d, msg %s", e.Code, e.Msg)
}
