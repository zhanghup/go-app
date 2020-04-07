package api

import (
	"context"
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/toolxorm"
	"reflect"
)

func (this *Resolver) Create(ctx context.Context, tab interface{}, obj interface{}) (string, error) {
	id := ""
	sess := this.DBS.NewSession(ctx)
	err := sess.TS(func(sess *toolxorm.Session) error {
		tools.Rft.DeepSet(tab, func(t reflect.Type, v reflect.Value, tf reflect.StructField) bool {
			switch tf.Name {
			case "Id":
				if t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.String {
					if v.Pointer() == 0 {
						id = tools.Str.Uid()
						v.Set(reflect.ValueOf(&id))
					}
				} else if t.Kind() == reflect.String {
					if v.String() == "" {
						v.Set(reflect.ValueOf(tools.Str.Uid()))
					}
				}
			case "Status":
				if t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Int {
					if v.Pointer() == 0 {
						v.Set(reflect.ValueOf(tools.Ptr.Int(1)))
					}
				}
			}
			return true
		})

		_, err := sess.Sess.Insert(tab)
		if err != nil {
			return err
		}
		_, err = sess.Sess.Table(tab).Where("id = ?", id).Update(obj)
		return err
	})

	return id, err
}

func (this *Resolver) Update(ctx context.Context, tab interface{}, id string, obj interface{}) (bool, error) {
	err := this.DBS.NewSession(ctx).TS(func(sess *toolxorm.Session) error {
		_, err := sess.Sess.Table(tab).Where("id = ?", id).Update(beans.Bean{})
		if err != nil {
			return err
		}
		_, err = sess.Sess.Table(tab).Where("id = ?", id).AllCols().Update(obj)
		return err
	})

	return err == nil, err
}

func (this *Resolver) Removes(ctx context.Context, table interface{}, ids []string) (bool, error) {
	err := this.DBS.NewSession(ctx).TS(func(sess *toolxorm.Session) error {
		_, err := sess.Sess.Table(table).In("id", ids).Delete(table)
		return err
	})
	return err == nil, err
}
