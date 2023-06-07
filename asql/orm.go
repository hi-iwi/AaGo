package asql

import (
	"strconv"
	"strings"
)

type Stmt struct {
	Cond_   string
	orderby string
	offset  uint
	limit   uint
}

func (o *Stmt) Concat(operator, field, asqlGrammar string) *Stmt {
	s := MakeASQL(asqlGrammar).Fmt(field)
	if o.Cond_ != "" {
		o.Cond_ += " " + operator + " "
	}
	o.Cond_ += s
	return o
}

func (o *Stmt) And(field, asqlGrammar string) *Stmt {
	return o.Concat("AND", field, asqlGrammar)
}

func (o *Stmt) Or(field, asqlGrammar string) *Stmt {
	return o.Concat("OR", field, asqlGrammar)
}

func (o *Stmt) OrderBy(keyword string) *Stmt {
	if o.orderby != "" {
		o.orderby += ","
	}
	o.orderby += keyword
	return o
}

func (o *Stmt) Limit(offset, limit uint) *Stmt {
	o.offset = offset
	o.limit = limit
	return o
}

func (o *Stmt) TryOrderBy(keyword string) *Stmt {
	if o.orderby == "" {
		o.orderby = keyword
	}
	return o
}

func (o *Stmt) TryLimit(offset, limit uint) *Stmt {
	if limit == 0 {
		o.offset = offset
		o.limit = limit
	}
	return o
}

func (o *Stmt) LimitStmt() string {
	if o.limit == 0 {
		if o.offset == 0 {
			o.limit = 20
		} else {
			o.limit = 10
		}
	}
	a := strconv.FormatUint(uint64(o.offset), 10)
	b := strconv.FormatUint(uint64(o.limit), 10)
	return "LIMIT " + a + "," + b
}

func (o *Stmt) OrderByStmt() string {
	return "ORDER BY " + o.orderby
}

func (o *Stmt) Stmt() string {
	var s strings.Builder
	if o.Cond_ != "" {
		s.WriteString(" WHERE ")
		s.WriteString(o.Cond_)
	}
	if o.orderby != "" {
		s.WriteString(" ORDER BY ")
		s.WriteString(o.orderby)
	}
	s.WriteByte(' ')
	s.WriteString(o.LimitStmt())
	return s.String()
}
