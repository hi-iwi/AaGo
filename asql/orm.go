package asql

import (
	"strconv"
	"strings"
)

type Stmt struct {
	Cond_    string
	OrderBy_ string
	Offset_  uint
	Limit_   uint
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
	if o.OrderBy_ != "" {
		o.OrderBy_ += ","
	}
	o.OrderBy_ += keyword
	return o
}

func (o *Stmt) Limit(offset, limit uint) *Stmt {
	o.Offset_ = offset
	o.Limit_ = limit
	return o
}

func (o *Stmt) TryOrderBy(keyword string) *Stmt {
	if o.OrderBy_ == "" {
		o.OrderBy_ = keyword
	}
	return o
}

func (o *Stmt) TryLimit(offset, limit uint) *Stmt {
	if limit > 0 {
		o.Offset_ = offset
		o.Limit_ = limit
	}
	return o
}

func (o *Stmt) LimitStmt() string {
	if o.Limit_ == 0 {
		o.Limit_ = 100
	}
	a := strconv.FormatUint(uint64(o.Offset_), 10)
	b := strconv.FormatUint(uint64(o.Limit_), 10)
	return "LIMIT " + a + "," + b
}

func (o *Stmt) Stmt() string {
	var s strings.Builder
	if o.Cond_ != "" {
		s.WriteString(" WHERE ")
		s.WriteString(o.Cond_)
	}
	if o.OrderBy_ != "" {
		s.WriteString(" ORDER BY ")
		s.WriteString(o.OrderBy_)
	}

	s.WriteByte(' ')
	s.WriteString(o.LimitStmt())
	return s.String()
}
