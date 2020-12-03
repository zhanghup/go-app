package initia

import (
	"context"
	"github.com/zhanghup/go-app/service/directive"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
	"xorm.io/xorm"
)

func InitDBTemplate(db *xorm.Engine) {
	dbs := txorm.NewEngine(db)

	dbs.TemplateFuncAdd("users",InitDBTemplateFn(`
	`))
}

func InitDBTemplateFn(sqlstr string) func(ctx context.Context) string {
	return func(ctx context.Context) string {
		user := directive.MyInfo(ctx).Info.User
		return tools.Str.Tmp(sqlstr, map[string]interface{}{"id": user.Id}).String()
	}
}
