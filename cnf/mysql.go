package cnf

import (
	"fmt"
)

type Mysql struct {
	id        string
	Scheme    string
	Authority string
	Username  string
	Password  string
	Db        string
}

func (s Mysql) Id() string {
	if s.id == "" {
		if s.Scheme != "" {
			s.id += s.Scheme + "://"
		}
		if s.Username != "" {
			s.id += s.Username + "@"
		}

		s.id += s.Authority

		if s.Password != "" {
			s.id += ":" + s.Password
		}
	}
	return s.id
}

/*
	"user:password@/dbname"
*/
func (s Mysql) DatasrcName() string {
	return fmt.Sprintf("%s:%s@%s(%s)/%s", s.Username, s.Password, s.Scheme, s.Authority, s.Db)
}
