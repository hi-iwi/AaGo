package asql

import "regexp"

// 应当禁止 DELETE|DROP|TRUNCATE|RENAME|ALTER 等权限
func defenseInjection(v string) string {
	if len(v) < 13 {
		return v
	}
	v = regexp.MustCompile(`(?i)(INSERT|REPLACE)(LOW_PRIORITY|DELAYED|HIGH_PRIORITY|\s|IGNORE)*\sINTO\s`).ReplaceAllLiteralString(v, "")
	v = regexp.MustCompile(`(?i)UPDATE(LOW_PRIORITY|\s|IGNORE)*\sSET\s`).ReplaceAllLiteralString(v, "")
	// delete T1,T2 FROM T1 INNER JOIN... or delete from t1
	// DELETE t1.field1,t1.field2 FROM t1
	// 应当禁止 delete 权限
	//v = regexp.MustCompile(`(?i)DELETE(LOW_PRIORITY|\s|QUICK|IGNORE)*[\s\w,.\*]*\sFROM\s`).ReplaceAllLiteralString(v, "")
	//v = regexp.MustCompile(`(?i)(DROP|TRUNCATE|RENAME|ALTER)\s+TABLE\s`).ReplaceAllLiteralString(v, "")
	v = regexp.MustCompile(`(?i)</?script>?`).ReplaceAllLiteralString(v, "")
	return v
}
