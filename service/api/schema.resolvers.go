package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/zhanghup/go-app/service/api/lib"
)

func (r *mutationResolver) World(ctx context.Context) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Stat(ctx context.Context) (interface{}, error) {
	return statsReport(), nil
}

func (r *queryResolver) Hello(ctx context.Context) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns lib.MutationResolver implementation.
func (r *Resolver) Mutation() lib.MutationResolver { return &mutationResolver{r} }

// Query returns lib.QueryResolver implementation.
func (r *Resolver) Query() lib.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
