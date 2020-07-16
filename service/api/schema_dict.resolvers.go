package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/service/api/lib"
	"github.com/zhanghup/go-app/service/event"
)

func (r *dictResolver) Values(ctx context.Context, obj *beans.Dict) ([]beans.DictItem, error) {
	if obj.Id == nil {
		return nil, nil
	}
	c := make([]beans.DictItem, 0)
	err := r.Loader(ctx).Slice(c, "select * from dict_item where code in :keys", nil, "Code", "").Load(*obj.Id, &c)
	return c, err
}

func (r *mutationResolver) DictCreate(ctx context.Context, input lib.NewDict) (bool, error) {
	_, err := r.Create(ctx, new(beans.Dict), input)
	if err != nil {
		return false, err
	}
	go event.DictChange()
	return true, nil
}

func (r *mutationResolver) DictUpdate(ctx context.Context, id string, input lib.UpdDict) (bool, error) {
	ok, err := r.Update(ctx, new(beans.Dict), id, input)
	if err != nil {
		return false, err
	}
	if ok {
		go event.DictChange()
	}
	return ok, err
}

func (r *mutationResolver) DictRemoves(ctx context.Context, ids []string) (bool, error) {
	ok, err := r.Removes(ctx, new(beans.Dict), ids)
	if err != nil {
		return false, err
	}
	if ok {
		go event.DictChange()
	}
	return ok, err
}

func (r *mutationResolver) DictItemCreate(ctx context.Context, input lib.NewDictItem) (bool, error) {
	_, err := r.Create(ctx, new(beans.DictItem), input)
	if err != nil {
		return false, err
	}
	go event.DictChange()
	return true, nil
}

func (r *mutationResolver) DictItemUpdate(ctx context.Context, id string, input lib.UpdDictItem) (bool, error) {
	ok, err := r.Update(ctx, new(beans.DictItem), id, input)
	if err != nil {
		return false, err
	}
	if ok {
		go event.DictChange()
	}
	return ok, err
}

func (r *mutationResolver) DictItemRemoves(ctx context.Context, ids []string) (bool, error) {
	ok, err := r.Removes(ctx, new(beans.DictItem), ids)
	if err != nil {
		return false, err
	}
	if ok {
		go event.DictChange()
	}
	return ok, err
}

func (r *queryResolver) Dicts(ctx context.Context) ([]beans.Dict, error) {
	dicts := make([]beans.Dict, 0)
	err := r.DBS.SF(` select * from dict u where 1 = 1 order by u.code`).Find(&dicts)
	return dicts, err
}

func (r *queryResolver) Dict(ctx context.Context, id string) (*beans.Dict, error) {
	return r.DictLoader(ctx, id)
}

// Dict returns lib.DictResolver implementation.
func (r *Resolver) Dict() lib.DictResolver { return &dictResolver{r} }

type dictResolver struct{ *Resolver }
