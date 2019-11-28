package directive

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"strings"
)

type PermObjects map[string]string
type Perms map[string][]string

func Perm() func(ctx context.Context, obj interface{}, next graphql.Resolver, entity string, perm string) (res interface{}, err error) {
	return func(ctx context.Context, obj interface{}, next graphql.Resolver, entity string, perm string) (res interface{}, err error) {
		md := MewMe(ctx)
		user := md.User()

		// root 无线操作权限
		if user.Id != nil && *user.Id == "root" {
			return next(ctx)
		}

		if !md.Admin() {
			data, ok := md.PermObjs()[entity]
			if !ok {
				return nil, errors.New("【无操作权限 1】")
			}
			if strings.Contains(data, perm) {
				return nil, errors.New("【无操作权限 2】")
			}
		}

		return next(ctx)
	}
}
