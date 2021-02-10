package awxmp

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/zhanghup/go-app/service/awxmp/source"
)

func (r *mutationResolver) World(ctx context.Context) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Hello(ctx context.Context) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) Hello(ctx context.Context) (<-chan *string, error) {
	panic(fmt.Errorf("not implemented"))
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
