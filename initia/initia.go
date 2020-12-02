package initia

import (
	"context"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
	"xorm.io/xorm"
)

func InitDBTemplate(db *xorm.Engine) {
	dbs := txorm.NewEngine(db)
	dbs.TemplateFuncAdd("users", func(ctx context.Context) string {
		return tools.Str.Tmp(`
			users as (select id from user)
		`, map[string]interface{}{}).String()
	})
}
