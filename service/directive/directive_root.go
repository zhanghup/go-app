package directive

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"xorm.io/xorm"
)

func Root(db *xorm.Engine) func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	return func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
		user := MyInfo(ctx)
		if *user.Info.User.Id != "root" {
			return nil, errors.New("只有root用户才能够访问")
		}
		return next(ctx)
	}
}
