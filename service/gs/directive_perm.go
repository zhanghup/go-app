package gs

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
)

type PermObjects map[string]string
type Perms map[string][]string

func Perm() func(ctx context.Context, obj interface{}, next graphql.Resolver, entity string, perm string) (res interface{}, err error) {
	return func(ctx context.Context, obj interface{}, next graphql.Resolver, entity string, perm string) (res interface{}, err error) {
		return next(ctx)
	}
}
