package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/service/api/source"
)

func (r *mutationResolver) PlanStepCreate(ctx context.Context, input source.NewPlanStep) (string, error) {
	return r.Create(r.SessCtx(ctx), new(beans.PlanStep), input)
}

func (r *mutationResolver) PlanStepUpdate(ctx context.Context, id string, input source.UpdPlanStep) (bool, error) {
	return r.Update(r.SessCtx(ctx), new(beans.PlanStep), id, input)
}

func (r *mutationResolver) PlanStepRemoves(ctx context.Context, ids []string) (bool, error) {
	return r.Removes(r.SessCtx(ctx), new(beans.PlanStep), ids)
}

func (r *queryResolver) PlanSteps(ctx context.Context, query source.QPlanStep) (*source.PlanSteps, error) {
	plans := make([]beans.PlanStep, 0)
	i, err := r.DBS().SF(`
		select 
			* 
		from 
			plan_step p 
		where 1 = 1 
			{{ if .status }} and p.status = :status {{ end }}
		`,
		map[string]interface{}{
			"status": query.Status,
		}).Page2(query.Index, query.Size, query.Count, &plans)
	return &source.PlanSteps{Total: &i, Data: plans}, err
}

func (r *queryResolver) PlanStep(ctx context.Context, id string) (*beans.PlanStep, error) {
	return r.PlanStepLoader(ctx, id)
}
