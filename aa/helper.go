package aa

import (
	"encoding/json"
	"fmt"
	"github.com/hi-iwi/AaGo/atype"
	"log"
	"time"
)

func Println(args ...any) {
	ns := time.Now().Format("2006-01-02 15:04:05")
	for _, arg := range args {
		msg, ok := arg.(string)
		if !ok {
			s, err := json.Marshal(arg)
			if err != nil {
				msg = err.Error()
			} else {
				msg = string(s)
			}
		}
		log.Println(msg)
		fmt.Println(ns + " " + msg)
	}
}

func (app *Aa) Now() time.Time {
	return time.Now().In(app.Cfg().TimeLocation)
}

func (app *Aa) FmtNow() string {
	return app.Now().Format("2006-01-02 15:04:05")
}
func (app *Aa) FmtDate() string {
	return app.Now().Format("2006-01-02")
}
func (app *Aa) FmtTime() string {
	return app.Now().Format("15:04:05")
}

func (app *Aa) ParseDatetime(s string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02 15:04:05", s, app.Cfg().TimeLocation)
}
func (app *Aa) ParseDate(s string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02", s, app.Cfg().TimeLocation)
}
func (app *Aa) Datetime() atype.Datetime {
	return atype.ToDatetime(app.Now())
}
func (app *Aa) Date() atype.Date {
	return atype.ToDate(app.Now())
}
