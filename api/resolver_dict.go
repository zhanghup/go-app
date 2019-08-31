package api

import (
	"context"
	"github.com/zhanghup/go-app"
	"github.com/zhanghup/go-app/api/gs"
)

type DictResolver struct {
	*Resolver
}

func (this *Resolver) Dict() gs.DictResolver {
	return DictResolver{this}
}

func (this DictResolver) Values(ctx context.Context, obj *app.Dict) ([]*app.DictItem, error) {
	panic("implement me")
}

func (this queryResolver) Dicts(ctx context.Context, query gs.QDict) (*gs.Dicts, error) {
	panic("implement me")
}

func (this queryResolver) Dict(ctx context.Context, id string) (*app.Dict, error) {
	panic("implement me")
}

func (this mutationResolver) DictCreate(ctx context.Context, input gs.NewDict) (*app.Dict, error) {
	panic("implement me")
}

func (this mutationResolver) DictUpdate(ctx context.Context, id string, input gs.UpdDict) (bool, error) {
	panic("implement me")
}

func (this mutationResolver) DictRemoves(ctx context.Context, ids []string) (bool, error) {
	panic("implement me")
}

func (this mutationResolver) DictItemCreate(ctx context.Context, input gs.NewDictItem) (*app.Dict, error) {
	panic("implement me")
}

func (this mutationResolver) DictItemUpdate(ctx context.Context, id string, input gs.UpdDictItem) (bool, error) {
	panic("implement me")
}

func (this mutationResolver) DictItemRemoves(ctx context.Context, ids []string) (bool, error) {
	panic("implement me")
}

