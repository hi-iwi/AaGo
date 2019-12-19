package ae

import (
	"fmt"
	"strconv"

	"github.com/luexu/AaGo/dict"
)

type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewE(code int) *Error {
	return NewError(code, dict.Code2Msg(code))
}

func NewErr(msg string, args ...interface{}) *Error {
	return NewError(500, fmt.Sprintf(msg, args...))
}

func NewError(code int, msg string) *Error {
	return &Error{
		Code: code,
		Msg:  msg,
	}
}

func (e *Error) Error() string {
	return "code " + strconv.Itoa(e.Code) + ", msg " + e.Msg
}
