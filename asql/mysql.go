package asql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/hi-iwi/AaGo/ae"
	"github.com/hi-iwi/AaGo/atype"
	"strconv"
	"strings"
)

const InArgs = "(...)" // InArgs 占位符
const MaxUint64IdLen = 20

func isInArgsTag(t interface{}) bool {
	v, ok := t.(string)
	return ok && v == InArgs
}
func uniqueKey(xargs []interface{}) string {
	var k strings.Builder
	for _, arg := range xargs {
		if !isInArgsTag(arg) {
			k.WriteString(atype.New(arg).String())
		}
	}
	return k.String()
}

// 一定要记得close rows，释放连接资源
func qsInUnion(ctx context.Context, db *DB, unionAll bool, format string, ids []uint64, demoTbLen int, args map[string][]interface{}, inArgs map[string][]uint64) (*sql.Rows, *ae.Error) {
	l := len(ids)
	var qs strings.Builder
	qs.Grow((len(format) + demoTbLen + l*(MaxUint64IdLen+1)) * l)
	var x bool
	union := " UNION "
	if unionAll {
		union += " ALL "
	}

	// sku 会分表，所以用union
	for k, arg := range args {
		if x {
			qs.WriteString(union)
		}
		x = true
		subids := inArgs[k]
		var s strings.Builder
		s.Grow(len(subids) * (MaxUint64IdLen + 1))
		for i, id := range subids {
			if i > 0 {
				s.WriteByte(',')
			}
			s.WriteString(strconv.FormatUint(id, 10))
		}
		for i, a := range arg {
			if isInArgsTag(a) {
				arg[i] = s.String()
			}
		}
		qs.WriteString(fmt.Sprintf(format, arg...))
	}
	// append LIMIT ?
	s := qs.String() + " LIMIT ?"
	return db.Query(ctx, s, l)
}

// 全表union all
// union 会过滤重复数据，性能稍差点；union all 不会过滤
// 尾部会自动添加 LIMIT ?
func InUnionAllTablesQs(ctx context.Context, db *DB, format string, ids []uint64, ptbs []string, xargs func(string, uint64) []interface{}) (*sql.Rows, *ae.Error) {
	args := make(map[string][]interface{}, 0)
	inArgs := make(map[string][]uint64, 0)
	var demoTable string
	for _, id := range ids {
		for _, ptb := range ptbs {
			arg := xargs(ptb, id)
			k := uniqueKey(arg)
			if _, ok := args[k]; ok {
				inArgs[k] = append(inArgs[k], id)
			} else {
				if demoTable == "" {
					demoTable = k
				}
				args[k] = arg
				inArgs[k] = []uint64{id}
			}
		}
	}
	return qsInUnion(ctx, db, true, format, ids, len(demoTable), args, inArgs)
}

// 处理按查询id分表的连表操作，不用全表union all
// union 会过滤重复数据，性能稍差点；union all 不会过滤
func InUnionAllQs(ctx context.Context, db *DB, format string, ids []uint64, xargs func(uint64) []interface{}) (*sql.Rows, *ae.Error) {
	args := make(map[string][]interface{}, 0)
	inArgs := make(map[string][]uint64, 0)
	var demoTable string
	for _, id := range ids {
		arg := xargs(id)
		k := uniqueKey(arg)
		if _, ok := args[k]; ok {
			inArgs[k] = append(inArgs[k], id)
		} else {
			if demoTable == "" {
				demoTable = k
			}
			args[k] = arg
			inArgs[k] = []uint64{id}
		}
	}
	return qsInUnion(ctx, db, true, format, ids, len(demoTable), args, inArgs)
}
