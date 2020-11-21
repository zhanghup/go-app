package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/service/api/source"
)

func (r *deptResolver) ODept(ctx context.Context, obj *beans.Dept) (*beans.Dept, error) {
	if obj.Pid == nil {
		return nil, nil
	}
	return r.Resolver.DeptLoader(ctx, *obj.Pid)
}

func (r *mutationResolver) DeptCreate(ctx context.Context, input source.NewDept) (string, error) {
	return r.Create(ctx, new(beans.Dept), input)
}

func (r *mutationResolver) DeptUpdate(ctx context.Context, id string, input source.UpdDept) (bool, error) {
	return r.Update(ctx, new(beans.Dept), id, input)
}

func (r *mutationResolver) DeptRemoves(ctx context.Context, ids []string) (bool, error) {
	return r.Removes(ctx, new(beans.Dept), ids)
}

func (r *queryResolver) Depts(ctx context.Context, query source.QDept) (*source.Depts, error) {
	depts := make([]beans.Dept, 0)

	i, err := r.DBS.SF(`
		select 
			* 
		from 
			dept d 
		where 1 = 1 
			{{ if .pid }} and d.pid = :pid {{ end }} 
		order by d.weight
		`,
		map[string]interface{}{
			"pid": query.Pid,
		}).Page2(query.Index, query.Size, query.Count, &depts)
	return &source.Depts{Total: &i, Data: depts}, err
}

func (r *queryResolver) Dept(ctx context.Context, id string) (*beans.Dept, error) {
	return r.Resolver.DeptLoader(ctx, id)
}

func (r *queryResolver) DeptTree(ctx context.Context) (interface{}, error) {
	depts := make([]beans.Dept, 0)
	err := r.DB.Find(&depts)
	if err != nil {
		return nil, err
	}
	return r.DeptTreeHelp(depts, "", true), err
}

// Dept returns source.DeptResolver implementation.
func (r *Resolver) Dept() source.DeptResolver { return &deptResolver{r} }

type deptResolver struct{ *Resolver }
