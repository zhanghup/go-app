package api

import (
	"context"
	"github.com/go-xorm/xorm"
	"github.com/zhanghup/go-app"
	"github.com/zhanghup/go-tools"
)

func (this *Resolver) Create(ctx context.Context, tab interface{}, obj interface{}, commit ...bool) (string, error) {
	flag := true
	if len(commit) > 0 && !commit[0] {
		flag = false
	}
	id := tools.ObjectString()
	ctx, err := this.DB.Ts(ctx, func(s *xorm.Session) error {

		_, err := s.Table(tab).Insert(&app.Bean{
			Id:     id,
			Status: tools.Ptr().Int(1),
		})
		if err != nil {
			return err
		}
		_, err = s.Table(tab).Where("id = ?", id).Update(obj)
		return err
	}, flag)

	return *id, err
}

func (this *Resolver) Update(ctx context.Context, tab interface{}, id string, obj interface{}, commit ...bool) (bool, error) {
	flag := true
	if len(commit) > 0 && !commit[0] {
		flag = false
	}
	ctx, err := this.DB.Ts(ctx, func(s *xorm.Session) error {
		_, err := s.Table(tab).Where("id = ?", id).Update(app.Bean{})
		if err != nil {
			return err
		}
		_, err = s.Table(tab).Where("id = ?", id).AllCols().Update(obj)
		return err
	}, flag)

	return err == nil, err
}

func (this *Resolver) Removes(ctx context.Context, table interface{}, ids []string, commit ...bool) (bool, error) {
	flag := true
	if len(commit) > 0 && !commit[0] {
		flag = false
	}
	ctx, err := this.DB.Ts(ctx, func(s *xorm.Session) error {
		_, err := s.Table(table).In("id", ids).Delete(table)
		return err
	}, flag)

	return err == nil, err
}
