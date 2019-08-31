package gs

import (
	"context"
	"errors"
	"github.com/go-xorm/xorm"
	"github.com/zhanghup/go-tools"
	"net/http"
	"reflect"
	"sync"
	"time"
)

const DATA_LOADEN = "go-app-dataloaden"

type Loader interface {
	Common(obj interface{}) *CommonLoader
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

func (this *dataLoaden) Common(obj interface{}) *CommonLoader {
	this.sync.Lock()
	defer this.sync.Unlock()

	ty := reflect.TypeOf(obj)
	if ty.Kind() == reflect.Ptr {
		ty = ty.Elem()
	}

	objType := "CommonFn-" + ty.PkgPath() + "." + ty.Name()
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

func commonGetField(vl reflect.Value, fieldname string) (*reflect.Value, error) {

	// 这里有一个问题，可能会优先取到继承对象中的相同字段
	for i := 0; i < vl.Type().NumField(); i++ {
		v := vl.Field(i)
		t := vl.Type().Field(i)
		if t.Name == fieldname {
			return &v, nil
		}
		if t.Type.Kind() == reflect.Struct {
			return commonGetField(v, fieldname)
		}
	}
	return nil, errors.New("没有找到对应的属性值")

}

// CommonLoader batches and caches requests
type CommonLoader struct {
	// this method provides the data for the loader
	fetch func(keys []string) ([]interface{}, []error)

	// how long to done before sending a batch
	wait time.Duration

	// this will limit the maximum number of keys to send in one batch, 0 = no limit
	maxBatch int

	// INTERNAL

	// lazily created cache
	cache map[string]interface{}

	// the current batch. keys will continue to be collected until timeout is hit,
	// then everything will be sent to the fetch method and out to the listeners
	batch *commonBatch

	// mutex to prevent races
	mu sync.Mutex
}

type commonBatch struct {
	keys    []string
	data    []interface{}
	error   []error
	closing bool
	done    chan struct{}

	table interface{} // 保存数据库model对象
}

// Load a dict by key, batching and caching will be applied automatically
func (l *CommonLoader) Load(key string) (interface{}, error) {
	return l.LoadThunk(key)()
}

// LoadThunk returns a function that when called will block waiting for a dict.
// This method should be used if you want one goroutine to make requests to many
// different data loaders without blocking until the thunk is called.
func (l *CommonLoader) LoadThunk(key string) func() (interface{}, error) {
	l.mu.Lock()
	if it, ok := l.cache[key]; ok {
		l.mu.Unlock()
		return func() (interface{}, error) {
			return it, nil
		}
	}
	if l.batch == nil {
		l.batch = &commonBatch{done: make(chan struct{})}
	}
	batch := l.batch
	pos := batch.keyIndex(l, key)
	l.mu.Unlock()

	return func() (interface{}, error) {
		<-batch.done

		var data interface{}
		if pos < len(batch.data) {
			data = batch.data[pos]
		}

		var err error
		// its convenient to be able to return a single error for everything
		if len(batch.error) == 1 {
			err = batch.error[0]
		} else if batch.error != nil {
			err = batch.error[pos]
		}

		if err == nil {
			l.mu.Lock()
			l.unsafeSet(key, data)
			l.mu.Unlock()
		}

		return data, err
	}
}

func (l *CommonLoader) unsafeSet(key string, value interface{}) {
	if l.cache == nil {
		l.cache = map[string]interface{}{}
	}
	l.cache[key] = value
}

// keyIndex will return the location of the key in the batch, if its not found
// it will add the key to the batch
func (b *commonBatch) keyIndex(l *CommonLoader, key string) int {
	for i, existingKey := range b.keys {
		if key == existingKey {
			return i
		}
	}

	pos := len(b.keys)
	b.keys = append(b.keys, key)
	if pos == 0 {
		go b.startTimer(l)
	}

	if l.maxBatch != 0 && pos >= l.maxBatch-1 {
		if !b.closing {
			b.closing = true
			l.batch = nil
			go b.end(l)
		}
	}

	return pos
}

func (b *commonBatch) startTimer(l *CommonLoader) {
	time.Sleep(l.wait)
	l.mu.Lock()

	// we must have hit a batch limit and are already finalizing this batch
	if b.closing {
		l.mu.Unlock()
		return
	}

	l.batch = nil
	l.mu.Unlock()

	b.end(l)
}

func (b *commonBatch) end(l *CommonLoader) {
	b.data, b.error = l.fetch(b.keys)
	close(b.done)
}
