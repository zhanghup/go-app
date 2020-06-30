package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/service/api/lib"
)

func (r *dictResolver) Values(ctx context.Context, obj *beans.Dict) ([]beans.DictItem, error) {
	if obj.Id == nil {
		return nil, nil
	}
	c := make([]beans.DictItem, 0)
	err := r.Loader(ctx).Slice(c, "select * from dict_item where id in :keys", nil, "Code", "").Load(*obj.Id, &c)
	return c, err
}

func (r *mutationResolver) DictCreate(ctx context.Context, input lib.NewDict) (*beans.Dict, error) {
	id, err := r.Create(ctx, new(beans.Dict), input)
	if err != nil {
		return nil, err
	}
	return r.DictLoader(ctx, id)
}

func (r *mutationResolver) DictUpdate(ctx context.Context, id string, input lib.UpdDict) (bool, error) {
	return r.Update(ctx, new(beans.Dict), id, input)
}

func (r *mutationResolver) DictRemoves(ctx context.Context, ids []string) (bool, error) {
	return r.Removes(ctx, new(beans.Dict), ids)
}

func (r *mutationResolver) DictItemCreate(ctx context.Context, input lib.NewDictItem) (*beans.DictItem, error) {
	id, err := r.Create(ctx, new(beans.DictItem), input)
	if err != nil {
		return nil, err
	}
	return r.DictItemLoader(ctx, id)
}

func (r *mutationResolver) DictItemUpdate(ctx context.Context, id string, input lib.UpdDictItem) (bool, error) {
	return r.Update(ctx, new(beans.DictItem), id, input)
}

func (r *mutationResolver) DictItemRemoves(ctx context.Context, ids []string) (bool, error) {
	return r.Removes(ctx, new(beans.DictItem), ids)
}

func (r *queryResolver) Dicts(ctx context.Context, query lib.QDict) (*lib.Dicts, error) {
	dicts := make([]beans.Dict, 0)
	total, err := r.DBS.SF(`
		select * from dict u
		where 1 = 1
		order by u.code
	`).Page2(query.Index, query.Size, query.Count, &dicts)
	return &lib.Dicts{Data: dicts, Total: &total}, err
}

func (r *queryResolver) Dict(ctx context.Context, id string) (*beans.Dict, error) {
	return r.DictLoader(ctx, id)
}

// Dict returns lib.DictResolver implementation.
func (r *Resolver) Dict() lib.DictResolver { return &dictResolver{r} }

type dictResolver struct{ *Resolver }
