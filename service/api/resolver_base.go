package api

import (
	"context"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
	"reflect"
	"time"
)

func (this *ResolverTools) Create(ctx context.Context, tab interface{}, obj interface{}) (string, error) {
	id := ""
	err := this.Sess(ctx).TS(func(sess txorm.ISession) error {
		tools.Rft.DeepSet(tab, func(t reflect.Type, v reflect.Value, tf reflect.StructField) bool {
			switch tf.Name {
			case "Id":
				if t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.String {
					if v.Pointer() == 0 {
						id = tools.UUID()
						v.Set(reflect.ValueOf(&id))
					}
				} else if t.Kind() == reflect.String {
					if v.String() == "" {
						v.Set(reflect.ValueOf(tools.UUID()))
					}
				}
			case "Status":
				if t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.String {
					if v.Pointer() == 0 {
						v.Set(reflect.ValueOf(tools.PtrOfString("1")))
					}
				}
			case "Weight":
				if t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Int {
					if v.Pointer() == 0 {
						v.Set(reflect.ValueOf(tools.PtrOfInt(int(time.Now().Unix() - 1610541047))))
					}
				}
			}
			return true
		})

		_, err := sess.S().Insert(tab)
		if err != nil {
			return err
		}
		if obj != nil {
			_, err = sess.S().Table(tab).Where("id = ?", id).Update(obj)
		}
		return err
	})

	return id, err
}

func (this *ResolverTools) Update(ctx context.Context, tab interface{}, id string, obj interface{}) (bool, error) {
	err := this.Sess(ctx).TS(func(sess txorm.ISession) error {
		_, err := sess.S().Table(tab).Where("id = ?", id).Update(tab)
		if err != nil {
			return err
		}
		_, err = sess.S().Table(tab).Where("id = ?", id).AllCols().Update(obj)
		return err
	})

	return err == nil, err
}

func (this *ResolverTools) Removes(ctx context.Context, table interface{}, ids []string) (bool, error) {
	err := this.Sess(ctx).TS(func(sess txorm.ISession) error {
		_, err := sess.S().Table(table).In("id", ids).Delete(table)
		return err
	})
	return err == nil, err
}
