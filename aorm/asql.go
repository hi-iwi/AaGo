package aorm

import (
	"fmt"
	"strings"
)

type ASQL struct {
	EqualTo          string
	UnequalTo        string
	Like             string
	In               []string
	GreaterOrEqualTo string
	LessThan         string
	GreaterThan      string
	LessOrEqualTo    string
	StartWith        string
	EndWith          string
}

// urlencode special code is start with `%`, e.g. `%23`. Url special characters: + % / ? = % # &
// name=Iwi                                          name=Iwi
// name=::Iwi:                                       name likes Iwi
// name=::Iwi                                        name ends with Iwi
// name=:Iwi:                                        name starts Iwi
// name=:Iwi,Tom                                     name in [Iwi, Tom]
// create_at=2019-06-01 00:00:00                       create_at = 2019-06-01 00:00:00
// create_at=:2019-06-01 00:00:00~2019-06-01 01:00:00  create_at >= 2019-06-01 00:00:00 && create_at < 2019-06-01 00:00:00
// create_at=:2019-06-01 00:00:00~                     create_at >= 2019-06-01 00:00:00
// create_at=:~2019-06-01 01:00:00                     create_at < 2019-06-01 00:00:00

func MakeASQL(v string) ASQL {
	// a simple way to defense SQL injection
	v = defenseInjection(v)

	if len(v) > 2 && v[0] == ':' {
		v = v[1:]
		if v[0] == ':' {
			if v[len(v)-1] == ':' {
				return ASQL{Like: v[1 : len(v)-1]}
			}
			return ASQL{EndWith: v[1:]}
		} else if v[len(v)-1] == ':' {
			return ASQL{StartWith: v[:len(v)-1]}
		}

		if strings.Index(v, "~") > -1 {
			rng := strings.Split(v, "~")
			return ASQL{GreaterOrEqualTo: rng[0], LessThan: rng[1]}
		}

		if strings.Index(v, ",") > 0 {
			return ASQL{In: strings.Split(strings.Trim(v, ","), ",")}
		}
	}
	return ASQL{EqualTo: v}
}

func (q ASQL) Fmt(k string) string {
	if q.EqualTo != "" {
		return fmt.Sprintf(`%s="%s"`, toMySqlFieldName(k), q.EqualTo)
	} else if q.UnequalTo != "" {
		return fmt.Sprintf(`%s!="%s"`, toMySqlFieldName(k), q.EqualTo)
	} else if q.Like != "" {
		return fmt.Sprintf(`%s LIKE "%%%s%%"`, toMySqlFieldName(k), q.Like)
	} else if len(q.In) > 0 {
		r := ""
		ins := strings.Join(q.In, `","`)
		r += fmt.Sprintf(`%s IN ("%s")`, toMySqlFieldName(k), ins)
		return r
	} else if q.GreaterOrEqualTo != "" {
		if q.LessThan != "" {
			return fmt.Sprintf(`%s>="%s" AND %s<"%s"`, toMySqlFieldName(k), q.GreaterOrEqualTo, toMySqlFieldName(k), q.LessThan)
		} else if q.LessOrEqualTo != "" {
			return fmt.Sprintf(`%s>="%s" AND %s<="%s"`, toMySqlFieldName(k), q.GreaterOrEqualTo, toMySqlFieldName(k), q.LessOrEqualTo)
		}
		return fmt.Sprintf(`%s>="%s"`, toMySqlFieldName(k), q.GreaterOrEqualTo)
	} else if q.LessThan != "" {
		return fmt.Sprintf(`%s<"%s"`, toMySqlFieldName(k), q.LessThan)
	} else if q.GreaterThan != "" {
		if q.LessThan != "" {
			return fmt.Sprintf(`%s>"%s" AND %s<"%s"`, toMySqlFieldName(k), q.GreaterThan, toMySqlFieldName(k), q.LessThan)
		} else if q.LessOrEqualTo != "" {
			return fmt.Sprintf(`%s>"%s" AND %s<="%s"`, toMySqlFieldName(k), q.GreaterThan, toMySqlFieldName(k), q.LessOrEqualTo)
		}
		return fmt.Sprintf(`%s>"%s"`, toMySqlFieldName(k), q.GreaterThan)
	} else if q.LessOrEqualTo != "" {
		return fmt.Sprintf(`%s<="%s"`, toMySqlFieldName(k), q.LessOrEqualTo)
	} else if q.StartWith != "" {
		return fmt.Sprintf(`%s LIKE "%s%%"`, toMySqlFieldName(k), q.StartWith)
	} else if q.EndWith != "" {
		return fmt.Sprintf(`%s LIKE "%%%s"`, toMySqlFieldName(k), q.EndWith)
	}
	return ""
}
