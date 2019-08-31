//go:generate go run github.com/99designs/gqlgen

package api

import (
	"context"
	"github.com/go-xorm/xorm"
	"github.com/zhanghup/go-app/api/gs"
	"github.com/zhanghup/go-tools"
)

type Resolver struct {
	DB     func(ctx context.Context) *xorm.Session
	Loader func(ctx context.Context) gs.Loader
}

func (r *Resolver) Mutation() gs.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() gs.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (this mutationResolver) World(ctx context.Context) (*string, error) {
	return tools.Ptr().String("hello"), nil
}

type queryResolver struct{ *Resolver }

func (this queryResolver) Hello(ctx context.Context) (*string, error) {
	return tools.Ptr().String("world"), nil
}
