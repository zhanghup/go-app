package api

import (
	"context"
	"github.com/go-xorm/xorm"
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/toolxorm"
)

func (this *Resolver) Create(ctx context.Context, tab interface{}, obj interface{}) (string, error) {
	id := tools.Str.Uid()
	sess := this.DBS.NewSession(ctx)
	err := sess.TS(func(sess *toolxorm.Session) error {
		_, err := sess.Sess.Table(tab).Insert(&beans.Bean{
			Id:     &id,
			Status: tools.Ptr.Int(1),
		})
		if err != nil {
			return err
		}
		_, err = sess.Sess.Table(tab).Where("id = ?", id).Update(obj)
		return err
	})

	return id, err
}

func (this *Resolver) Update(ctx context.Context, tab interface{}, id string, obj interface{}, commit ...bool) (bool, error) {
	ctx, err := this.DB.Ts(ctx, func(s *xorm.Session) error {
		_, err := s.Table(tab).Where("id = ?", id).Update(beans.Bean{})
		if err != nil {
			return err
		}
		_, err = s.Table(tab).Where("id = ?", id).AllCols().Update(obj)
		return err
	}, commit...)

	return err == nil, err
}

func (this *Resolver) Removes(ctx context.Context, table interface{}, ids []string, commit ...bool) (bool, error) {
	ctx, err := this.DB.Ts(ctx, func(s *xorm.Session) error {
		_, err := s.Table(table).In("id", ids).Delete(table)
		return err
	}, commit...)

	return err == nil, err
}
