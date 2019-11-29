package initia

import (
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/cfg"
	"github.com/zhanghup/go-tools"
)

func InitDictCode(dict beans.Dict, dictitems []beans.DictItem) {
	hisDict := beans.Dict{}
	ok, err := cfg.DB().Engine().Table(hisDict).Where("code = ?", dict.Code).Get(&hisDict)
	if err != nil {
		panic(err)
	}

	// 若不存在当前字典项，则新增
	if !ok {
		dict.Id = tools.ObjectString()
		dict.Status = tools.Ptr().Int(1)
		_, err = cfg.DB().Engine().Table(dict).Insert(dict)
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
			_, err = cfg.DB().Engine().Table(o).Insert(o)
			if err != nil {
				panic(err)
			}
		}
	}

}

func initDict() {
	initDictSys()
	initDictSta()
	if cfg.WxqyEnable() {
		initWxqy()
	}
}

func initDictSys() {
	InitDictCode(beans.Dict{Code: tools.Ptr().String("SYS0001"), Name: tools.Ptr().String("用户类型")}, []beans.DictItem{
		{Name: tools.Ptr().String("平台用户"), Value: tools.Ptr().String("0")},
		{Name: tools.Ptr().String("微信公众号用户"), Value: tools.Ptr().String("1")},
		{Name: tools.Ptr().String("微信小程序用户"), Value: tools.Ptr().String("2")},
		{Name: tools.Ptr().String("微信企业号用户"), Value: tools.Ptr().String("3")},
	})

	InitDictCode(beans.Dict{Code: tools.Ptr().String("SYS0002"), Name: tools.Ptr().String("对象权限")}, []beans.DictItem{
		{Name: tools.Ptr().String("新增"), Value: tools.Ptr().String("C")}, // 就是root用户
		{Name: tools.Ptr().String("查询"), Value: tools.Ptr().String("R")},
		{Name: tools.Ptr().String("编辑"), Value: tools.Ptr().String("U")},
		{Name: tools.Ptr().String("删除"), Value: tools.Ptr().String("D")},
		{Name: tools.Ptr().String("管理"), Value: tools.Ptr().String("M")},
	})

	InitDictCode(beans.Dict{Code: tools.Ptr().String("SYS0003"), Name: tools.Ptr().String("对象列表")}, []beans.DictItem{
		{Name: tools.Ptr().String("用户"), Value: tools.Ptr().String("user")},
		{Name: tools.Ptr().String("数据字典"), Value: tools.Ptr().String("dict")},
	})
}

func initDictSta() {
	InitDictCode(beans.Dict{Code: tools.Ptr().String("STA0001"), Name: tools.Ptr().String("数据状态")}, []beans.DictItem{
		{Name: tools.Ptr().String("启用"), Value: tools.Ptr().String("1")},
		{Name: tools.Ptr().String("禁用"), Value: tools.Ptr().String("0")},
	})

	InitDictCode(beans.Dict{Code: tools.Ptr().String("STA0002"), Name: tools.Ptr().String("人物性别")}, []beans.DictItem{
		{Name: tools.Ptr().String("未知"), Value: tools.Ptr().String("0")},
		{Name: tools.Ptr().String("男"), Value: tools.Ptr().String("1")},
		{Name: tools.Ptr().String("女"), Value: tools.Ptr().String("2")},
	})
}

func initWxqy() {
	InitDictCode(beans.Dict{Code: tools.Ptr().String("WXQY001"), Name: tools.Ptr().String("菜单类型")}, []beans.DictItem{
		{Name: tools.Ptr().String("点击推事件"), Value: tools.Ptr().String("click")},
		{Name: tools.Ptr().String("跳转URL"), Value: tools.Ptr().String("view")},
		{Name: tools.Ptr().String("扫码推事件"), Value: tools.Ptr().String("scancode_push")},
		{Name: tools.Ptr().String("扫码推事件且弹出“消息接收中”提示框"), Value: tools.Ptr().String("scancode_waitmsg")},
		{Name: tools.Ptr().String("弹出系统拍照发图"), Value: tools.Ptr().String("pic_sysphoto")},
		{Name: tools.Ptr().String("弹出拍照或者相册发图"), Value: tools.Ptr().String("pic_photo_or_album")},
		{Name: tools.Ptr().String("弹出企业微信相册发图器"), Value: tools.Ptr().String("pic_weixin")},
		{Name: tools.Ptr().String("弹出地理位置选择器"), Value: tools.Ptr().String("location_select")},
	})
}
