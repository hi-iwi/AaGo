package asql

import (
	"github.com/hi-iwi/AaGo/atype"
	"strconv"
	"strings"
)

type Cond struct {
	Constraint strings.Builder
	orderby    string
	offset     uint
	limit      uint
}

func NewCond(paging atype.Paging) *Cond {
	return &Cond{offset: paging.Offset, limit: paging.Limit}
}
func (c *Cond) WriteString(s string) *Cond {
	c.Constraint.WriteByte(' ')
	c.Constraint.WriteString(s)
	return c
}
func (c *Cond) Write(operator, s string) *Cond {
	if c.Constraint.Len() == 0 {
		operator = ""
	} else {
		operator += " "
	}
	return c.WriteString(operator + s)
}
func (c *Cond) Concat(operator, field, asqlGrammar string) *Cond {
	s := MakeASQL(asqlGrammar).Fmt(field)
	if c.Constraint.Len() > 0 {
		c.Constraint.WriteByte(' ')
		c.Constraint.WriteString(operator)
		c.Constraint.WriteByte(' ')
	}
	c.Constraint.WriteString(s)
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

func (c *Cond) Try(orderBy string, offset, limit uint) *Cond {
	if c.orderby == "" {
		c.orderby = orderBy
	}
	if limit == 0 {
		c.offset = offset
		c.limit = limit
	}
	return c
}

func (c *Cond) LimitN() uint {
	if c.limit > 0 {
		return c.limit
	}
	return 10
}

func (c *Cond) LimitStmt() string {
	limit := c.LimitN()
	a := strconv.FormatUint(uint64(c.offset), 10)
	b := strconv.FormatUint(uint64(limit), 10)
	return "LIMIT " + a + "," + b
}

func (c *Cond) OrderByStmt() string {
	return "ORDER BY " + c.orderby
}

func (c *Cond) Stmt() string {
	var s strings.Builder
	if c.Constraint.Len() > 0 {
		s.WriteString(" WHERE ")
		s.WriteString(c.Constraint.String())
	}
	if c.orderby != "" {
		s.WriteString(" ORDER BY ")
		s.WriteString(c.orderby)
	}
	s.WriteByte(' ')
	s.WriteString(c.LimitStmt())
	return s.String()
}
