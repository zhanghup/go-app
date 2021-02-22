package initia

import (
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/tog"
	"xorm.io/xorm"
)

func InitDict(db *xorm.Engine) {
	// SYS
	{
		// 字典类型 SYS001
		InitDictCode(db, "SYS", "001", "字典类型", 1, []DictInfoItem{
			{"1", "系统类型", "SYS", "", 1},
			{"2", "系统状态", "STA", "", 1},
			{"3", "业务类型", "BUS", "", 1},
		})

		// 账号类型 SYS002
		InitDictCode(db, "SYS", "002", "账号类型", 1, []DictInfoItem{
			{"1", "用户密码", "password", "", 1},
		})

		// 权限类型 SYS003
		InitDictCode(db, "SYS", "003", "权限类型", 1, []DictInfoItem{
			{"1", "菜单权限", "menu", "", 1},
		})

		// 权限状态 SYS004
		//InitDictCode(db, "SYS", "004", "权限状态", 1, []DictInfoItem{
		//	{"1", "新增", "C", "", 1},
		//	{"2", "查询", "R", "", 1},
		//	{"3", "编辑", "U", "", 1},
		//	{"4", "删除", "D", "", 1},
		//	{"5", "管理", "M", "", 1},
		//})

		// 权限类型 SYS005
		InitDictCode(db, "SYS", "005", "消息模板类型", 1, []DictInfoItem{
			{"1", "确认框", "confirm", "", 1},
			{"2", "消息提示", "message", "", 1},
			{"3", "通知", "notice", "", 1},
		})

		// 消息等级 SYS006
		InitDictCode(db, "SYS", "006", "消息等级", 1, []DictInfoItem{
			{"1", "严重", "0", "", 1},
			{"2", "重要", "1", "", 1},
			{"3", "次要", "2", "", 1},
			{"4", "普通", "3", "", 1},
		})

		// 消息推送目标 SYS007
		InitDictCode(db, "SYS", "007", "消息推送目标", 1, []DictInfoItem{
			{"1", "网页", "web", "", 1},
		})

		// 消息状态 SYS008
		InitDictCode(db, "SYS", "008", "消息状态", 1, []DictInfoItem{
			{"1", "已读", "0", "", 1},
			{"2", "未读", "1", "", 1},
			{"5", "已确认", "4", "", 1},
		})

		// 业务对象 SYS009
		InitDictCode(db, "SYS", "009", "业务对象", 1, []DictInfoItem{
			{"0", "账户管理", "account", "C:新增,U:更新,D:删除", 1},
			{"1", "组织管理", "dept", "C:新增,U:更新,D:删除", 1},
			{"2", "用户管理", "user", "C:新增,U:更新,D:删除", 1},
			{"3", "数据字典", "dict", "C:新增,U:更新,D:删除", 1},
			{"4", "数据字典项", "dict_item", "C:新增,U:更新,D:删除,MST:排序", 1},
			{"5", "角色管理", "role", "C:新增,U:更新,D:删除,MPM:菜单权限,MO:对象权限,MWU:角色分配", 1},
			{"6", "定时任务", "cron", "MSO:暂停,MST:开启,MR:执行一次", 1},
			{"7", "菜单管理", "menu", "C:新增,U:菜单调整", 1},
		})

		// 菜单类型 SYS010
		InitDictCode(db, "SYS", "010", "菜单类型", 1, []DictInfoItem{
			{"1", "目录菜单", "0", "", 1},
			{"2", "路由菜单", "1", "", 1},
		})
	}

	// STA
	{
		// 数据状态 STA001
		InitDictCode(db, "STA", "001", "数据状态", 1, []DictInfoItem{
			{"1", "启用", "1", "", 1},
			{"2", "禁用", "0", "", 1},
		})

		// 人物性别 STA002
		InitDictCode(db, "STA", "002", "人物性别", 1, []DictInfoItem{
			{"1", "男", "1", "", 1},
			{"2", "女", "2", "", 1},
			{"3", "未知", "3", "", 1},
		})

		// 运行状态 STA003
		InitDictCode(db, "STA", "003", "运行状态", 1, []DictInfoItem{
			{"1", "开始", "start", "", 1},
			{"2", "停止", "stop", "", 1},
		})

		// 执行结果 STA004
		InitDictCode(db, "STA", "004", "执行结果", 1, []DictInfoItem{
			{"1", "成功", "success", "", 1},
			{"2", "失败", "error", "", 1},
			{"3", "拒绝", "refuse", "", 1},
		})

		// 是否 STA005
		InitDictCode(db, "STA", "005", "是否", 1, []DictInfoItem{
			{"1", "是", "1", "", 1},
			{"2", "否", "0", "", 1},
		})
	}

	// BUS
	{
		// 组织类型 BUS001
		InitDictCode(db, "BUS", "001", "组织类型", 0, []DictInfoItem{
			{"1", "普通组织", "1", "", 0},
		})

		// 用户类型 BUS001
		InitDictCode(db, "BUS", "002", "用户类型", 0, []DictInfoItem{
			{"1", "普通用户", "1", "", 0},
		})
	}

	// WXMP
	{
		// 用户性别 WXMP001
		InitDictCode(db, "WXMP", "001", "用户性别", 0, []DictInfoItem{
			{"1", "未知", "0", "", 0},
			{"2", "男性", "1", "", 0},
			{"3", "女性", "2", "", 0},
		})

		// 支付类型 BUS001
		InitDictCode(db, "WXMP", "002", "支付类型", 0, []DictInfoItem{
		})

		// 支付状态 WXMP001
		InitDictCode(db, "WXMP", "003", "支付状态", 0, []DictInfoItem{
			{"1", "未支付", "0", "", 0},
			{"2", "已支付", "1", "", 0},
			{"3", "已取消", "2", "", 0},
			{"4", "支付成功", "3", "", 0},
			{"5", "支付失败", "4", "", 0},
		})
	}
}

func InitDictItem(db *xorm.Engine, code, id, name, value, ext string, weight, disabled int) {
	itemid := code + "-" + id

	dictItem := beans.DictItem{}
	ok, err := db.Where("id = ?", itemid).Get(&dictItem)
	if err != nil {
		tog.Error(err.Error())
		return
	}

	newItem := beans.DictItem{
		Bean: beans.Bean{
			Id:     &itemid,
			Weight: &weight,
			Status: tools.PtrOfString("1"),
		},
		Code:     &code,
		Name:     &name,
		Value:    &value,
		Disabled: &disabled,
		Ext:      &ext,
	}

	if ok {
		_, err = db.Where("id = ?", itemid).Update(newItem)
		if err != nil {
			tog.Error(err.Error())
			return
		}
	} else {
		_, err = db.Insert(newItem)
		if err != nil {
			tog.Error(err.Error())
			return
		}
	}
}

func InitDictCode(db *xorm.Engine, typeArg, code, name string, disabled int, items []DictInfoItem) {
	hisDict := beans.Dict{}
	id := typeArg + code
	ok, err := db.Table(hisDict).Where("id = ?", id).Get(&hisDict)
	if err != nil {
		tog.Error(err.Error())
		return
	}
	if !ok {
		hisDict.Id = &id
		hisDict.Status = tools.PtrOfString("1")
		hisDict.Code = &id
		hisDict.Name = &name
		hisDict.Type = &typeArg
		hisDict.Disabled = &disabled
		_, err = db.Insert(hisDict)
		if err != nil {
			tog.Error(err.Error())
			return
		}
	} else {
		hisDict.Id = &id
		hisDict.Status = tools.PtrOfString("1")
		hisDict.Code = &id
		hisDict.Name = &name
		hisDict.Type = &typeArg
		hisDict.Disabled = &disabled
		_, err = db.Where("id = ?", id).Update(hisDict)
		if err != nil {
			tog.Error(err.Error())
			return
		}
	}

	for i, item := range items {
		if len(item.Id) == 0 {
			panic("DictInfoItem 的id属性必须被指定，并且为不重复数据")
		}
		itemid := id + "-" + item.Id

		dictItem := beans.DictItem{}
		ok, err := db.Where("id = ?", itemid).Get(&dictItem)
		if err != nil {
			tog.Error(err.Error())
			return
		}

		newItem := beans.DictItem{
			Bean: beans.Bean{
				Id:     &itemid,
				Weight: &i,
				Status: tools.PtrOfString("1"),
			},
			Code:     &id,
			Name:     &item.Name,
			Value:    &item.Value,
			Disabled: &item.Disabled,
			Ext:      &item.Ext,
		}

		if ok {
			_, err = db.Where("id = ?", itemid).Update(newItem)
			if err != nil {
				tog.Error(err.Error())
				return
			}
		} else {
			_, err = db.Insert(newItem)
			if err != nil {
				tog.Error(err.Error())
				return
			}
		}

	}
}

type DictInfoItem struct {
	Id       string
	Name     string
	Value    string
	Ext      string
	Disabled int
}
