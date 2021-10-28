package dtype

import (
	"fmt"
	"log"
	"reflect"
)

// Tag Get tag of a struct.
// e.g.  struct { Iwi string `name:"iwi"`, Nation string `name:"nation"`}   dtype.Tag(stru, "Iwi", "name")
func Tag(u interface{}, field string, tagname string) string {
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

func NameTag(u interface{}, field string) string {
	return Tag(u, field, "name")
}

// ValueByTag Get value by its tag in a struct
func ValueByTag(u interface{}, tagname string, tag string) (interface{}, error) {
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

func ValueByName(u interface{}, tag string) (interface{}, error) {
	return ValueByTag(u, "name", tag)
}
