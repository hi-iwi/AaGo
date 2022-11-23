package aenum

import "time"

const AaBirthDate = int64(690134400) // 秒   1991-11-15

// https://dev.mysql.com/doc/refman/8.0/en/datetime.html
// DATETIME values is '1000-01-01 00:00:00.000000' to '9999-12-31 23:59:59.999999',
// TIMESTAMP values is '1970-01-01 00:00:01.000000' to '2038-01-19 03:14:07.999999'
// 两个时间最小、最大均不同，所以一定要开启 sql-mode  NO_ZERO_DATE，保持最小为 0000-00-00 00:00:00
const DateMin = "0000-00-00"
const DateMax = "9999-12-31"
const DateRegExp = `([12]\d{3}-[01]\d-[0-3]\d)|(0000-00-00)|(9999-12-31)`

const DatetimeMin = "0000-00-00 00:00:00"
const DatetimeMax = "9999-12-31 23:59:59"
const DatetimeRegExp = `([12]\d{3}-[01]\d-[0-3]\d\s[0-2]\d:[0-5]\d:[0-5]\d)|(0000-00-00\s00:00:00)|(9999-12-31\s23:59:59)`

const GoDateLayout = "2006-01-02"
const GoDatetimeLayout = "2006-01-02 15:04:05"
const GoTimeLayout = "15:04:05"

func IsMinDate(d string) bool {
	return d == "" || d == "0000" || d == "0000-00" || d == DateMin || d == "1000" || d == "1000-01" || d == "1000-01-01"
}
func IsMaxDate(d string) bool {
	return d == "9999" || d == "9999-12" || d == DateMax
}
func IsMinDatetime(d string) bool {
	return d == DatetimeMin || d == "1000-01-01 00:00:00" || IsMinDate(d)
}
func IsMaxDatetime(d string) bool {
	return d == DatetimeMax || IsMaxDate(d)
}
func ToDate(d string) (string, error) {
	if IsMinDate(d) {
		return DateMin, nil
	}
	if IsMaxDate(d) {
		return DateMax, nil
	}
	_, err := time.Parse(GoDateLayout, d)
	return d, err
}
func ToDatetime(d string) (string, error) {
	if IsMinDatetime(d) {
		return DatetimeMin, nil
	}
	if IsMaxDatetime(d) {
		return DatetimeMax, nil
	}
	_, err := time.Parse(GoDatetimeLayout, d)
	return d, err
}
