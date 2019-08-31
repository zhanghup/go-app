package api

import (
	"context"
	"github.com/zhanghup/go-app"
	"github.com/zhanghup/go-tools"
)

func (this *Resolver) Create(ctx context.Context, tab interface{}, obj interface{}) (string, error) {
	id := tools.ObjectString()
	_, err := this.DB(ctx).Table(tab).Insert(&app.Bean{
		Id:     id,
		Status: tools.Ptr().Int(1),
	})
	if err != nil {
		return "", err
	}
	_, err = this.DB(ctx).Table(tab).Where("id = ?", id).Update(obj)
	return *id, err
}

func (this *Resolver) Update(ctx context.Context, tab interface{}, id string, obj interface{}) (bool, error) {
	_, err := this.DB(ctx).Table(tab).Where("id = ?", id).Update(app.Bean{})
	if err != nil {
		return false, err
	}
	_, err = this.DB(ctx).Table(tab).Where("id = ?", id).AllCols().Update(obj)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (this *Resolver) Remove(ctx context.Context, tab interface{}, id string) {

}

func (this *Resolver) UserLoader(ctx context.Context, id string) (*app.User, error) {
	obj, err := this.Loader(ctx).Common(new(app.User)).Load(id)
	if err != nil {
		return nil, err
	}
	user, ok := obj.(app.User)
	if !ok {
		return nil, nil
	}
	return &user, nil
}
