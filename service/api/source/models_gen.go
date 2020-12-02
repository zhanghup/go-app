// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package source

import (
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/service/event"
)

type Accounts struct {
	Total *int            `json:"total"`
	Data  []beans.Account `json:"data"`
}

type CronLogs struct {
	Total *int            `json:"total"`
	Data  []beans.CronLog `json:"data"`
}

type Crons struct {
	Total *int         `json:"total"`
	Data  []beans.Cron `json:"data"`
}

type Depts struct {
	Total *int         `json:"total"`
	Data  []beans.Dept `json:"data"`
}

type IPermObj struct {
	// 对象
	Object string `json:"object"`
	// 操作权限
	Mask string `json:"mask"`
}

type Message struct {
	Action  event.MsgAction `json:"action"`
	Message *beans.MsgInfo  `json:"message"`
}

type NewAccount struct {
	// 用户ID
	UID string `json:"uid"`
	// 账号类型 dict: SYS002
	Type string `json:"type"`
	// 用户名
	Username *string `json:"username"`
	// 密码
	Password *string `json:"password"`
	// 是否为管理员
	Admin *int `json:"admin"`
	// 是否为默认账户，默认账户可以在用户列表中可见并且维护
	Default *int `json:"default"`
	// 排序
	Weight *int `json:"weight"`
	// 状态{dict:STA001}
	Status *string `json:"status"`
}

type NewDept struct {
	// 组织类型{dict:BUS001}
	Type *string `json:"type"`
	// 组织代码
	Code *string `json:"code"`
	// 组织名称
	Name *string `json:"name"`
	// 组织头像
	Avatar *string `json:"avatar"`
	// 父级组织ID
	Pid *string `json:"pid"`
	// 备注
	Remark *string `json:"remark"`
	// 排序
	Weight *int `json:"weight"`
	// 状态{dict:STA001}
	Status *string `json:"status"`
}

type NewDict struct {
	// 字典类型{dict:SYS001}
	Type *string `json:"type"`
	// 字典编码
	Code *string `json:"code"`
	// 字典名称
	Name *string `json:"name"`
	// 备注
	Remark *string `json:"remark"`
	// 排序
	Weight *int `json:"weight"`
	// 状态{dict:STA001}
	Status *string `json:"status"`
}

type NewDictItem struct {
	// 字典id
	Code *string `json:"code"`
	// 名称
	Name *string `json:"name"`
	// 值
	Value *string `json:"value"`
	// 扩展
	Ext *string `json:"ext"`
	// 排序
	Weight *int `json:"weight"`
	// 状态{dict:STA001}
	Status *string `json:"status"`
}

type NewMsgConfirm struct {
	// 确认备注
	Remark *string `json:"remark"`
}

type NewRole struct {
	// 角色描述
	Name *string `json:"name"`
	// 角色名称
	Desc *string `json:"desc"`
	// 排序
	Weight *int `json:"weight"`
	// 状态{dict:STA001}
	Status *string `json:"status"`
}

type NewUser struct {
	User    map[string]interface{} `json:"user"`
	Account *NewAccount            `json:"account"`
}

type PermObj struct {
	// 对象
	Object string `json:"object"`
	// 操作权限
	Mask string `json:"mask"`
}

type QAccount struct {
	UID   string `json:"uid"`
	Index *int   `json:"index"`
	Size  *int   `json:"size"`
	Count *bool  `json:"count"`
}

type QCron struct {
	Keyword *string `json:"keyword"`
	Index   *int    `json:"index"`
	Size    *int    `json:"size"`
	Count   *bool   `json:"count"`
}

type QCronLog struct {
	Cron    string  `json:"cron"`
	Keyword *string `json:"keyword"`
	Index   *int    `json:"index"`
	Size    *int    `json:"size"`
	Count   *bool   `json:"count"`
}

type QDept struct {
	Pid   *string `json:"pid"`
	Index *int    `json:"index"`
	Size  *int    `json:"size"`
	Count *bool   `json:"count"`
}

type QDict struct {
	// 字典类型
	Type *string `json:"type"`
}

type QMsgInfo struct {
	// 接收者
	Receiver *string `json:"receiver"`
	// 消息类型{dict:SYS005}
	Type *string `json:"type"`
	// 消息级别{dict: SYS006}
	Level *string `json:"level"`
	// 消息接收平台{dict:SYS007}
	Target *string `json:"target"`
	// 弹出消息是否必须确认{dict:STA005}
	MustConfirm *string `json:"must_confirm"`
	// 确认平台{dict:SYS007}
	ConfirmTarget *string `json:"confirm_target"`
	// 已读平台{dict:SYS007}
	ReadTarget *string `json:"read_target"`
	// 消息状态{ dict:SYS008}
	State *string `json:"state"`
	Index *int    `json:"index"`
	Size  *int    `json:"size"`
}

type QMsgTemplate struct {
	// 名称模糊查询
	Name *string `json:"name"`
	// 编码模糊查询
	Code *string `json:"code"`
}

type QMyMsgInfo struct {
	// 消息类型{dict:SYS005}
	Type *string `json:"type"`
	// 消息级别{dict: SYS006}
	Level *string `json:"level"`
	// 消息接收平台{dict:SYS007}
	Target *string `json:"target"`
	// 弹出消息是否必须确认{dict:STA005}
	MustConfirm *string `json:"must_confirm"`
	// 确认平台{dict:SYS007}
	ConfirmTarget *string `json:"confirm_target"`
	// 已读平台{dict:SYS007}
	ReadTarget *string `json:"read_target"`
	// 消息状态{ dict:SYS008}
	State *string `json:"state"`
	Index *int    `json:"index"`
	Size  *int    `json:"size"`
}

type QRole struct {
	Keyword *string `json:"keyword"`
	Index   *int    `json:"index"`
	Size    *int    `json:"size"`
	Count   *bool   `json:"count"`
}

type QUser struct {
	Keyword *string `json:"keyword"`
	// 获取当前权限下的用户
	Withrole *bool `json:"withrole"`
	// 状态查询[-1:全部,0:禁止,1:启用]
	Status *int  `json:"status"`
	Index  *int  `json:"index"`
	Size   *int  `json:"size"`
	Count  *bool `json:"count"`
}

type Roles struct {
	Total *int         `json:"total"`
	Data  []beans.Role `json:"data"`
}

type UpdAccount struct {
	// 账号类型 dict: SYS002
	Type string `json:"type"`
	// 用户名
	Username *string `json:"username"`
	// 密码
	Password *string `json:"password"`
	// 是否为管理员
	Admin *int `json:"admin"`
	// 是否为默认账户，默认账户可以在用户列表中可见并且维护
	Default *int `json:"default"`
	// 排序
	Weight *int `json:"weight"`
	// 状态{dict:STA001}
	Status *string `json:"status"`
}

type UpdDept struct {
	// 组织类型{dict:BUS001}
	Type *string `json:"type"`
	// 组织代码
	Code *string `json:"code"`
	// 组织名称
	Name *string `json:"name"`
	// 组织头像
	Avatar *string `json:"avatar"`
	// 父级组织ID
	Pid *string `json:"pid"`
	// 备注
	Remark *string `json:"remark"`
	// 排序
	Weight *int `json:"weight"`
	// 状态{dict:STA001}
	Status *string `json:"status"`
}

type UpdDict struct {
	// 字典类型{dict:SYS001}
	Type *string `json:"type"`
	// 字典名称
	Name *string `json:"name"`
	// 备注
	Remark *string `json:"remark"`
	// 排序
	Weight *int `json:"weight"`
	// 状态{dict:STA001}
	Status *string `json:"status"`
}

type UpdDictItem struct {
	// 名称
	Name *string `json:"name"`
	// 值
	Value *string `json:"value"`
	// 扩展
	Ext *string `json:"ext"`
	// 排序
	Weight *int `json:"weight"`
	// 状态{dict:STA001}
	Status *string `json:"status"`
}

type UpdMsgTemplate struct {
	// 模板名称
	Name *string `json:"name"`
	// 消息分类{dict:SYS005}
	Type *string `json:"type"`
	// 消息等级{dict:SYS006}
	Level *string `json:"level"`
	// 消息推送平台{dict:SYS007}
	Target *string `json:"target"`
	// 消息超时时间（秒）
	Expire *int64 `json:"expire"`
	// 消息是否必须确认{dict:STA005}
	MustConfirm *string `json:"must_confirm"`
	// 消息提示图片
	ImgPath *string `json:"img_path"`
	// 备注
	Remark *string `json:"remark"`
}

type UpdRole struct {
	// 角色描述
	Name *string `json:"name"`
	// 角色名称
	Desc *string `json:"desc"`
	// 排序
	Weight *int `json:"weight"`
	// 状态{dict:STA001}
	Status *string `json:"status"`
}

type UpdUser struct {
	User    map[string]interface{} `json:"user"`
	Account *UpdAccount            `json:"account"`
}

type Users struct {
	Total *int         `json:"total"`
	Data  []beans.User `json:"data"`
}
