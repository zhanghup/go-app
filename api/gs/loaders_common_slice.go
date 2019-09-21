package gs

import (
	"reflect"
	"sync"
	"time"
)

type CommonSliceLoader struct {
	fetch    func(keys []string) (interface{}, []error)
	wait     time.Duration
	maxBatch int
	cache    interface{}
	batch    *commonSliceLoaderBatch
	mu       sync.Mutex
}

type commonSliceLoaderBatch struct {
	keys    []string
	data    interface{}
	error   []error
	closing bool
	done    chan struct{}

	table interface{} // 保存数据库model对象
}

/*
 @param result 的类型 *[]*interface{}
*/

func (l *CommonSliceLoader) Load(key string, result interface{}) error {
	datas, err := l.LoadThunk(key)()
	if err != nil {
		return err
	}
	if datas == nil {
		return nil
	}

	rtype := reflect.TypeOf(result)
	if rtype.Kind() != reflect.Ptr || rtype.Elem().Kind() != reflect.Slice {
		panic("传入参数result类型异常，应为*[]*interface{}")
	}

	rval := reflect.New(rtype.Elem()).Elem()

	vl := reflect.ValueOf(datas)
	if vl.Type().Kind() == reflect.Slice {
		for i := 0; i < vl.Len(); i++ {
			rval = reflect.Append(rval, vl.Index(i))
		}
	}
	reflect.ValueOf(result).Elem().Set(rval)
	return nil
}

func (l *CommonSliceLoader) LoadThunk(key string) func() (interface{}, error) {
	l.mu.Lock()

	if l.cache != nil {
		if it := reflect.ValueOf(l.cache).MapIndex(reflect.ValueOf(key)); it.IsValid() {
			l.mu.Unlock()
			return func() (interface{}, error) {
				return it.Interface(), nil
			}
		}
	}
	if l.batch == nil {
		l.batch = &commonSliceLoaderBatch{done: make(chan struct{})}
	}
	batch := l.batch
	pos := batch.keyIndex(l, key)
	l.mu.Unlock()

	return func() (interface{}, error) {
		<-batch.done
		if batch.data == nil {
			return nil, nil
		}

		var data interface{}
		rg := reflect.ValueOf(batch.data).MapRange()
		for ; rg.Next(); {
			v := rg.Value()
			k := rg.Key()
			if k.String() == key {
				data = v.Interface()
			}
		}

		if data == nil {
			return nil, nil
		}

		var err error
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

//func (l *CommonSliceLoader) Prime(key string, value []interface{}) bool {
//	l.mu.Lock()
//	var found bool
//	if _, found = l.cache[key]; !found {
//		cpy := make([]interface{}, len(value))
//		copy(cpy, value)
//		l.unsafeSet(key, cpy)
//	}
//	l.mu.Unlock()
//	return !found
//}

func (l *CommonSliceLoader) unsafeSet(key string, value interface{}) {
	var cache reflect.Value
	if l.cache == nil {
		cache = reflect.MakeMap(reflect.MapOf(reflect.TypeOf(key), reflect.TypeOf(value)))
	} else {
		cache = reflect.ValueOf(l.cache)
	}

	cache.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(value))
	l.cache = cache.Interface()
}

func (b *commonSliceLoaderBatch) keyIndex(l *CommonSliceLoader, key string) int {
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

func (b *commonSliceLoaderBatch) startTimer(l *CommonSliceLoader) {
	time.Sleep(l.wait)
	l.mu.Lock()

	if b.closing {
		l.mu.Unlock()
		return
	}

	l.batch = nil
	l.mu.Unlock()

	b.end(l)
}

func (b *commonSliceLoaderBatch) end(l *CommonSliceLoader) {
	b.data, b.error = l.fetch(b.keys)
	close(b.done)
}
