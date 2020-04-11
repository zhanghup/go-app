package loaders

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/toolxorm"
	"reflect"
	"sync"
	"xorm.io/xorm"
)

type Loader interface {
	Object(table interface{}, sql string, param map[string]interface{}, keyField string) *ObjectLoader
}

type dataLoaden struct {
	sync  sync.Locker
	store tools.ICache
	db    *xorm.Engine
}

func NewDataLoaden(db *xorm.Engine) Loader {
	return &dataLoaden{sync: &sync.Mutex{}, db: db, store: tools.CacheCreate()}
}

func (this *dataLoaden) Object(table interface{}, sql string, param map[string]interface{}, keyField string) *ObjectLoader {
	query := map[string]interface{}{}
	if param != nil {
		query = param
	}
	path := reflect.TypeOf(tools.Rft.RealValue(table)).PkgPath()
	name := reflect.TypeOf(tools.Rft.RealValue(table)).Name()
	key := fmt.Sprintf("table: %s/%s, sql: %s, param: %s, field: %s", path, name, sql, tools.Str.JSONString(query), keyField)
	key = tools.Crypto.MD5([]byte(key))
	this.sync.Lock()
	defer this.sync.Unlock()
	objLoader := new(ObjectLoader)
	if this.store.Get(key, objLoader) {
		return objLoader
	}
	objLoader = &ObjectLoader{
		sync:     &sync.Mutex{},
		db:       toolxorm.NewEngine(this.db),
		keyField: keyField,
		sql:      sql,
		param:    query,
		table:    table,
	}
	this.store.Set(key, objLoader)
	return objLoader
}

//
//import (
//	"context"
//	"errors"
//	"fmt"
//	"github.com/zhanghup/go-tools"
//	"net/http"
//	"reflect"
//	"strings"
//	"sync"
//	"time"
//	"xorm.io/xorm"
//)
//
//const DATA_LOADEN = "go-app-dataloaden"
//
//type Loader interface {
//	Object(obj interface{}) *CommonLoader
//	Slice(obj interface{}, key string, param ...map[string]interface{}) *CommonSliceLoader
//}
//
//type dataLoaden struct {
//	sync  sync.Locker
//	store tools.ICache
//	db    *xorm.Engine
//}
//
//func DataLoaden(ctx context.Context) Loader {
//	return ctx.Value(DATA_LOADEN).(*dataLoaden)
//}
//func DataLoadenMiddleware(db *xorm.Engine, next http.HandlerFunc) http.HandlerFunc {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		ctx := context.WithValue(r.Context(), DATA_LOADEN, &dataLoaden{
//			sync:  &sync.Mutex{},
//			store: tools.CacheCreate(),
//			db:    db,
//		})
//		next.ServeHTTP(w, r.WithContext(ctx))
//	})
//}
//
//func (this *dataLoaden) Object(obj interface{}) *CommonLoader {
//	this.sync.Lock()
//	defer this.sync.Unlock()
//
//	ty := reflect.TypeOf(obj)
//	if ty.Kind() == reflect.Ptr {
//		ty = ty.Elem()
//	}
//
//	objType := "Common-Object-" + ty.PkgPath() + "." + ty.Name()
//	var lo *CommonLoader
//	ok := this.store.Get(objType, lo)
//	if ok {
//		return lo
//	}
//
//	defer this.store.Set(objType, lo)
//
//	lo = &CommonLoader{
//		maxBatch: 128,
//		wait:     1 * time.Millisecond,
//		fetch: func(keys []string) ([]interface{}, []error) {
//
//			ll := reflect.New(reflect.SliceOf(reflect.TypeOf(obj).Elem()))
//
//			err := this.db.Table(obj).In("id", keys).Find(ll.Interface())
//			if err != nil {
//				return nil, []error{err}
//			}
//
//			ll = ll.Elem()
//			rmap := map[string]interface{}{}
//			for i := 0; i < ll.Len(); i++ {
//				vl := ll.Index(i)
//				id, err := commonGetField(vl, "Id")
//				if err != nil {
//					return nil, []error{err}
//				}
//				rmap[id.Elem().String()] = vl.Interface()
//			}
//
//			rs := make([]interface{}, len(keys))
//			for i, str := range keys {
//				rs[i] = rmap[str]
//			}
//
//			return rs, nil
//		},
//	}
//	return lo
//}
//
//func (this *dataLoaden) Slice(obj interface{}, key string, params ...map[string]interface{}) *CommonSliceLoader {
//	this.sync.Lock()
//	defer this.sync.Unlock()
//
//	param := map[string]interface{}{}
//	if len(params) > 0 {
//		param = params[0]
//	}
//
//	ty := reflect.TypeOf(obj)
//	if ty.Kind() == reflect.Ptr {
//		ty = ty.Elem()
//	}
//
//	objType := fmt.Sprintf("Common-Slice-%s.%s,key=%v&param=%v", ty.PkgPath(), ty.Name(), key, tools.Str.JSONString(param))
//	var lo *CommonSliceLoader
//	ok := this.store.Get(objType, lo)
//	if ok {
//		return lo
//	}
//	defer this.store.Set(objType, lo)
//	lo = &CommonSliceLoader{
//		maxBatch: 128,
//		wait:     1 * time.Millisecond,
//		fetch: func(keys []string) (i interface{}, errors []error) {
//			tySlice := reflect.SliceOf(reflect.TypeOf(obj).Elem())
//			tyMap := reflect.MapOf(reflect.TypeOf(""), tySlice)
//
//			ll := reflect.New(tySlice)
//
//			s := this.db.Table(obj).In(key, keys)
//			if param != nil {
//				for k, v := range param {
//					if strings.Index(k, "[]") == 0 {
//						s.In(strings.ReplaceAll(k, "[]", ""), v)
//					} else {
//						s.And(fmt.Sprintf("%s = ?", k), v)
//					}
//				}
//			}
//			err := s.Find(ll.Interface())
//			if err != nil {
//				return nil, []error{err}
//			}
//
//			ll = ll.Elem()
//
//			rmap := reflect.MakeMap(tyMap)
//
//			for i := 0; i < ll.Len(); i++ {
//				vl := ll.Index(i)
//				id, err := commonGetField(vl, UpTitle(key))
//				if err != nil {
//					return nil, []error{err}
//				}
//				if r := rmap.MapIndex(id.Elem()); r.IsValid() {
//					r = reflect.Append(r, vl)
//					rmap.SetMapIndex(id.Elem(), r)
//				} else {
//					sl := reflect.New(tySlice).Elem()
//					sl = reflect.Append(sl, vl)
//					rmap.SetMapIndex(id.Elem(), sl)
//				}
//			}
//			return rmap.Interface(), nil
//		},
//	}
//	return lo
//}
//func UpTitle(str string) string {
//	ss := ""
//	for _, s := range strings.Split(str, "_") {
//		ss += strings.Title(s)
//	}
//	return ss
//}
