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
	Object(table interface{}, sql string, param map[string]interface{}, keyField string, resultField string) *ObjectLoader
	Slice(table interface{}, sql string, param map[string]interface{}, keyField string, resultField string) *SliceLoader
}

type dataLoaden struct {
	db *xorm.Engine

	objSync  sync.Locker
	objStore map[string]*ObjectLoader

	sliceSync  sync.Locker
	sliceStore map[string]*SliceLoader
}

func NewDataLoaden(db *xorm.Engine) Loader {
	return &dataLoaden{
		db:       db,
		objSync:  &sync.Mutex{},
		objStore: map[string]*ObjectLoader{},

		sliceSync:  &sync.Mutex{},
		sliceStore: map[string]*SliceLoader{},
	}
}

func (this *dataLoaden) Object(table interface{}, sql string, param map[string]interface{}, keyField string, resultField string) *ObjectLoader {
	requestTable := reflect.TypeOf(table)
	if requestTable.Kind() == reflect.Ptr {
		requestTable = requestTable.Elem()
	}
	if requestTable.Kind() == reflect.Array || requestTable.Kind() == reflect.Slice {
		requestTable = requestTable.Elem()
		if requestTable.Kind() == reflect.Ptr {
			requestTable = requestTable.Elem()
		}
		if requestTable.Kind() != reflect.Struct {
			panic("table 数据结构异常")
		}
	}
	resultTable := requestTable
	if len(resultField) > 0 {
		f, ok := resultTable.FieldByName(resultField)
		if !ok {
			panic("resultField 不存在")
		}
		resultTable = f.Type
	}

	query := map[string]interface{}{}
	if param != nil {
		query = param
	}
	path := reflect.TypeOf(tools.Rft.RealValue(table)).PkgPath()
	name := reflect.TypeOf(tools.Rft.RealValue(table)).Name()
	key := fmt.Sprintf("table: %s/%s, sql: %s, param: %s, field: %s,resultField: %s", path, name, sql, tools.Str.JSONString(query), keyField, resultField)
	key = tools.Crypto.MD5([]byte(key))
	this.objSync.Lock()
	defer this.objSync.Unlock()
	if l, ok := this.objStore[key]; ok {
		return l
	}
	objLoader := &ObjectLoader{
		sync:     &sync.RWMutex{},
		db:       toolxorm.NewEngine(this.db),
		keyField:     keyField,
		resultField:  resultField,
		sql:          sql,
		param:        query,
		requestTable: requestTable,
		resultTable:  resultTable,
	}
	this.objStore[key] = objLoader
	return objLoader
}

func (this *dataLoaden) Slice(table interface{}, sql string, param map[string]interface{}, keyField string, resultField string) *SliceLoader {
	requestTable := reflect.TypeOf(table)
	if requestTable.Kind() == reflect.Ptr {
		requestTable = requestTable.Elem()
	}
	if requestTable.Kind() == reflect.Array || requestTable.Kind() == reflect.Slice {
		requestTable = requestTable.Elem()
		if requestTable.Kind() == reflect.Ptr {
			requestTable = requestTable.Elem()
		}
		if requestTable.Kind() != reflect.Struct {
			panic("table 数据结构异常")
		}
	}
	resultTable := requestTable
	if len(resultField) > 0 {
		f, ok := resultTable.FieldByName(resultField)
		if !ok {
			panic("resultField 不存在")
		}
		resultTable = f.Type
	}

	query := map[string]interface{}{}
	if param != nil {
		query = param
	}

	path := reflect.TypeOf(tools.Rft.RealValue(table)).PkgPath()
	name := reflect.TypeOf(tools.Rft.RealValue(table)).Name()
	key := fmt.Sprintf("table: %s/%s, sql: %s, param: %s, field: %s,resultField: %s,resultTable: %s", path, name, sql, tools.Str.JSONString(query), keyField, resultField, resultTable.PkgPath()+"/"+resultTable.Name())
	key = tools.Crypto.MD5([]byte(key))

	this.sliceSync.Lock()
	defer this.sliceSync.Unlock()
	if l, ok := this.sliceStore[key]; ok {
		return l
	}
	sliceLoader := &SliceLoader{
		sync:         &sync.RWMutex{},
		db:           toolxorm.NewEngine(this.db),
		keyField:     keyField,
		resultField:  resultField,
		sql:          sql,
		param:        query,
		requestTable: requestTable,
		resultTable:  resultTable,
	}
	this.sliceStore[key] = sliceLoader
	return sliceLoader
}
