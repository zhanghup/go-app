// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package source

import (
	"github.com/zhanghup/go-app/beans"
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

type MenuLocal struct {
	ID *string `json:"id"`
	// 菜单类型
	Type *string `json:"type"`
	// 菜单名称
	Name *string `json:"name"`
	// 菜单标题
	Title *string `json:"title"`
	// 菜单图标
	Icon *string `json:"icon"`
	// 子菜单
	Children []MenuLocal `json:"children"`
}

type Message struct {
	Message  *beans.MsgInfo     `json:"message"`
	Template *beans.MsgTemplate `json:"template"`
}

type NewAccount struct {
	// 用户ID
	UID *string `json:"uid"`
	// 账号类型 dict: SYS002
	Type *string `json:"type"`
	// 用户名
	Username *string `json:"username"`
	// 密码
	Password *string `json:"password"`
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
	// 负责人
	Leader *string `json:"leader"`
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

type NewMenu struct {
	// 菜单标题
	Title *string `json:"title"`
	// 菜单图标
	Icon *string `json:"icon"`
	// 上级菜单
	Parent *string `json:"parent"`
	// 排序
	Weight *int `json:"weight"`
	// 状态{dict:STA001}
	Status *string `json:"status"`
}

type NewMsgConfirm struct {
	// 确认备注
	Remark *string `json:"remark"`
}

type NewPlan struct {
	// 计划名称
	Name *string `json:"name"`
	// 计划开始时间
	Pstime *int64 `json:"pstime"`
	// 计划结束时间
	Petime *int64 `json:"petime"`
	// 排序
	Weight *int `json:"weight"`
	// 状态{dict:STA001}
	Status *string `json:"status"`
}

type NewPlanStep struct {
	// 排序
	Weight *int `json:"weight"`
	// 状态{dict:STA001}
	Status *string `json:"status"`
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
	User    *NewUserInfo `json:"user"`
	Account *NewAccount  `json:"account"`
	Roles   []string     `json:"roles"`
}

type NewUserInfo struct {
	// 所属部门
	Dept *string `json:"dept"`
	// 用户类型{dict:BUS002}
	Type *string `json:"type"`
	// 用户名称
	Name *string `json:"name"`
	// 头像
	Avatar *string `json:"avatar"`
	// 工号
	Sn *string `json:"sn"`
	// 身份证
	IDCard *string `json:"id_card"`
	// 出生年月
	Birth *int `json:"birth"`
	// 性别{dict:STA002}
	Sex *string `json:"sex"`
	// 移动电话
	Mobile *string `json:"mobile"`
	// 排序
	Weight *int `json:"weight"`
	// 状态{dict:STA001}
	Status *string `json:"status"`
	// 备注
	Remark *string `json:"remark"`
	// 是否为管理员{dict:STA005}
	Admin *string `json:"admin"`
}

type PermObj struct {
	// 对象
	Object string `json:"object"`
	// 操作权限
	Mask string `json:"mask"`
}

type PlanSteps struct {
	Total *int             `json:"total"`
	Data  []beans.PlanStep `json:"data"`
}

type Plans struct {
	Total *int         `json:"total"`
	Data  []beans.Plan `json:"data"`
}

type QAccount struct {
	// 账号类型{dict:SYS002}
	Type *string `json:"type"`
	// 用户id
	UID *string `json:"uid"`
	// 用户名查询
	Username *string `json:"username"`
	// 状态{dict:STA001}
	Status *string `json:"status"`
	Index  *int    `json:"index"`
	Size   *int    `json:"size"`
	Count  *bool   `json:"count"`
}

type QCron struct {
	Keyword *string `json:"keyword"`
	// 是否启动定时任务{dict:STA003}
	State *string `json:"state"`
	// 任务名称 - 模糊查询
	Name *string `json:"name"`
	// 任务结果状态{dict:STA004}
	Result *string `json:"result"`
	// 状态{dict:STA001}
	Status *string `json:"status"`
	Index  *int    `json:"index"`
	Size   *int    `json:"size"`
	Count  *bool   `json:"count"`
}

type QCronLog struct {
	Cron    string  `json:"cron"`
	Keyword *string `json:"keyword"`
	Index   *int    `json:"index"`
	Size    *int    `json:"size"`
	Count   *bool   `json:"count"`
}

type QDept struct {
	Pid *string `json:"pid"`
	// 状态{dict:STA001}
	Status *string `json:"status"`
	// 名称模糊查询
	Name *string `json:"name"`
	// 部门编码
	Code  *string `json:"code"`
	Index *int    `json:"index"`
	Size  *int    `json:"size"`
	Count *bool   `json:"count"`
}

type QDict struct {
	// 字典类型
	Type  *string  `json:"type"`
	Types []string `json:"types"`
	Dicts []string `json:"dicts"`
	// 状态{dict:STA001}
	Status *string `json:"status"`
}

type QMenu struct {
	// 状态{dict:STA001}
	Status *string `json:"status"`
}

type QMsgHistory struct {
	// 消息ID
	Info  *string `json:"info"`
	Index *int    `json:"index"`
	Size  *int    `json:"size"`
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
	// 确认平台{dict:SYS007}
	ConfirmTarget *string `json:"confirm_target"`
	// 已读平台{dict:SYS007}
	ReadTarget *string `json:"read_target"`
	// 消息状态{ dict:SYS008}
	State *string `json:"state"`
	// 状态{dict:STA001}
	Status *string `json:"status"`
	Index  *int    `json:"index"`
	Size   *int    `json:"size"`
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
	// 确认平台{dict:SYS007}
	ConfirmTarget *string `json:"confirm_target"`
	// 已读平台{dict:SYS007}
	ReadTarget *string `json:"read_target"`
	// 消息状态{ dict:SYS008}
	State *string `json:"state"`
	// 状态{dict:STA001}
	Status *string `json:"status"`
	Index  *int    `json:"index"`
	Size   *int    `json:"size"`
}

type QPlan struct {
	// 状态{dict:STA001}
	Status *string `json:"status"`
	Index  *int    `json:"index"`
	Size   *int    `json:"size"`
	Count  *bool   `json:"count"`
}

type QPlanStep struct {
	// 状态{dict:STA001}
	Status *string `json:"status"`
	Index  *int    `json:"index"`
	Size   *int    `json:"size"`
	Count  *bool   `json:"count"`
}

type QRole struct {
	Keyword *string `json:"keyword"`
	// 状态{dict:STA001}
	Status *string `json:"status"`
	Index  *int    `json:"index"`
	Size   *int    `json:"size"`
	Count  *bool   `json:"count"`
}

type QUser struct {
	Keyword *string `json:"keyword"`
	// 部门ID
	Dept *string `json:"dept"`
	// 用户类型{dict:BUS002}
	Type *string `json:"type"`
	// 用户名称模糊查询
	Name *string `json:"name"`
	// 工号模糊查询
	Sn *string `json:"sn"`
	// 性别{dict:STA002}
	Sex *string `json:"sex"`
	// 是否为管理员{dict:STA005}
	Admin *string `json:"admin"`
	// 状态{dict:STA001}
	Status *string `json:"status"`
	Index  *int    `json:"index"`
	Size   *int    `json:"size"`
	Count  *bool   `json:"count"`
}

type Roles struct {
	Total *int         `json:"total"`
	Data  []beans.Role `json:"data"`
}

type UpdAccount struct {
	// 账号类型 dict: SYS002
	Type *string `json:"type"`
	// 用户名
	Username *string `json:"username"`
	// 密码
	Password *string `json:"password"`
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
	// 负责人
	Leader *string `json:"leader"`
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
	// 扩展
	Ext *string `json:"ext"`
	// 状态{dict:STA001}
	Status *string `json:"status"`
}

type UpdMenu struct {
	// 菜单标题
	Title *string `json:"title"`
	// 菜单图标
	Icon *string `json:"icon"`
	// 上级菜单
	Parent *string `json:"parent"`
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
	// 是否推送管理员{dict:STA005}
	ToAdmin *string `json:"to_admin"`
	// 消息提示图片
	ImgPath *string `json:"img_path"`
	// 备注
	Remark *string `json:"remark"`
	// 消息延时
	Delay *int64 `json:"delay"`
	// 消息提前提醒时间
	Alert *int64 `json:"alert"`
	// 消息模板
	Template *string `json:"template"`
}

type UpdPlan struct {
	// 计划名称
	Name *string `json:"name"`
	// 计划开始时间
	Pstime *int64 `json:"pstime"`
	// 计划结束时间
	Petime *int64 `json:"petime"`
	// 排序
	Weight *int `json:"weight"`
	// 状态{dict:STA001}
	Status *string `json:"status"`
}

type UpdPlanStep struct {
	// 排序
	Weight *int `json:"weight"`
	// 状态{dict:STA001}
	Status *string `json:"status"`
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
	User    *UpdUserInfo `json:"user"`
	Account *UpdAccount  `json:"account"`
	Roles   []string     `json:"roles"`
}

type UpdUserInfo struct {
	// 所属部门
	Dept *string `json:"dept"`
	// 用户类型{dict:BUS002}
	Type *string `json:"type"`
	// 用户名称
	Name *string `json:"name"`
	// 头像
	Avatar *string `json:"avatar"`
	// 工号
	Sn *string `json:"sn"`
	// 身份证
	IDCard *string `json:"id_card"`
	// 出生年月
	Birth *int `json:"birth"`
	// 性别{dict:STA002}
	Sex *string `json:"sex"`
	// 移动电话
	Mobile *string `json:"mobile"`
	// 排序
	Weight *int `json:"weight"`
	// 状态{dict:STA001}
	Status *string `json:"status"`
	// 备注
	Remark *string `json:"remark"`
	// 是否为管理员{dict:STA005}
	Admin *string `json:"admin"`
}

type Users struct {
	Total *int         `json:"total"`
	Data  []beans.User `json:"data"`
}
