package initia

import (
	"context"
	"github.com/zhanghup/go-app/gs"
	"github.com/zhanghup/go-app/service/directive"
	"github.com/zhanghup/go-tools"
)

func InitDBTemplate() {
	gs.InfoBegin("数据库模板")

	// 当前用户可以查看的用户列表
	InitDBTemplateFn("with_role_user", `
		{{ if .admin }}
			select id from user
		{{ else }}
			select distinct ruu.uid id from user
			join role_user ru on ru.uid = user.id and user.id = '{{ .uid }}' 
			join role_user ruu on ruu.role = ru.role
		{{ end }}
	`)

	// 当前用户可以查看的角色
	InitDBTemplateFn("with_role", `
		{{ if .admin }}
			select id from role
		{{ else }}
			select distinct ru.role id from user
			join role_user ru on ru.uid = user.id and user.id = '{{ .uid }}' 
		{{ end }}
	`)
}

func InitDBTemplateFn(name, sqlstr string) {
	gs.InfoSuccess(`数据库模板`, name)

	gs.DBA().TemplateFuncAdd(name, func(ctx context.Context) string {
		user := directive.MyInfo(ctx)
		return tools.StrTmp(sqlstr, map[string]interface{}{
			"uid":   user.Id,
			"admin": user.Admin,
		}).String()
	})
}
