package ae

import (
	"fmt"

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

func (e *Error) Error() string {
	return fmt.Sprintf("code %d, msg %s", e.Code, e.Msg)
}
