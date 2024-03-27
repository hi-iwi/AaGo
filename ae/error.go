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

func New(err error) *Error {
	if err == nil {
		return nil
	}
	return NewErr(err.Error())
}

func NewE(code int) *Error {
	return NewError(code, dict.Code2Msg(code))
}

func NewErr(msg string, args ...any) *Error {
	return NewError(500, fmt.Sprintf(msg, args...))
}

func NewError(code int, msg string) *Error {
	return &Error{
		Code: code,
		Msg:  msg,
	}
}
func NewOk(ok bool) *Error {
	if ok {
		return nil
	}
	return NewE(500)
}
func E(e *Error, format string, a ...any) *Error {
	if e == nil || format == "" {
		return nil
	}
	if len(a) > 0 {
		format = fmt.Sprintf(format, a...)
	}
	e.Msg += " " + format
	return e
}

func Text(e *Error) string {
	if e == nil {
		return "nil"
	}
	return e.Text()
}

// 不要用 Error()，要不然跟 error.Error() 容易造成失误性panic
func (e *Error) Text() string {
	return e.Msg + " [" + strconv.Itoa(e.Code) + "]"
}
func (e *Error) NoMatched() bool {
	return e.Code == NotFound.Code || e.Code == NoRows.Code || e.Code == Gone.Code
}
func (e *Error) IsServerError() bool {
	return e.Code > 499
}

func (e *Error) IsRetryWith() bool {
	return e.Code == 449 && e.Msg != ""
}
