package dtype

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"sort"
	"strings"
)

type JoinType int

const (
	JoinSortedBit JoinType = 1 << 20
)

const (
	JoinKeys                  JoinType = 1 << iota // k
	JoinValues                                     // v
	JoinMySqlValues                                // "v"
	JoinKV                                         //kv
	JoinJSON                                       // "k":"v"
	JoinMySQL                                      // `t`.`k`="v"
	JoinMySqlFullLike                              // `t`.`k` LIKE "%v%"
	JoinMySqlStartWith                             // `t`.`k` LIKE "v%"
	JoinMySqlEndWith                               // `t`.`k` LIKE "%v"
	JoinMySqlLessThan                              // `t`.`k` < "v"
	JoinMySqlGreaterThan                           // `t`.`k` > "v"
	JoinMySqlGreaterOrEqualTo                      // `t`.`k` >= "v"
	JoinMySqlLessOrEqualTo                         // `t`.`k` <= "v"
	JoinURL                                        // k=v
	JoinSortedValues          = JoinSortedBit | JoinValues
	JoinSortedMySqlValues     = JoinSortedBit | JoinMySqlValues
	JoinSortedKV              = JoinSortedBit | JoinKV
	JoinSortedJSON            = JoinSortedBit | JoinJSON
	JoinSortedMySQL           = JoinSortedBit | JoinMySQL
	JoinSortedURL             = JoinSortedBit | JoinURL
)

func toMySqlFieldName(k string) string {
	fields := strings.Split(k, ".")
	for i, field := range fields {
		fields[i] = "`" + field + "`"
	}
	return strings.Join(fields, ".")
}

func byJoinType(ty JoinType, k string, v interface{}) string {
	var val string
	// @TODO separate sql from here
	if w, ok := v.(sql.NullBool); ok {
		val = String(w.Bool)
	} else if w, ok := v.(sql.NullFloat64); ok {
		val = String(w.Float64)
	} else if w, ok := v.(sql.NullInt64); ok {
		val = String(w.Int64)
	} else if w, ok := v.(sql.NullString); ok {
		val = w.String
	} else {
		val = String(v)
	}

	switch ty {
	case JoinKeys:
		return k
	case JoinSortedValues, JoinValues:
		return val
	case JoinSortedMySqlValues, JoinMySqlValues:
		return fmt.Sprintf(`"%s"`, val)
	case JoinSortedKV, JoinKV:
		return k + val
	case JoinSortedJSON, JoinJSON:
		return fmt.Sprintf(`"%s":"%s"`, k, val)
	case JoinSortedMySQL, JoinMySQL:
		return fmt.Sprintf(`%s="%s"`, toMySqlFieldName(k), val)
	case JoinMySqlFullLike:
		return fmt.Sprintf(`%s LIKE "%%%s%%"`, toMySqlFieldName(k), val)
	case JoinMySqlStartWith:
		return fmt.Sprintf(`%s LIKE "%s%%"`, toMySqlFieldName(k), val)
	case JoinMySqlEndWith:
		return fmt.Sprintf(`%s LIKE "%%%s"`, toMySqlFieldName(k), val)
	case JoinMySqlLessThan:
		return fmt.Sprintf(`%s<"%s"`, toMySqlFieldName(k), val)
	case JoinMySqlGreaterThan:
		return fmt.Sprintf(`%s>"%s"`, toMySqlFieldName(k), val)
	case JoinMySqlGreaterOrEqualTo:
		return fmt.Sprintf(`%s>="%s"`, toMySqlFieldName(k), val)
	case JoinMySqlLessOrEqualTo:
		return fmt.Sprintf(`%s<="%s"`, toMySqlFieldName(k), val)
	case JoinSortedURL, JoinURL:
		return fmt.Sprintf(`%s=%s`, k, val)
	}
	return ""
}

// JoinTagsByElements(stru, JoinUnsortedBit, " AND ", "json", "Name", "Age")
func JoinTagsByElements(u interface{}, ty JoinType, sep string, tagname string, eles ...string) (ret string) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	tags := make([]string, len(eles))

	t := reflect.TypeOf(u)
	for i, ele := range eles {
		for j := 0; j < t.NumField(); j++ {
			f := t.Field(j)
			if f.Name == ele {
				tags[i] = f.Tag.Get(tagname)
			}
		}
	}
	return JoinByTags(u, ty, sep, tagname, tags...)
}

// JoinByTags(stru, JoinUnsortedBit, " AND ", "json", "name", "age")
func JoinByTags(u interface{}, ty JoinType, sep string, tagname string, tags ...string) (ret string) {

	defer func() {
		if err := recover(); err != nil {
			log.Printf("[error] dtype.JoinByTags: %s", err)
		}
	}()

	if ty&JoinSortedBit > 0 {
		sort.Strings(tags)
	}

	t := reflect.TypeOf(u)
	var found bool
	for g := 0; g < len(tags); g++ {
		tag := tags[g]
		found = false
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			al := f.Tag.Get(tagname)
			if al == tag {
				found = true
				ret += sep + byJoinType(ty, tag, reflect.ValueOf(u).FieldByName(f.Name).Interface())
			}
		}
		if !found {
			panic(fmt.Sprintf(`not found %s:"%s"`, tagname, tag))
		}
	}
	if len(ret) > len(sep) {
		ret = ret[len(sep):]
	}
	return
}

func JoinByAlias(u interface{}, ty JoinType, sep string, alias ...string) string {
	return JoinByTags(u, ty, sep, "alias", alias...)
}

func JoinAliasByElements(u interface{}, ty JoinType, sep string, eles ...string) string {
	return JoinTagsByElements(u, ty, sep, "alias", eles...)
}
