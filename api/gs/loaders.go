package gs

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/zhanghup/go-tools"
	"net/http"
	"reflect"
	"strings"
	"sync"
	"time"
)

const DATA_LOADEN = "go-app-dataloaden"

type Loader interface {
	Object(obj interface{}) *CommonLoader
	Slice(obj interface{}, key string, param ...map[string]interface{}) *CommonSliceLoader
}

type dataLoaden struct {
	sync  sync.Mutex
	store tools.IMap
	db    *xorm.Session
}

func DataLoaden(ctx context.Context) Loader {
	return ctx.Value(DATA_LOADEN).(*dataLoaden)
}
func DataLoadenMiddleware(db *xorm.Engine, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), DATA_LOADEN, &dataLoaden{
			sync:  sync.Mutex{},
			store: tools.NewCache(),
			db:    db.Context(r.Context()),
		})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (this *dataLoaden) Object(obj interface{}) *CommonLoader {
	this.sync.Lock()
	defer this.sync.Unlock()

	ty := reflect.TypeOf(obj)
	if ty.Kind() == reflect.Ptr {
		ty = ty.Elem()
	}

	objType := "Common-Object-" + ty.PkgPath() + "." + ty.Name()
	o := this.store.Get(objType)
	if o != nil {
		return o.(*CommonLoader)
	}
	var lo *CommonLoader
	defer this.store.Set(objType, lo)

	lo = &CommonLoader{
		maxBatch: 128,
		wait:     1 * time.Millisecond,
		fetch: func(keys []string) ([]interface{}, []error) {

			ll := reflect.New(reflect.SliceOf(reflect.TypeOf(obj).Elem()))

			err := this.db.Table(obj).In("id", keys).Find(ll.Interface())
			if err != nil {
				return nil, []error{err}
			}

			ll = ll.Elem()
			rmap := map[string]interface{}{}
			for i := 0; i < ll.Len(); i++ {
				vl := ll.Index(i)
				id, err := commonGetField(vl, "Id")
				if err != nil {
					return nil, []error{err}
				}
				rmap[id.Elem().String()] = vl.Interface()
			}

			rs := make([]interface{}, len(keys))
			for i, str := range keys {
				rs[i] = rmap[str]
			}

			return rs, nil
		},
	}
	return lo
}

func (this *dataLoaden) Slice(obj interface{}, key string, params ...map[string]interface{}) *CommonSliceLoader {
	this.sync.Lock()
	defer this.sync.Unlock()

	param := map[string]interface{}{}
	if len(params) > 0 {
		param = params[0]
	}

	ty := reflect.TypeOf(obj)
	if ty.Kind() == reflect.Ptr {
		ty = ty.Elem()
	}

	objType := fmt.Sprintf("Common-Slice-%s.%s,key=%v&param=%v", ty.PkgPath(), ty.Name(), key, tools.Str().JSONString(param))
	o := this.store.Get(objType)
	if o != nil {
		return o.(*CommonSliceLoader)
	}
	this.store.Set(objType, &CommonSliceLoader{
		maxBatch: 128,
		wait:     1 * time.Millisecond,
		fetch: func(keys []string) (i interface{}, errors []error) {
			tySlice := reflect.SliceOf(reflect.TypeOf(obj))
			tyMap := reflect.MapOf(reflect.TypeOf(""), tySlice)

			ll := reflect.New(tySlice)

			s := this.db.Table(obj).In(key, keys)
			if param != nil {
				for k, v := range param {
					if strings.Index(k, "[]") == 0 {
						s.In(strings.ReplaceAll(k, "[]", ""), v)
					} else {
						s.And(fmt.Sprintf("%s = ?", k), v)
					}
				}
			}
			err := s.Find(ll.Interface())
			if err != nil {
				return nil, []error{err}
			}

			ll = ll.Elem()

			rmap := reflect.MakeMap(tyMap)

			for i := 0; i < ll.Len(); i++ {
				vl := ll.Index(i)
				id, err := commonGetField(vl.Elem(), UpTitle(key))
				if err != nil {
					return nil, []error{err}
				}
				if r := rmap.MapIndex(id.Elem()); r.IsValid() {
					r = reflect.Append(r, vl)
					rmap.SetMapIndex(id.Elem(), r)
				} else {
					sl := reflect.New(tySlice).Elem()
					sl = reflect.Append(sl, vl)
					rmap.SetMapIndex(id.Elem(), sl)
				}
			}
			return rmap.Interface(), nil
		},
	})
	return this.store.Get(objType).(*CommonSliceLoader)
}
func UpTitle(str string) string {
	ss := ""
	for _, s := range strings.Split(str, "_") {
		ss += strings.Title(s)
	}
	return ss
}

func commonGetField(vl reflect.Value, fieldname string) (*reflect.Value, error) {

	// 这里有一个问题，可能会优先取到继承对象中的相同字段
	for i := 0; i < vl.Type().NumField(); i++ {
		v := vl.Field(i)
		t := vl.Type().Field(i)
		if t.Name == fieldname {
			return &v, nil
		}
		if t.Type.Kind() == reflect.Struct {
			commonGetField(v, fieldname)
		}
	}
	return nil, errors.New("没有找到对应的属性值")

}
