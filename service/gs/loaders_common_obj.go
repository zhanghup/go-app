package gs

import (
	"reflect"
	"sync"
	"time"
)

type CommonLoader struct {
	fetch    func(keys []string) ([]interface{}, []error)
	wait     time.Duration
	maxBatch int
	cache    map[string]interface{}
	batch    *commonBatch
	mu       sync.Mutex
}

type commonBatch struct {
	keys    []string
	data    []interface{}
	error   []error
	closing bool
	done    chan struct{}

	table interface{} // 保存数据库model对象
}

func (l *CommonLoader) Load(key string, result interface{}) (interface{}, error) {
	rtype := reflect.TypeOf(result)
	if rtype.Kind() != reflect.Ptr || rtype.Elem().Kind() != reflect.Struct {
		panic("传入参数result类型异常，应为*struct{}")
	}

	result, err := l.LoadThunk(key)()
	if err != nil {
		return nil, err
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(result))
	return nil, nil
}

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
