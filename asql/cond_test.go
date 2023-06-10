package asql_test

import (
	"github.com/hi-iwi/AaGo/asql"
	"testing"
)

func TestCond(t *testing.T) {
	var cond = &asql.Cond{}
	cond.And("t.id", "100")
	if cond.Stmt() != " WHERE `t`.`id`=\"100\" LIMIT 0,20" {
		t.Errorf("test cond failed `%s`", cond.Stmt())
	}
}
