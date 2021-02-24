package ca

import (
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/gs"
	"github.com/zhanghup/go-app/service/event"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/tog"
)

type dictcacheinfo struct {
	Dict     beans.Dict
	DictItem []beans.DictItem
}

type dictCache struct {
	data tools.ICache
}

var DictCache *dictCache

func (this *dictCache) init() error {
	dicts := make([]beans.Dict, 0)
	dictitems := make([]beans.DictItem, 0)
	err := gs.DB().Find(&dicts)
	if err != nil {
		return err
	}
	err = gs.DB().Find(&dictitems)
	if err != nil {
		return err
	}
	dictmap := map[string]*dictcacheinfo{}
	for _, o := range dicts {
		if o.Id == nil {
			continue
		}
		dictmap[*o.Id] = &dictcacheinfo{Dict: o}
	}

	for _, o := range dictitems {
		if o.Code == nil {
			continue
		}
		dict, ok := dictmap[*o.Code]
		if !ok {
			continue
		}
		dict.DictItem = append(dict.DictItem, o)
		dictmap[*o.Code] = dict
	}

	for _, v := range dictmap {
		if v.Dict.Code == nil {
			continue
		}
		this.data.Set(*v.Dict.Code, v)
	}
	return nil
}

func (this *dictCache) Get(dict string) (*beans.Dict, []beans.DictItem, bool) {
	o := this.data.Get(dict)
	if o == nil {
		return nil, nil, false
	}
	v := o.(*dictcacheinfo)
	return &v.Dict, v.DictItem, true
}

func (this *dictCache) GetName(dictCode, value string) string {
	_, items, ok := this.Get(dictCode)
	if !ok {
		return ""
	}
	if items == nil {
		return ""
	}
	for _, s := range items {
		if s.Value != nil && *s.Value == value {
			return *s.Name
		}
	}
	return ""
}

func (this *dictCache) GetValue(dictCode, name string) string {
	_, items, ok := this.Get(dictCode)
	if !ok {
		return ""
	}
	if items == nil {
		return ""
	}
	for _, s := range items {
		if s.Name != nil && *s.Name == name {
			return *s.Value
		}
	}
	return ""
}

func initDict() {
	gs.InfoBegin("数据字典缓存")
	if DictCache != nil {
		return
	}
	DictCache = &dictCache{
		data: tools.CacheCreate(),
	}

	err := DictCache.init()
	if err != nil {
		tog.Error(err.Error())
		gs.InfoError("数据字典缓存")
	} else {
		gs.InfoSuccess("数据字典缓存")
	}
	go event.DictChangeSubscribe(func() {
		err := DictCache.init()
		if err != nil {
			tog.Error(err.Error())
		}
	})
}
