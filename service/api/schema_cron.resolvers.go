package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/service/api/lib"
)

func (r *mutationResolver) CronStop(ctx context.Context, id *string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CronStart(ctx context.Context, id *string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CronRun(ctx context.Context, id *string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Crons(ctx context.Context, query lib.QCron) (*lib.Crons, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Cron(ctx context.Context, id string) (*beans.Cron, error) {
	panic(fmt.Errorf("not implemented"))
}
