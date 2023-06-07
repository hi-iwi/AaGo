package asql

import (
	"strconv"
	"strings"
)

type Cond struct {
	Constraint string
	orderby    string
	offset     uint
	limit      uint
}

func (c *Cond) Concat(operator, field, asqlGrammar string) *Cond {
	s := MakeASQL(asqlGrammar).Fmt(field)
	if c.Constraint != "" {
		c.Constraint += " " + operator + " "
	}
	c.Constraint += s
	return c
}

func (c *Cond) And(field, asqlGrammar string) *Cond {
	return c.Concat("AND", field, asqlGrammar)
}

func (c *Cond) Or(field, asqlGrammar string) *Cond {
	return c.Concat("OR", field, asqlGrammar)
}

func (c *Cond) OrderBy(keyword string) *Cond {
	if c.orderby != "" {
		c.orderby += ","
	}
	c.orderby += keyword
	return c
}

func (c *Cond) Limit(offset, limit uint) *Cond {
	c.offset = offset
	c.limit = limit
	return c
}

func (c *Cond) TryOrderBy(keyword string) *Cond {
	if c.orderby == "" {
		c.orderby = keyword
	}
	return c
}

func (c *Cond) TryLimit(offset, limit uint) *Cond {
	if limit == 0 {
		c.offset = offset
		c.limit = limit
	}
	return c
}

func (c *Cond) LimitStmt() string {
	if c.limit == 0 {
		if c.offset == 0 {
			c.limit = 20
		} else {
			c.limit = 10
		}
	}
	a := strconv.FormatUint(uint64(c.offset), 10)
	b := strconv.FormatUint(uint64(c.limit), 10)
	return "LIMIT " + a + "," + b
}

func (c *Cond) OrderByStmt() string {
	return "ORDER BY " + c.orderby
}

func (c *Cond) Stmt() string {
	var s strings.Builder
	if c.Constraint != "" {
		s.WriteString(" WHERE ")
		s.WriteString(c.Constraint)
	}
	if c.orderby != "" {
		s.WriteString(" ORDER BY ")
		s.WriteString(c.orderby)
	}
	s.WriteByte(' ')
	s.WriteString(c.LimitStmt())
	return s.String()
}
