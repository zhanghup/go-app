package initia

import (
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-tools"
	"xorm.io/xorm"
)

func NewDict(db *xorm.Engine) *Dict {
	return &Dict{db}
}

type Dict struct {
	DB *xorm.Engine
}

func (this *Dict) initDictCode(dict beans.Dict, dictitems []beans.DictItem) {
	hisDict := beans.Dict{}
	ok, err := this.DB.Table(hisDict).Where("code = ?", dict.Code).Get(&hisDict)
	if err != nil {
		panic(err)
	}

	// 若不存在当前字典项，则新增
	if !ok {
		dict.Id = tools.Ptr.Uid()
		dict.Status = tools.Ptr.Int(1)
		_, err = this.DB.Table(dict).Insert(dict)
		if err != nil {
			panic(err)
		}
	} else {
		return
	}

	// 只有不存在字典项的时候，新增字典项，然后增加具体条目
	if len(dictitems) > 0 {
		for i, o := range dictitems {
			o.Weight = &i
			o.Status = tools.Ptr.Int(1)
			o.Id = tools.Ptr.Uid()
			o.Code = dict.Id
			_, err = this.DB.Table(o).Insert(o)
			if err != nil {
				panic(err)
			}
		}
	}
}

func (this *Dict) InitDict() {
	this.initDictSys()
	this.initDictSta()
}

func (this *Dict) initDictSys() {
	this.initDictCode(beans.Dict{Code: tools.Ptr.String("SYS0001"), Name: tools.Ptr.String("用户类型")}, []beans.DictItem{
		{Name: tools.Ptr.String("平台用户"), Value: tools.Ptr.String("0")},
	})

	this.initDictCode(beans.Dict{Code: tools.Ptr.String("SYS0002"), Name: tools.Ptr.String("对象权限")}, []beans.DictItem{
		{Name: tools.Ptr.String("新增"), Value: tools.Ptr.String("C")}, // 就是root用户
		{Name: tools.Ptr.String("查询"), Value: tools.Ptr.String("R")},
		{Name: tools.Ptr.String("编辑"), Value: tools.Ptr.String("U")},
		{Name: tools.Ptr.String("删除"), Value: tools.Ptr.String("D")},
		{Name: tools.Ptr.String("管理"), Value: tools.Ptr.String("M")},
	})

	this.initDictCode(beans.Dict{Code: tools.Ptr.String("SYS0003"), Name: tools.Ptr.String("对象列表")}, []beans.DictItem{
		{Name: tools.Ptr.String("用户"), Value: tools.Ptr.String("user")},
		{Name: tools.Ptr.String("数据字典"), Value: tools.Ptr.String("dict")},
	})

	this.initDictCode(beans.Dict{Code: tools.Ptr.String("SYS0004"), Name: tools.Ptr.String("表单字段类型")}, []beans.DictItem{
		{Name: tools.Ptr.String("隐藏域"), Value: tools.Ptr.String("hide")},
		{Name: tools.Ptr.String("字符串"), Value: tools.Ptr.String("string")},
		{Name: tools.Ptr.String("数字"), Value: tools.Ptr.String("number")},
		{Name: tools.Ptr.String("时间"), Value: tools.Ptr.String("date")},
		{Name: tools.Ptr.String("时间范围"), Value: tools.Ptr.String("date-range")},
		{Name: tools.Ptr.String("经纬度"), Value: tools.Ptr.String("position")},
		{Name: tools.Ptr.String("当前时间"), Value: tools.Ptr.String("current-time")},
		{Name: tools.Ptr.String("当前人员"), Value: tools.Ptr.String("current-user")},
	})

}

func (this *Dict) initDictSta() {
	this.initDictCode(beans.Dict{Code: tools.Ptr.String("STA0001"), Name: tools.Ptr.String("数据状态")}, []beans.DictItem{
		{Name: tools.Ptr.String("启用"), Value: tools.Ptr.String("1")},
		{Name: tools.Ptr.String("禁用"), Value: tools.Ptr.String("0")},
	})

	this.initDictCode(beans.Dict{Code: tools.Ptr.String("STA0002"), Name: tools.Ptr.String("人物性别")}, []beans.DictItem{
		{Name: tools.Ptr.String("未知"), Value: tools.Ptr.String("0")},
		{Name: tools.Ptr.String("男"), Value: tools.Ptr.String("1")},
		{Name: tools.Ptr.String("女"), Value: tools.Ptr.String("2")},
	})
}
