package directive

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/zhanghup/go-app/service/ca"
)

type PermObjects map[string]string
type Perms map[string][]string

func Perm() func(ctx context.Context, obj interface{}, next graphql.Resolver, entity string, perm string, remark *string) (res interface{}, err error) {
	return func(ctx context.Context, obj interface{}, next graphql.Resolver, entity string, perm string, remark *string) (res interface{}, err error) {
		a, b, c := ca.DictCache.Get("SYS0002")
		fmt.Println(a, b, c)
		//md := MyInfo(ctx)
		//user := md.User()
		//
		//// root 无限操作权限
		//if user.Id != nil && *user.Id == "root" {
		//	res, err = next(ctx)
		//}
		//
		//state := 1
		//if !md.Admin() {
		//	data, ok := md.PermObjs()[entity]
		//	if !ok {
		//		state = -1
		//	}
		//	if strings.Contains(data, perm) {
		//		state = -1
		//	}
		//}
		//
		//if state == 1 {
		//	res, err = next(ctx)
		//}
		//
		//go tools.Run(func() {
		//	dc := DictCacheInfo()
		//	_, dis, ok := dc.Get("SYS0002")
		//	var msg string
		//	for _, o := range dis {
		//		if o.Value == nil || o.Name == nil {
		//			continue
		//		}
		//		if perm == *o.Value {
		//			msg = *o.Name
		//			break
		//		}
		//	}
		//	fmt.Println(msg)
		//	if ok {
		//		model := beans.OperateLog{
		//			Bean: beans.Bean{
		//				Id:     tools.Ptr.Uid(),
		//				Status: &state,
		//			},
		//			Type: &entity,
		//			Opt:  &perm,
		//			Uid:  user.Id,
		//		}
		//		if err != nil {
		//			model.State = tools.Ptr.Int(0)
		//			model.Msg = tools.Ptr.String(err.Error())
		//		} else {
		//
		//		}
		//
		//
		//	}
		//})

		return next(ctx)
	}
}
