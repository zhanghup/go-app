package initia

import (
	"context"
	"github.com/zhanghup/go-app/service/directive"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
	"github.com/zhanghup/go-tools/tog"
	"xorm.io/xorm"
)

func InitDBTemplate(db *xorm.Engine) {
	tog.Info("初始化数据库模板。。。")

	dbs := txorm.NewEngine(db)

	// 当前用户可以查看的用户列表
	dbs.TemplateFuncAdd("with_role_user", InitDBTemplateFn(`
		{{ if .admin }}
			select id from user
		{{ else }}
			select distinct ruu.uid id from user
			join role_user ru on ru.uid = user.id and user.id = '{{ .uid }}' 
			join role_user ruu on ruu.role = ru.role
		{{ end }}
	`))

	// 当前用户可以查看的角色
	dbs.TemplateFuncAdd("with_role", InitDBTemplateFn(`
		{{ if .admin }}
			select id from role
		{{ else }}
			select distinct ru.role id from user
			join role_user ru on ru.uid = user.id and user.id = '{{ .uid }}' 
		{{ end }}
	`))
}

func InitDBTemplateFn(sqlstr string) func(ctx context.Context) string {
	return func(ctx context.Context) string {
		user := directive.MyInfo(ctx)
		return tools.StrTmp(sqlstr, map[string]interface{}{
			"uid":   user.Id,
			"admin": user.Admin,
		}).String()
	}
}
