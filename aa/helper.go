package aa

import (
	"time"
)

func (app *Aa) Now() time.Time {
	return time.Now().In(app.Cfg().TimeLocation)
}

func (app *Aa) Datetime() string {
	return app.Now().Format("2006-01-02 15:04:05")
}
func (app *Aa) Date() string {
	return app.Now().Format("2006-01-02")
}
func (app *Aa) Time() string {
	return app.Now().Format("15:04:05")
}

func (app *Aa) ParseDatetime(s string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02 15:04:05", s, app.Cfg().TimeLocation)
}
