package aenum

const AaBirthDate = int64(690134400) // 秒   1991-11-15 00:00:00

// https://dev.mysql.com/doc/refman/8.0/en/datetime.html
// DATETIME values is '1000-01-01 00:00:00.000000' to '9999-12-31 23:59:59.999999',
// TIMESTAMP values is '1970-01-01 00:00:01.000000' to '2038-01-19 03:14:07.999999'
// 两个时间最小、最大均不同，所以一定要开启 sql-mode  NO_ZERO_DATE，保持最小为 0000-00-00 00:00:00
const DateMin = "0000-00-00"
const DatetimeMin = "0000-00-00 00:00:00"
const DateMax = "9999-12-31"
const DatetimeMax = "9999-12-31 23:59:59"

const GoDateLayout = "2006-01-02"
const GoDatetimeLayout = "2006-01-02 15:04:05"

func ToDatetime(d string) string {
	if d == "" || d[0] == '0' || d[0:7] == "1000-00" || d == "1970-01-01" || d == "1970-01-01 00:00:01" {
		return DatetimeMin
	}
	return d
}
