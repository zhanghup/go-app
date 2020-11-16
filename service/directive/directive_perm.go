package directive

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-tools"
	"strings"
	"xorm.io/xorm"
)

type PermObjects map[string]string
type Perms map[string][]string

func Perm(db *xorm.Engine) func(ctx context.Context, obj interface{}, next graphql.Resolver, entity string, perm string, remark *string) (res interface{}, err error) {
	return func(ctx context.Context, obj interface{}, next graphql.Resolver, entity string, perm string, remark *string) (res interface{}, err error) {
		user := MyInfo(ctx)
		lg := beans.OperateLog{
			Bean: beans.Bean{
				Id:     tools.Ptr.Uid(),
				Status: tools.Ptr.Int(1),
			},
			Type:  &entity,
			Opt:   &perm,
			Uid:   user.Info.User.Id,
			Uname: user.Info.User.Name,
		}

		// root 无限操作权限
		if user.Info.User.Id != nil && *user.Info.User.Id == "root" {
			res, err = next(ctx)
			if err != nil {
				lg.State = tools.Ptr.String("error") // 失败
				lg.Msg = tools.Ptr.String(err.Error())
			} else {
				lg.State = tools.Ptr.String("success") // 成功
			}
		} else {
			// 管理员
			if user.Info.Admin {
				res, err = next(ctx)
				if err != nil {
					lg.State = tools.Ptr.String("error") // 失败
					lg.Msg = tools.Ptr.String(err.Error())
				} else {
					lg.State = tools.Ptr.String("success") // 成功
				}
			} else {
				// 非管理员
				data, ok := user.Info.PermObjects[entity]
				if ok && strings.Contains(data, perm) {
					res, err = next(ctx)
					if err != nil {
						lg.State = tools.Ptr.String("error") // 失败
						lg.Msg = tools.Ptr.String(err.Error())
					} else {
						lg.State = tools.Ptr.String("success") // 成功
					}
				} else {
					lg.State = tools.Ptr.String("refuse") // 拒绝
				}
			}
		}

		input := graphql.GetOperationContext(ctx)
		lg.Gql = &input.RawQuery
		lg.GqlVariables = tools.Ptr.String(tools.Str.JSONString(input.Variables))
		go db.Insert(lg)
		return
	}
}
