package loaders

import (
	"fmt"
	"github.com/zhanghup/go-tools/database/toolxorm"
	"reflect"
	"sync"
	"time"
)

type ObjectLoader struct {
	sql      string
	param    map[string]interface{}
	db       *toolxorm.Engine
	keyField string
	table    interface{}

	cache map[string]interface{}
	batch *objectLoaderBatch
	sync  sync.Locker
}

type objectLoaderBatch struct {
	keys    []string
	data    []interface{}
	error   error
	closing bool
	done    chan struct{}
}

func (this *ObjectLoader) fetch(keys []string) ([]interface{}, error) {
	this.sync.Lock()
	query := map[string]interface{}{}
	for k, v := range this.param {
		query[k] = v
	}
	query["keys"] = keys
	this.sync.Unlock()
	this.db.SF(this.sql, query)
	fmt.Println("------------------------------=================", keys)
	return nil, nil
}

func (l *ObjectLoader) Load(key string, result interface{}) error {
	i, err := l.LoadThunk(key)()
	if err != nil {
		return err
	}
	if i == nil {
		return nil
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(i).Elem())
	return nil
}

func (l *ObjectLoader) LoadThunk(key string) func() (interface{}, error) {
	l.sync.Lock()
	if it, ok := l.cache[key]; ok {
		l.sync.Unlock()
		return func() (interface{}, error) {
			return it, nil
		}
	}
	if l.batch == nil {
		l.batch = &objectLoaderBatch{done: make(chan struct{})}
	} else if l.batch.closing {
		l.batch.keys = nil
		l.batch.data = nil
		l.batch.error = nil
		l.batch.closing = false
		l.batch.done = make(chan struct{})
	}
	batch := l.batch
	pos := batch.keyIndex(l, key)
	l.sync.Unlock()

	return func() (interface{}, error) {
		<-batch.done

		var data interface{}
		if pos < len(batch.data) {
			data = batch.data[pos]
		}

		if batch.error == nil {
			l.sync.Lock()
			l.unsafeSet(key, data)
			l.sync.Unlock()
		}

		return data, batch.error
	}
}

func (l *ObjectLoader) unsafeSet(key string, value interface{}) {
	if l.cache == nil {
		l.cache = map[string]interface{}{}
	}
	l.cache[key] = value
}

func (b *objectLoaderBatch) keyIndex(l *ObjectLoader, key string) int {
	fmt.Println(b.keys, "keyskeyskeyskeyskeyskeyskeyskeyskeys")
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

	return pos
}

func (b *objectLoaderBatch) startTimer(l *ObjectLoader) {
	time.Sleep(time.Millisecond * 5)
	l.sync.Lock()

	if b.closing {
		l.sync.Unlock()
		return
	}

	l.sync.Unlock()
	b.end(l)
}

func (b *objectLoaderBatch) end(l *ObjectLoader) {
	b.data, b.error = l.fetch(b.keys)
	close(b.done)
	b.closing = true

}
