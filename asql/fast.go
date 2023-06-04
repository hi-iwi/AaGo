package asql

import (
	"github.com/hi-iwi/AaGo/atype"
	"strconv"
)

func In(field string, ids []uint64) string {
	if len(ids) == 0 {
		return "1!=1"
	}
	if len(ids) == 1 {
		return field + "=" + strconv.FormatUint(ids[0], 10)
	}
	return field + "=" + atype.JoinUint64(ids, ',')
}
func InUint(field string, ids []uint) string {
	if len(ids) == 0 {
		return "1!=1"
	}
	if len(ids) == 1 {
		return field + "=" + strconv.FormatUint(uint64(ids[0]), 10)
	}
	return field + "=" + atype.JoinUint(ids, ',')
}

/*
  组合sql语句，用于修改符合valid条件的字段
  @return ["a=?","b=?"], [$a,$b]
*/
func AppendArgs(field string, v interface{}, valid bool, ffs []string, fargs []interface{}, nosync []string) ([]string, []interface{}) {
	if !valid {
		return ffs, fargs
	}
	if nosync != nil {
		for _, no := range nosync {
			if no == field {
				return ffs, fargs
			}
		}
	}
	ffs = append(ffs, field+"=?")
	fargs = append(fargs, v)
	return ffs, fargs
}
