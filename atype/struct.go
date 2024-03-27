package atype

import (
	"fmt"
	"log"
	"reflect"
)

// Tag Get tag of a struct.
// e.g.  struct { Iwi string `name:"iwi"`, Nation string `name:"nation"`}   atype.Tag(stru, "Iwi", "name")
func Tag(u any, field string, tagname string) string {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("field %s, tagname: %s error: %s\n", field, tagname, err)
		}
	}()

	t := reflect.TypeOf(u)
	if f, ok := t.FieldByName(field); ok {
		return f.Tag.Get(tagname)
	}

	return ""
}

func NameTag(u any, field string) string {
	return Tag(u, field, "name")
}

// ValueByTag Get value by its tag in a struct
func ValueByTag(u any, tagname string, tag string) (any, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	t := reflect.TypeOf(u)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		al := f.Tag.Get(tagname)
		if al == tag {
			return reflect.ValueOf(u).FieldByName(f.Name).Interface(), nil
		}
	}

	return nil, fmt.Errorf(`filed with tag %s:"%s" not found`, tagname, tag)
}

func ValueByName(u any, tag string) (any, error) {
	return ValueByTag(u, "name", tag)
}
