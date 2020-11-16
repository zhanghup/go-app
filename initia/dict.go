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
	Name     string
	Value    string
	Ext      string
	Disabled int
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
					Status: tools.Ptr.String("1"),
				},
				Code: tools.Ptr.String(ty + dict.Code),
				Name: &dict.Name,
				Type: &ty,
			})
			if err != nil {
				panic(err)
			}
		} else {
			continue
		}

		// 只有不存在字典项的时候，新增字典项，然后增加具体条目
		if len(dict.Children) > 0 {
			for i, o := range dict.Children {
				_, err = db.Table(beans.DictItem{}).Insert(beans.DictItem{
					Bean: beans.Bean{
						Id:     tools.Ptr.Uid(),
						Weight: &i,
						Status: tools.Ptr.String("1"),
					},
					Code:     tools.Ptr.String(ty + dict.Code),
					Name:     &o.Name,
					Value:    &o.Value,
					Disabled: &o.Disabled,
					Ext:      &o.Ext,
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
	initDictBus(db)
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
		{Code: "001", Name: "字典类型", Children: []DictInfoItem{
			{Name: "系统类型", Value: "SYS", Disabled: 1},
			{Name: "系统状态", Value: "STA", Disabled: 1},
			{Name: "业务类型", Value: "BUS", Disabled: 1},
		}},
		{Code: "002", Name: "账号类型", Children: []DictInfoItem{
			{Name: "用户密码", Value: "password", Disabled: 1},
		}},
		{Code: "003", Name: "权限类型", Children: []DictInfoItem{
		}},
		{Code: "004", Name: "权限状态", Children: []DictInfoItem{
			{Name: "新增", Value: "C", Disabled: 1},
			{Name: "查询", Value: "R", Disabled: 1},
			{Name: "编辑", Value: "U", Disabled: 1},
			{Name: "删除", Value: "D", Disabled: 1},
			{Name: "管理", Value: "M", Disabled: 1},
		}},
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
			{Name: "启用", Value: "1", Disabled: 1},
			{Name: "禁用", Value: "0", Disabled: 1},
		}},
		{Code: "002", Name: "人物性别", Children: []DictInfoItem{
			{Name: "男", Value: "1", Disabled: 1},
			{Name: "女", Value: "2", Disabled: 1},
			{Name: "未知", Value: "3", Disabled: 1},
		}},
		{Code: "003", Name: "运行状态", Children: []DictInfoItem{
			{Name: "开始", Value: "start", Disabled: 1},
			{Name: "停止", Value: "stop", Disabled: 1},
		}},
		{Code: "004", Name: "执行结果", Children: []DictInfoItem{
			{Name: "成功", Value: "success", Disabled: 1},
			{Name: "失败", Value: "error", Disabled: 1},
			{Name: "拒绝", Value: "refuse", Disabled: 1},
		}},
	})
}

func initDictBus(db *xorm.Engine) {
	err := txorm.NewEngine(db).SF(`delete from dict_item where code in (select code from dict where type = 'BUS')`).Exec()
	if err != nil {
		tog.Error(err.Error())
	}
	err = txorm.NewEngine(db).SF(`delete from dict where type = 'BUS'`).Exec()
	if err != nil {
		tog.Error(err.Error())
	}
	InitDictCode(db, "BUS", []DictInfo{
		{Code: "001", Name: "组织类型", Children: []DictInfoItem{
		}},
		{Code: "002", Name: "用户类型", Children: []DictInfoItem{
		}},
	})
}
