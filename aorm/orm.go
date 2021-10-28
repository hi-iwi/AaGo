package aorm

import (
	"log"
	"reflect"
)

type Limit struct {
	Offset int `name:"offset"`
	Limit  int `name:"limit"`
}

type CombineOperator string

const (
	OrCombineOperator  = ") OR ("
	AndCombineOperator = ") AND ("
)

type ORM struct {
	CombineOperator CombineOperator
	Limit           *Limit
	OrderBy         string
	OrderByDesc     string
	AndFields       []string
	OrFields        []string
	LikeFields      []string

	Index []string
	And   map[string]ASQL
	Or    map[string]ASQL
}

// name u struct; ele name name
func name(u interface{}, ele string) string {
	if !(ele[0] >= 'A' && ele[0] <= 'Z') {
		return ele
	}

	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	t := reflect.TypeOf(u)
	for j := 0; j < t.NumField(); j++ {
		f := t.Field(j)
		if f.Name == ele {
			return f.Tag.Get("name")
		}
	}
	return ele
}

func indexes(idx []string, asqls map[string]ASQL) []string {
	if len(idx) == 0 {
		idx = make([]string, 0)
	}
	newIdx := idx
	for k, _ := range asqls {
		ins := false
		for i := 0; i < len(idx); i++ {
			if idx[i] == k {
				ins = true
				break
			}
		}
		if !ins {
			newIdx = append(newIdx, k)
		}
	}
	return newIdx
}

func (orm ORM) WithWhere(t interface{}) string {
	ands := ""
	ors := ""

	idx := indexes(orm.Index, orm.And)

	for i := 0; i < len(idx); i++ {
		for k, a := range orm.And {
			if idx[i] != k {
				continue
			}
			if f := a.Fmt(name(t, k)); f != "" {
				ands += " AND " + f
			}
		}
	}

	idx = indexes(orm.Index, orm.Or)
	for i := 0; i < len(idx); i++ {
		for k, a := range orm.Or {
			if idx[i] != k {
				continue
			}
			if f := a.Fmt(name(t, k)); f != "" {
				ors += " OR " + f
			}
		}
	}
	operator := string(orm.CombineOperator)
	return Where("(", And(t, orm.AndFields...), ands, operator, Or(t, orm.OrFields...), ors, operator, Like(t, orm.LikeFields...), ")")
}
