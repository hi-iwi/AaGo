package ae

import (
	"fmt"
	"github.com/hi-iwi/AaGo/dict"
	"strconv"
)

type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

var (
	NotFound = &Error{404, "not found"} // refer to redis.Nil, sql.ErrNoRows
)

func New(err error) *Error {
	if err == nil {
		return nil
	}
	return NewErr(err.Error())
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
	return e.Msg + " [" + strconv.Itoa(e.Code) + "]"
}
func (e *Error) NotFound() bool {
	return e.Code == NotFound.Code && e.Msg == NotFound.Msg
}
func (e *Error) IsServerError() bool {
	return e.Code > 499
}
