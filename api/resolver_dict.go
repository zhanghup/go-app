package api

import (
	"context"
	"github.com/zhanghup/go-app"
	"github.com/zhanghup/go-app/api/gs"
)

type DictResolver struct {
	*Resolver
}

func (this *Resolver) DictLoader(ctx context.Context, id string) (*app.Dict, error) {
	obj, err := this.Loader(ctx).Object(new(app.Dict)).Load(id)
	if err != nil {
		return nil, err
	}
	dict, ok := obj.(app.Dict)
	if !ok {
		return nil, nil
	}
	return &dict, nil
}

func (this *Resolver) DictItemLoader(ctx context.Context, id string) (*app.DictItem, error) {
	obj, err := this.Loader(ctx).Object(new(app.DictItem)).Load(id)
	if err != nil {
		return nil, err
	}
	dict, ok := obj.(app.DictItem)
	if !ok {
		return nil, nil
	}
	return &dict, nil
}

func (this *Resolver) Dict() gs.DictResolver {
	return DictResolver{this}
}

func (this DictResolver) Values(ctx context.Context, obj *app.Dict) ([]*app.DictItem, error) {
	if obj.Code == nil {
		return nil, nil
	}
	var c []*app.DictItem
	err := this.Loader(ctx).Slice(new(app.DictItem), "code").Load(*obj.Code, &c)
	return c, err
}

func (this queryResolver) Dicts(ctx context.Context, query gs.QDict) (*gs.Dicts, error) {
	dicts := make([]*app.Dict, 0)
	_, total, err := this.DB.SF(`
		select * from {{ table "dict" }} u
		where 1 = 1
	`).Page2(query.Index, query.Size, query.Count, &dicts)
	return &gs.Dicts{Data: dicts, Total: &total}, err
}

func (this queryResolver) Dict(ctx context.Context, id string) (*app.Dict, error) {
	return this.DictLoader(ctx, id)
}

func (this mutationResolver) DictCreate(ctx context.Context, input gs.NewDict) (*app.Dict, error) {
	id, err := this.Create(ctx, new(app.Dict), input)
	if err != nil {
		return nil, err
	}
	return this.DictLoader(ctx, id)
}

func (this mutationResolver) DictUpdate(ctx context.Context, id string, input gs.UpdDict) (bool, error) {
	return this.Update(ctx, new(app.Dict), id, input)
}

func (this mutationResolver) DictRemoves(ctx context.Context, ids []string) (bool, error) {
	return this.Removes(ctx, new(app.Dict), ids)
}

func (this mutationResolver) DictItemCreate(ctx context.Context, input gs.NewDictItem) (*app.DictItem, error) {
	id, err := this.Create(ctx, new(app.DictItem), input)
	if err != nil {
		return nil, err
	}
	return this.DictItemLoader(ctx, id)
}

func (this mutationResolver) DictItemUpdate(ctx context.Context, id string, input gs.UpdDictItem) (bool, error) {
	return this.Update(ctx, new(app.DictItem), id, input)
}

func (this mutationResolver) DictItemRemoves(ctx context.Context, ids []string) (bool, error) {
	return this.Removes(ctx, new(app.DictItem), ids)
}
