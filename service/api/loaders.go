package api

import (
	"context"
	"github.com/zhanghup/go-app/beans"
)

func (this *Resolver) DictLoader(ctx context.Context, id string) (*beans.Dict, error) {
	result := new(beans.Dict)
	err := this.Loader(ctx).Object(result, "select * from dict where id in :keys", nil, "Id", "").Load(id, result)
	return result, err
}

func (this *Resolver) DictItemLoader(ctx context.Context, id string) (*beans.DictItem, error) {
	result := new(beans.DictItem)
	err := this.Loader(ctx).Object(result, "select * from dict_item where id in :keys", nil, "Id", "").Load(id, result)
	return result, err
}

func (this *Resolver) RoleLoader(ctx context.Context, id string) (*beans.Role, error) {
	result := new(beans.Role)
	err := this.Loader(ctx).Object(result, "select * from role where id in :keys", nil, "Id", "").Load(id, result)
	return result, err
}
func (this *Resolver) UserLoader(ctx context.Context, id string) (*beans.User, error) {
	result := new(beans.User)
	err := this.Loader(ctx).Object(result, "select * from user where id in :keys", nil, "Id", "").Load(id, result)
	return result, err
}