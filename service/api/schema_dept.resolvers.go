package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/service/api/source"
)

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
func (r *queryResolver) DeptTree__(items []beans.Dept, pid string, flag ...bool) interface{} {
	type DeptTreeItem struct {
		Id       *string     `json:"id"`
		Name     *string     `json:"name"`
		Code     *string     `json:"code"`
		Children interface{} `json:"children"`
	}

	results := make([]DeptTreeItem, 0)
	for _, o := range items {
		item := DeptTreeItem{
			Id:   o.Id,
			Name: o.Name,
			Code: o.Code,
		}

		if len(flag) > 0 && flag[0] {
			if o.Pid == nil {
				item.Children = r.DeptTree__(items, *o.Id)
				results = append(results, item)
			}
		} else {
			if *o.Pid == pid {
				item.Children = r.DeptTree__(items, *o.Id)
				results = append(results, item)
			}
		}
	}
	return results
}
func (r *queryResolver) DeptTree(ctx context.Context) (interface{}, error) {
	depts := make([]beans.Dept, 0)
	err := r.DB.Find(&depts)
	if err != nil {
		return nil, err
	}
	return r.DeptTree__(depts, "", true), err
}
