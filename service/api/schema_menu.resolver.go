package api

import (
	"context"
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/service/api/source"
)

func (r *mutationResolver) MenuUpdate(ctx context.Context, id string, input source.UpdMenu) (bool, error) {
	panic("implement me")
}

func (r *mutationResolver) MenuReload(ctx context.Context, menus []source.MenuLocal) (bool, error) {
	panic("implement me")
}


func (r *queryResolver) Menus(ctx context.Context, query source.QMenu) (*source.Menus, error) {
	panic("implement me")
}

func (r *queryResolver) Menu(ctx context.Context, id string) (*beans.Menu, error) {
	panic("implement me")
}
