package gs

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
)

func Perm() func(ctx context.Context, obj interface{}, next graphql.Resolver, entity string, perm string) (res interface{}, err error) {
	return func(ctx context.Context, obj interface{}, next graphql.Resolver, entity string, perm string) (res interface{}, err error) {
		fmt.Println(obj)
		return nil, nil
	}
}
