package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/zhanghup/go-app/service/api/lib"
	"github.com/zhanghup/go-tools"
	"time"
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

// Mutation returns lib.MutationResolver implementation.
func (r *Resolver) Mutation() lib.MutationResolver { return &mutationResolver{r} }

// Query returns lib.QueryResolver implementation.
func (r *Resolver) Query() lib.QueryResolver { return &queryResolver{r} }

// Subscription returns lib.SubscriptionResolver implementation.
func (r *Resolver) Subscription() lib.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
