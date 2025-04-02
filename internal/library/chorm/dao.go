package chorm

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type daoInstance interface {
	Table() string
	Ctx(ctx context.Context) *gdb.Model
}

// GetPkField 获取dao实例中的主键名称
func GetPkField(ctx context.Context, dao daoInstance) (string, error) {
	fields, err := dao.Ctx(ctx).TableFields(dao.Table())
	if err != nil {
		return "", err
	}
	if len(fields) == 0 {
		return "", gerror.New("field not found")
	}

	for _, field := range fields {
		if field.Key == "PRI" {
			return field.Name, nil
		}
	}
	return "", gerror.New("no primary key")
}

// IsUnique 验证字段是否唯一
func IsUnique(ctx context.Context, dao daoInstance, where g.Map, message string, pkId ...interface{}) error {
	if len(where) == 0 {
		return gerror.New("where condition cannot be empty")
	}

	m := dao.Ctx(ctx).Where(where)
	if len(pkId) > 0 {
		field, err := GetPkField(ctx, dao)
		if err != nil {
			return err
		}
		m = m.WhereNot(field, pkId[0])
	}

	count, err := m.Count()
	if err != nil {
		return err
	}

	if count > 0 {
		if message == "" {
			for k := range where {
				message = fmt.Sprintf("in the table：%s, %v not uniqued", dao.Table(), where[k])
				break
			}
		}
		return gerror.New(message)
	}
	return nil
}
