package asql

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
