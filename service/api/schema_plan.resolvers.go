package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/service/api/source"
)

func (r *mutationResolver) PlanCreate(ctx context.Context, input source.NewPlan) (string, error) {
	user := r.Me(ctx).Info.User
	plan := beans.Plan{
		Puid:   user.Id,
		Puname: user.Name,
	}
	return r.Create(ctx, &plan, input)
}

func (r *mutationResolver) PlanUpdate(ctx context.Context, id string, input source.UpdPlan) (bool, error) {
	return r.Update(ctx, new(beans.Plan), id, input)
}

func (r *mutationResolver) PlanRemoves(ctx context.Context, ids []string) (bool, error) {
	return r.Removes(ctx, new(beans.Dept), ids)
}

func (r *queryResolver) Plans(ctx context.Context, query source.QPlan) (*source.Plans, error) {
	plans := make([]beans.Plan, 0)
	i, err := r.DBS().SF(`
		select 
			* 
		from 
			plan p 
		where 1 = 1 
			{{ if .status }} and p.status = :status {{ end }}
		`,
		map[string]interface{}{
			"status": query.Status,
		}).Page2(query.Index, query.Size, query.Count, &plans)
	return &source.Plans{Total: &i, Data: plans}, err
}

func (r *queryResolver) Plan(ctx context.Context, id string) (*beans.Plan, error) {
	return r.PlanLoader(ctx, id)
}
