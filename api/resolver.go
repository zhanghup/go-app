//go:generate go run github.com/99designs/gqlgen

package api

import (
	"context"
	"github.com/zhanghup/go-app"
	"github.com/zhanghup/go-app/api/gs"
)

type Resolver struct{}

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

func (this queryResolver) Users(ctx context.Context) ([]*app.User, error) {
	return nil, nil
}

func (this queryResolver) Hello(ctx context.Context) (*string, error) {
	a := "world"
	return &a, nil
}
