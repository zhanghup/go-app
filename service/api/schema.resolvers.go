package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/zhanghup/go-app/service/api/source"
	"github.com/zhanghup/go-tools"
)

func (r *mutationResolver) World(ctx context.Context) (*string, error) {
	return tools.Ptr.String("world"), nil
}

func (r *queryResolver) Stat(ctx context.Context) (interface{}, error) {
	return statsReport(), nil
}

func (r *queryResolver) Hello(ctx context.Context) (*string, error) {
	return tools.Ptr.String("hello"), nil
}

func (r *subscriptionResolver) Hello(ctx context.Context) (<-chan *string, error) {
	c := make(chan *string, 10)
	go tools.RunWithContext(time.Second, ctx, func() {
		c <- tools.Ptr.String("hello world")
	})
	return c, nil
}

// Mutation returns source.MutationResolver implementation.
func (r *Resolver) Mutation() source.MutationResolver { return &mutationResolver{r} }

// Query returns source.QueryResolver implementation.
func (r *Resolver) Query() source.QueryResolver { return &queryResolver{r} }

// Subscription returns source.SubscriptionResolver implementation.
func (r *Resolver) Subscription() source.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
