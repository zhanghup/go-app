package directive

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-tools"
	"xorm.io/xorm"
)

type PermObjects map[string]string
type Perms map[string][]string

func Perm(db *xorm.Engine) func(ctx context.Context, obj interface{}, next graphql.Resolver, entity string, perm string, remark *string) (res interface{}, err error) {
	return func(ctx context.Context, obj interface{}, next graphql.Resolver, entity string, perm string, remark *string) (res interface{}, err error) {
		user := MyInfo(ctx)
		lg := beans.OperateLog{
			Bean: beans.Bean{
				Id:     tools.PtrOfUUID(),
				Status: tools.PtrOfString("1"),
			},
			Type:  &entity,
			Opt:   &perm,
			Uid:   &user.Id,
			Uname: &user.Name,
		}

		// root 无限操作权限
		if user.Id == "root" {
			res, err = next(ctx)
			if err != nil {
				lg.State = tools.PtrOfString("error") // 失败
				lg.Msg = tools.PtrOfString(err.Error())
			} else {
				lg.State = tools.PtrOfString("success") // 成功
			}
		} else {
			// 管理员
			if user.Admin {
				res, err = next(ctx)
				if err != nil {
					lg.State = tools.PtrOfString("error") // 失败
					lg.Msg = tools.PtrOfString(err.Error())
				} else {
					lg.State = tools.PtrOfString("success") // 成功
				}
			} else {
				// 非管理员
				//data, ok := user.Info.PermObjects[entity]
				//if ok && strings.Contains(data, perm) {
				res, err = next(ctx)
				if err != nil {
					lg.State = tools.PtrOfString("error") // 失败
					lg.Msg = tools.PtrOfString(err.Error())
				} else {
					lg.State = tools.PtrOfString("success") // 成功
				}
				//} else {
				//	lg.State = tools.Ptr.String("refuse") // 拒绝
				//}
			}
		}

		input := graphql.GetOperationContext(ctx)
		lg.Gql = &input.RawQuery
		lg.GqlVariables = tools.PtrOfString(tools.JSONString(input.Variables))
		go db.Insert(lg)

		return
	}
}
