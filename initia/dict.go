package initia

import (
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
	"github.com/zhanghup/go-tools/tog"
	"xorm.io/xorm"
)

type DictInfo struct {
	Code string
	Name string

	Children []DictInfoItem
}

type DictInfoItem struct {
	Name      string
	Value     string
	Extension string
	Disable   int
}

func InitDictCode(db *xorm.Engine, ty string, dicts []DictInfo) {
	for ii, dict := range dicts {
		hisDict := beans.Dict{}
		ok, err := db.Table(hisDict).Where("code = ?", ty+dict.Code).Get(&hisDict)
		if err != nil {
			panic(err)
		}

		// 若不存在当前字典项，则新增
		if !ok {
			_, err = db.Table(hisDict).Insert(beans.Dict{
				Bean: beans.Bean{
					Id:     tools.Ptr.Uid(),
					Weight: &ii,
					Status: tools.Ptr.Int(1),
				},
				Code: tools.Ptr.String(ty + dict.Code),
				Name: &dict.Name,
				Type: &ty,
			})
			if err != nil {
				panic(err)
			}
		} else {
			return
		}

		// 只有不存在字典项的时候，新增字典项，然后增加具体条目
		if len(dict.Children) > 0 {
			for i, o := range dict.Children {
				_, err = db.Table(beans.DictItem{}).Insert(beans.DictItem{
					Bean: beans.Bean{
						Id:     tools.Ptr.Uid(),
						Weight: &i,
						Status: tools.Ptr.Int(1),
					},
					Code:      tools.Ptr.String(ty + dict.Code),
					Name:      &o.Name,
					Value:     &o.Value,
					Disable:   &o.Disable,
					Extension: &o.Extension,
				})
				if err != nil {
					panic(err)
				}
			}
		}
	}
}

func InitDict(db *xorm.Engine) {
	initDictSys(db)
	initDictSta(db)
}

func initDictSys(db *xorm.Engine) {
	err := txorm.NewEngine(db).SF(`delete from dict_item where code in (select code from dict where type = 'SYS')`).Exec()
	if err != nil {
		tog.Error(err.Error())
	}
	err = txorm.NewEngine(db).SF(`delete from dict where type = 'SYS'`).Exec()
	if err != nil {
		tog.Error(err.Error())
	}
	InitDictCode(db, "SYS", []DictInfo{
		{Code: "000", Name: "字典类型", Children: []DictInfoItem{
			{Name: "系统类型", Value: "SYS", Disable: 1},
			{Name: "系统状态", Value: "STA", Disable: 1},
			{Name: "系统映射", Value: "AUT", Disable: 1},
		}},
		{Code: "001", Name: "用户类型", Children: []DictInfoItem{
			{Name: "平台用户", Value: "0", Disable: 1},
		}},
		{Code: "002", Name: "对象权限", Children: []DictInfoItem{
			{Name: "新增", Value: "C", Disable: 1},
			{Name: "查询", Value: "R", Disable: 1},
			{Name: "编辑", Value: "U", Disable: 1},
			{Name: "删除", Value: "D", Disable: 1},
			{Name: "管理", Value: "M", Disable: 1},
		}},
		{Code: "003", Name: "对象列表", Children: []DictInfoItem{
			{Name: "用户", Value: "user", Disable: 1},
			{Name: "数据字典", Value: "dict", Disable: 1},
		},},
	})
}

func initDictSta(db *xorm.Engine) {
	err := txorm.NewEngine(db).SF(`delete from dict_item where code in (select code from dict where type = 'STA')`).Exec()
	if err != nil {
		tog.Error(err.Error())
	}
	err = txorm.NewEngine(db).SF(`delete from dict where type = 'STA'`).Exec()
	if err != nil {
		tog.Error(err.Error())
	}
	InitDictCode(db, "STA", []DictInfo{
		{Code: "001", Name: "数据状态", Children: []DictInfoItem{
			{Name: "启用", Value: "1", Disable: 1},
			{Name: "禁用", Value: "0", Disable: 1,},
		}},
	})
}
