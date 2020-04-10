package api

import (
	"context"
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/service/api/lib"
)

type DictResolver struct {
	*Resolver
}

func (this *Resolver) Dict() lib.DictResolver {
	return DictResolver{this}
}

func (this *Resolver) DictLoader(ctx context.Context, id string) (*beans.Dict, error) {
	result := new(beans.Dict)
	_, err := this.Loader(ctx).Object(new(beans.Dict)).Load(id, result)
	return result, err
}

func (this *Resolver) DictItemLoader(ctx context.Context, id string) (*beans.DictItem, error) {
	result := new(beans.DictItem)
	_, err := this.Loader(ctx).Object(new(beans.DictItem)).Load(id, result)
	return result, err
}

func (this DictResolver) Values(ctx context.Context, obj *beans.Dict) ([]beans.DictItem, error) {
	if obj.Id == nil {
		return nil, nil
	}
	c := make([]beans.DictItem, 0)
	err := this.Loader(ctx).Slice(new(beans.DictItem), "code").Load(*obj.Id, &c)
	return c, err
}

func (this queryResolver) Dicts(ctx context.Context, query lib.QDict) (*lib.Dicts, error) {
	dicts := make([]beans.Dict, 0)
	total, err := this.DBS.SF(`
		select * from dict u
		where 1 = 1
		order by u.code
	`).Page2(query.Index, query.Size, query.Count, &dicts)
	return &lib.Dicts{Data: dicts, Total: &total}, err
}

func (this queryResolver) Dict(ctx context.Context, id string) (*beans.Dict, error) {
	return this.DictLoader(ctx, id)
}

func (this mutationResolver) DictCreate(ctx context.Context, input lib.NewDict) (*beans.Dict, error) {
	id, err := this.Create(ctx, new(beans.Dict), input)
	if err != nil {
		return nil, err
	}
	return this.DictLoader(ctx, id)
}

func (this mutationResolver) DictUpdate(ctx context.Context, id string, input lib.UpdDict) (bool, error) {
	return this.Update(ctx, new(beans.Dict), id, input)
}

func (this mutationResolver) DictRemoves(ctx context.Context, ids []string) (bool, error) {
	return this.Removes(ctx, new(beans.Dict), ids)
}

func (this mutationResolver) DictItemCreate(ctx context.Context, input lib.NewDictItem) (*beans.DictItem, error) {
	id, err := this.Create(ctx, new(beans.DictItem), input)
	if err != nil {
		return nil, err
	}
	return this.DictItemLoader(ctx, id)
}

func (this mutationResolver) DictItemUpdate(ctx context.Context, id string, input lib.UpdDictItem) (bool, error) {
	return this.Update(ctx, new(beans.DictItem), id, input)
}

func (this mutationResolver) DictItemRemoves(ctx context.Context, ids []string) (bool, error) {
	return this.Removes(ctx, new(beans.DictItem), ids)
}
