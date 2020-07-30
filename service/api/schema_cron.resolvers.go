package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/zhanghup/go-app/service/job"

	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/service/api/lib"
)

func (r *mutationResolver) CronStop(ctx context.Context, id string) (bool, error) {
	err := job.Stop(id)
	return err == nil, err
}

func (r *mutationResolver) CronStart(ctx context.Context, id string) (bool, error) {
	err := job.CheckJob(id)
	if err != nil{
		return false,err
	}

	err = job.Start(id)
	return err == nil, err
}

func (r *mutationResolver) CronRun(ctx context.Context, id string) (bool, error) {
	err := job.CheckJob(id)
	if err != nil{
		return false,err
	}

	job.Run(id)()
	return true, nil
}

func (r *queryResolver) Crons(ctx context.Context, query lib.QCron) (*lib.Crons, error) {
	users := make([]beans.Cron, 0)
	total, err := r.DBS.SF(`
		select * from cron u
		where 1 = 1
		{{ if .keyword }} and u.name like concat("%",?keyword,"%") {{ end }}
	`, map[string]interface{}{"keyword": query.Keyword}).Page2(query.Index, query.Size, query.Count, &users)
	return &lib.Crons{Data: users, Total: &total}, err
}

func (r *queryResolver) Cron(ctx context.Context, id string) (*beans.Cron, error) {
	return r.CronLoader(ctx, id)
}
