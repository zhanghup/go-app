package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/service/api/source"
	"github.com/zhanghup/go-app/service/event"
)

func (r *dictResolver) Values(ctx context.Context, obj *beans.Dict) ([]beans.DictItem, error) {
	if obj.Id == nil {
		return nil, nil
	}
	c := make([]beans.DictItem, 0)
	err := r.Loader(ctx).Slice(c, "select * from dict_item where code in :keys order by weight", nil, "Code", "").Load(*obj.Code, &c)
	return c, err
}

func (r *mutationResolver) DictCreate(ctx context.Context, input source.NewDict) (bool, error) {
	_, err := r.Create(ctx, new(beans.Dict), input)
	if err != nil {
		return false, err
	}
	go func() {
		time.Sleep(time.Second * 3)
		event.DictChange()
	}()
	return true, nil
}

func (r *mutationResolver) DictUpdate(ctx context.Context, id string, input source.UpdDict) (bool, error) {
	ok, err := r.Update(ctx, new(beans.Dict), id, input)
	if err != nil {
		return false, err
	}
	if ok {
		go func() {
			time.Sleep(time.Second * 3)
			event.DictChange()
		}()
	}
	return ok, err
}

func (r *mutationResolver) DictRemoves(ctx context.Context, ids []string) (bool, error) {
	ok, err := r.Removes(ctx, new(beans.Dict), ids)
	if err != nil {
		return false, err
	}
	if ok {
		go func() {
			time.Sleep(time.Second * 3)
			event.DictChange()
		}()
	}
	return ok, err
}

func (r *mutationResolver) DictItemCreate(ctx context.Context, input source.NewDictItem) (bool, error) {
	_, err := r.Create(ctx, new(beans.DictItem), input)
	if err != nil {
		return false, err
	}
	go func() {
		time.Sleep(time.Second * 3)
		event.DictChange()
	}()
	return true, nil
}

func (r *mutationResolver) DictItemUpdate(ctx context.Context, id string, input source.UpdDictItem) (bool, error) {
	ok, err := r.Update(ctx, new(beans.DictItem), id, input)
	if err != nil {
		return false, err
	}
	if ok {
		go func() {
			time.Sleep(time.Second * 3)
			event.DictChange()
		}()
	}
	return ok, err
}

func (r *mutationResolver) DictItemRemoves(ctx context.Context, ids []string) (bool, error) {
	ok, err := r.Removes(ctx, new(beans.DictItem), ids)
	if err != nil {
		return false, err
	}
	if ok {
		go func() {
			time.Sleep(time.Second * 3)
			event.DictChange()
		}()
	}
	return ok, err
}

func (r *mutationResolver) DictItemSort(ctx context.Context, code string, items []string) (bool, error) {
	for i, o := range items {
		err := r.Sess(ctx).SF(`update dict_item set weight = :weight where id = :id and code = :code`, map[string]interface{}{"weight": i, "id": o, "code": code}).Exec()
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

func (r *queryResolver) Dicts(ctx context.Context, query *source.QDict) ([]beans.Dict, error) {
	if query == nil {
		query = &source.QDict{}
	}
	dicts := make([]beans.Dict, 0)
	err := r.DBS(ctx).SF(`
		select 
			u.* 
		from 
			dict u 
		where 1 = 1 
		{{ if .type }} and u.type = :type {{ end }} 
		{{ if .types }} and u.type in :types {{ end }} 
		{{ if .status }} and u.status = :status {{ end }} 
		{{ if .dicts }} and u.code in :dicts {{ end }} 
		`,
		map[string]interface{}{
			"type":   query.Type,
			"types":  query.Types,
			"status": query.Status,
			"dicts":  query.Dicts,
		}).Order("u.code").Find(&dicts)
	return dicts, err
}

func (r *queryResolver) Dict(ctx context.Context, id string) (*beans.Dict, error) {
	return r.DictLoader(ctx, id)
}

// Dict returns source.DictResolver implementation.
func (r *Resolver) Dict() source.DictResolver { return &dictResolver{r} }

type dictResolver struct{ *Resolver }
