package aenum

import "strconv"

type Bank uint

const (
	BankICBC Bank = 1 // 中国工商银行
)

func (bank Bank) String() string {
	return strconv.FormatInt(int64(bank), 10)
}
func (bank Bank) Stringify() string {
	switch bank {
	case BankICBC:
		return "ICBC"
	}
	return ""
}
