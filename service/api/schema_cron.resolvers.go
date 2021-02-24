package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/zhanghup/go-app/gs"

	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/service/api/source"
	"github.com/zhanghup/go-app/service/job"
)

func (r *mutationResolver) CronStop(ctx context.Context, id string) (bool, error) {
	err := job.Stop(id)
	return err == nil, err
}

func (r *mutationResolver) CronStart(ctx context.Context, id string) (bool, error) {
	err := job.CheckJob(id)
	if err != nil {
		return false, err
	}

	err = job.Start(id)
	return err == nil, err
}

func (r *mutationResolver) CronRun(ctx context.Context, id string) (bool, error) {
	err := job.CheckJob(id)
	if err != nil {
		return false, err
	}

	job.Run(id)()
	return true, nil
}

func (r *queryResolver) Crons(ctx context.Context, query source.QCron) (*source.Crons, error) {
	users := make([]beans.Cron, 0)
	total, err := gs.DBS().SF(`
		select * from cron u
		where 1 = 1
		{{ if .keyword }} and u.name like concat("%",:keyword,"%") {{ end }}
		{{ if .state }} and u.state = :state {{ end }}
		{{ if .name }} and u.name like concat("%",:name,"%") {{ end }}
		{{ if .result }} and u.result = :result {{ end }}
	`, map[string]interface{}{
		"keyword": query.Keyword,
		"state":   query.State,
		"name":    query.Name,
		"result":  query.Result,
	}).Page2(query.Index, query.Size, query.Count, &users)
	return &source.Crons{Data: users, Total: &total}, err
}

func (r *queryResolver) Cron(ctx context.Context, id string) (*beans.Cron, error) {
	return r.CronLoader(ctx, id)
}

func (r *queryResolver) CronLogs(ctx context.Context, query source.QCronLog) (*source.CronLogs, error) {
	users := make([]beans.CronLog, 0)
	total, err := gs.DBS().SF(`
		select * from cron_log u
		where 1 = 1
		{{ if .keyword }} and u.name like concat("%",:keyword,"%") {{ end }}
		{{ if .cron }} and u.cron = :cron {{ end }}
		order by u.created desc
	`, map[string]interface{}{"keyword": query.Keyword, "cron": query.Cron}).Page2(query.Index, query.Size, query.Count, &users)
	return &source.CronLogs{Data: users, Total: &total}, err
}
