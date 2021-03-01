package ags

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	source1 "github.com/zhanghup/go-app/service/ags/source"
)

func (r *mutationResolver) Hello(ctx context.Context) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Hello(ctx context.Context) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) Hello(ctx context.Context) (<-chan *string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns source1.MutationResolver implementation.
func (r *Resolver) Mutation() source1.MutationResolver { return &mutationResolver{r} }

// Query returns source1.QueryResolver implementation.
func (r *Resolver) Query() source1.QueryResolver { return &queryResolver{r} }

// Subscription returns source1.SubscriptionResolver implementation.
func (r *Resolver) Subscription() source1.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
