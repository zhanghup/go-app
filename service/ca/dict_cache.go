package ca

import (
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/service/event"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/tog"
	"xorm.io/xorm"
)

type dictcacheinfo struct {
	Dict     beans.Dict
	DictItem []beans.DictItem
}

type dictCache struct {
	data tools.ICache
	db   *xorm.Engine
}

var DictCache *dictCache

func (this *dictCache) init() error {
	dicts := make([]beans.Dict, 0)
	dictitems := make([]beans.DictItem, 0)
	err := this.db.Find(&dicts)
	if err != nil {
		return err
	}
	err = this.db.Find(&dictitems)
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
	v := o.(dictcacheinfo)
	return &v.Dict, v.DictItem, true
}

func init() {
	event.XormDefaultInitSubscribeOnce(func(db *xorm.Engine) {
		if DictCache != nil {
			return
		}
		DictCache = &dictCache{
			db:   db,
			data: tools.CacheCreate(),
		}

		err := DictCache.init()
		if err != nil {
			tog.Error(err.Error())
		}
		go event.DictChangeSubscribe(func() {
			err := DictCache.init()
			if err != nil {
				tog.Error(err.Error())
			}
		})
	})
	return
}
