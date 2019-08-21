//go:generate go run github.com/99designs/gqlgen

package api

import (
	"context"
	"github.com/go-xorm/xorm"
	"github.com/zhanghup/go-app/api/gs"
)

type Resolver struct {
	DB xorm.Engine
}

func (r *Resolver) Mutation() gs.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() gs.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (this mutationResolver) World(ctx context.Context) (*string, error) {
	a := "hello"
	return &a, nil
}

type queryResolver struct{ *Resolver }

func (this queryResolver) Hello(ctx context.Context) (*string, error) {
	a := "world"
	return &a, nil
}
