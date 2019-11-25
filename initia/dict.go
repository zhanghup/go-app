package initia

import (
	"github.com/go-xorm/xorm"
	"github.com/zhanghup/go-app"
	"github.com/zhanghup/go-tools"
)

func InitDictCode(e *xorm.Engine, code, name, remark string, weight int, dictitems []app.DictItem) {
	dict := app.Dict{}
	ok, err := e.Table(dict).Where("code = ?", code).Get(&dict)
	if err != nil {
		panic(err)
	}

	// 若不存在当前字典项，则新增
	if !ok {
		dict.Code = &code
		dict.Name = &name
		dict.Remark = &remark
		dict.Weight = &weight
		dict.Id = tools.ObjectString()
		dict.Status = tools.Ptr().Int(1)
		_, err = e.Table(dict).Insert(dict)
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
			o.Status = tools.Ptr().Int(1)
			o.Id = tools.ObjectString()
			o.Code = dict.Id
			_, err = e.Table(o).Insert(o)
			if err != nil {
				panic(err)
			}
		}
	}

}

func initDict(e *xorm.Engine) {
	InitDictCode(e, "SYS0001", "用户类型", "", 1, []app.DictItem{
		{Name: tools.Ptr().String("超级管理员"), Value: tools.Ptr().String("0")}, // 就是root用户
		{Name: tools.Ptr().String("管理员"), Value: tools.Ptr().String("1")},
		{Name: tools.Ptr().String("普通用户"), Value: tools.Ptr().String("2")},
	})
	InitDictCode(e, "SYS0002", "对象权限", "", 1, []app.DictItem{
		{Name: tools.Ptr().String("新增"), Value: tools.Ptr().String("C")}, // 就是root用户
		{Name: tools.Ptr().String("查询"), Value: tools.Ptr().String("R")},
		{Name: tools.Ptr().String("编辑"), Value: tools.Ptr().String("U")},
		{Name: tools.Ptr().String("删除"), Value: tools.Ptr().String("D")},
		{Name: tools.Ptr().String("管理"), Value: tools.Ptr().String("M")},
	})
	InitDictCode(e, "SYS0003", "对象列表", "", 1, []app.DictItem{
		{Name: tools.Ptr().String("用户"), Value: tools.Ptr().String("user")},
		{Name: tools.Ptr().String("数据字典"), Value: tools.Ptr().String("dict")},
	})

	InitDictCode(e, "STA0001", "数据状态", "", 100, []app.DictItem{
		{Name: tools.Ptr().String("启用"), Value: tools.Ptr().String("1")},
		{Name: tools.Ptr().String("禁用"), Value: tools.Ptr().String("0")},
	})

	InitDictCode(e, "STA0002", "人物性别", "", 101, []app.DictItem{
		{Name: tools.Ptr().String("未知"), Value: tools.Ptr().String("0")},
		{Name: tools.Ptr().String("男"), Value: tools.Ptr().String("1")},
		{Name: tools.Ptr().String("女"), Value: tools.Ptr().String("2")},
	})
}
