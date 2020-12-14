// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package source

import (
	"github.com/zhanghup/go-app/beans"
)

type Accounts struct {
	Total *int            `json:"total" xorm:"total"`
	Data  []beans.Account `json:"data" xorm:"data"`
}

type CronLogs struct {
	Total *int            `json:"total" xorm:"total"`
	Data  []beans.CronLog `json:"data" xorm:"data"`
}

type Crons struct {
	Total *int         `json:"total" xorm:"total"`
	Data  []beans.Cron `json:"data" xorm:"data"`
}

type Depts struct {
	Total *int         `json:"total" xorm:"total"`
	Data  []beans.Dept `json:"data" xorm:"data"`
}

type IPermObj struct {
	// 对象
	Object string `json:"object" xorm:"object"`
	// 操作权限
	Mask string `json:"mask" xorm:"mask"`
}

type Message struct {
	Message  *beans.MsgInfo     `json:"message" xorm:"message"`
	Template *beans.MsgTemplate `json:"template" xorm:"template"`
}

type NewAccount struct {
	// 用户ID
	UID *string `json:"uid" xorm:"uid"`
	// 账号类型 dict: SYS002
	Type string `json:"type" xorm:"type"`
	// 用户名
	Username *string `json:"username" xorm:"username"`
	// 密码
	Password *string `json:"password" xorm:"password"`
	// 是否为默认账户，默认账户可以在用户列表中可见并且维护
	Default *int `json:"default" xorm:"default"`
	// 排序
	Weight *int `json:"weight" xorm:"weight"`
	// 状态{dict:STA001}
	Status *string `json:"status" xorm:"status"`
}

type NewDept struct {
	// 组织类型{dict:BUS001}
	Type *string `json:"type" xorm:"type"`
	// 组织代码
	Code *string `json:"code" xorm:"code"`
	// 组织名称
	Name *string `json:"name" xorm:"name"`
	// 组织头像
	Avatar *string `json:"avatar" xorm:"avatar"`
	// 父级组织ID
	Pid *string `json:"pid" xorm:"pid"`
	// 备注
	Remark *string `json:"remark" xorm:"remark"`
	// 负责人
	Leader *string `json:"leader" xorm:"leader"`
	// 排序
	Weight *int `json:"weight" xorm:"weight"`
	// 状态{dict:STA001}
	Status *string `json:"status" xorm:"status"`
}

type NewDict struct {
	// 字典类型{dict:SYS001}
	Type *string `json:"type" xorm:"type"`
	// 字典编码
	Code *string `json:"code" xorm:"code"`
	// 字典名称
	Name *string `json:"name" xorm:"name"`
	// 备注
	Remark *string `json:"remark" xorm:"remark"`
	// 排序
	Weight *int `json:"weight" xorm:"weight"`
	// 状态{dict:STA001}
	Status *string `json:"status" xorm:"status"`
}

type NewDictItem struct {
	// 字典id
	Code *string `json:"code" xorm:"code"`
	// 名称
	Name *string `json:"name" xorm:"name"`
	// 值
	Value *string `json:"value" xorm:"value"`
	// 扩展
	Ext *string `json:"ext" xorm:"ext"`
	// 排序
	Weight *int `json:"weight" xorm:"weight"`
	// 状态{dict:STA001}
	Status *string `json:"status" xorm:"status"`
}

type NewMsgConfirm struct {
	// 确认备注
	Remark *string `json:"remark" xorm:"remark"`
}

type NewPlan struct {
	// 计划名称
	Name *string `json:"name" xorm:"name"`
	// 计划开始时间
	Pstime *int64 `json:"pstime" xorm:"pstime"`
	// 计划结束时间
	Petime *int64 `json:"petime" xorm:"petime"`
	// 排序
	Weight *int `json:"weight" xorm:"weight"`
	// 状态{dict:STA001}
	Status *string `json:"status" xorm:"status"`
}

type NewPlanStep struct {
	// 排序
	Weight *int `json:"weight" xorm:"weight"`
	// 状态{dict:STA001}
	Status *string `json:"status" xorm:"status"`
}

type NewRole struct {
	// 角色描述
	Name *string `json:"name" xorm:"name"`
	// 角色名称
	Desc *string `json:"desc" xorm:"desc"`
	// 排序
	Weight *int `json:"weight" xorm:"weight"`
	// 状态{dict:STA001}
	Status *string `json:"status" xorm:"status"`
}

type NewUser struct {
	User    *NewUserInfo `json:"user" xorm:"user"`
	Account *NewAccount  `json:"account" xorm:"account"`
	Roles   []string     `json:"roles" xorm:"roles"`
}

type NewUserInfo struct {
	// 所属部门
	Dept *string `json:"dept" xorm:"dept"`
	// 用户类型{dict:BUS002}
	Type *string `json:"type" xorm:"type"`
	// 用户名称
	Name *string `json:"name" xorm:"name"`
	// 头像
	Avatar *string `json:"avatar" xorm:"avatar"`
	// 身份证
	IDCard *string `json:"id_card" xorm:"id_card"`
	// 出生年月
	Birth *int `json:"birth" xorm:"birth"`
	// 性别{dict:STA002}
	Sex *string `json:"sex" xorm:"sex"`
	// 移动电话
	Mobile *string `json:"mobile" xorm:"mobile"`
	// 排序
	Weight *int `json:"weight" xorm:"weight"`
	// 状态{dict:STA001}
	Status *string `json:"status" xorm:"status"`
	// 备注
	Remark *string `json:"remark" xorm:"remark"`
	// 是否为管理员{dict:STA005}
	Admin *string `json:"admin" xorm:"admin"`
}

type PermObj struct {
	// 对象
	Object string `json:"object" xorm:"object"`
	// 操作权限
	Mask string `json:"mask" xorm:"mask"`
}

type PlanSteps struct {
	Total *int             `json:"total" xorm:"total"`
	Data  []beans.PlanStep `json:"data" xorm:"data"`
}

type Plans struct {
	Total *int         `json:"total" xorm:"total"`
	Data  []beans.Plan `json:"data" xorm:"data"`
}

type QAccount struct {
	// 账号类型{dict:SYS002}
	Type *string `json:"type" xorm:"type"`
	// 用户id
	UID *string `json:"uid" xorm:"uid"`
	// 用户名查询
	Username *string `json:"username" xorm:"username"`
	// 状态{dict:STA001}
	Status *string `json:"status" xorm:"status"`
	Index  *int    `json:"index" xorm:"index"`
	Size   *int    `json:"size" xorm:"size"`
	Count  *bool   `json:"count" xorm:"count"`
}

type QCron struct {
	Keyword *string `json:"keyword" xorm:"keyword"`
	// 是否启动定时任务{dict:STA003}
	State *string `json:"state" xorm:"state"`
	// 任务名称 - 模糊查询
	Name *string `json:"name" xorm:"name"`
	// 任务结果状态{dict:STA004}
	Result *string `json:"result" xorm:"result"`
	// 状态{dict:STA001}
	Status *string `json:"status" xorm:"status"`
	Index  *int    `json:"index" xorm:"index"`
	Size   *int    `json:"size" xorm:"size"`
	Count  *bool   `json:"count" xorm:"count"`
}

type QCronLog struct {
	Cron    string  `json:"cron" xorm:"cron"`
	Keyword *string `json:"keyword" xorm:"keyword"`
	Index   *int    `json:"index" xorm:"index"`
	Size    *int    `json:"size" xorm:"size"`
	Count   *bool   `json:"count" xorm:"count"`
}

type QDept struct {
	Pid *string `json:"pid" xorm:"pid"`
	// 状态{dict:STA001}
	Status *string `json:"status" xorm:"status"`
	// 名称模糊查询
	Name  *string `json:"name" xorm:"name"`
	Index *int    `json:"index" xorm:"index"`
	Size  *int    `json:"size" xorm:"size"`
	Count *bool   `json:"count" xorm:"count"`
}

type QDict struct {
	// 字典类型
	Type *string `json:"type" xorm:"type"`
	// 状态{dict:STA001}
	Status *string `json:"status" xorm:"status"`
}

type QMsgHistory struct {
	// 消息ID
	Info  *string `json:"info" xorm:"info"`
	Index *int    `json:"index" xorm:"index"`
	Size  *int    `json:"size" xorm:"size"`
}

type QMsgInfo struct {
	// 接收者
	Receiver *string `json:"receiver" xorm:"receiver"`
	// 消息类型{dict:SYS005}
	Type *string `json:"type" xorm:"type"`
	// 消息级别{dict: SYS006}
	Level *string `json:"level" xorm:"level"`
	// 消息接收平台{dict:SYS007}
	Target *string `json:"target" xorm:"target"`
	// 确认平台{dict:SYS007}
	ConfirmTarget *string `json:"confirm_target" xorm:"confirm_target"`
	// 已读平台{dict:SYS007}
	ReadTarget *string `json:"read_target" xorm:"read_target"`
	// 消息状态{ dict:SYS008}
	State *string `json:"state" xorm:"state"`
	// 状态{dict:STA001}
	Status *string `json:"status" xorm:"status"`
	Index  *int    `json:"index" xorm:"index"`
	Size   *int    `json:"size" xorm:"size"`
}

type QMsgTemplate struct {
	// 名称模糊查询
	Name *string `json:"name" xorm:"name"`
	// 编码模糊查询
	Code *string `json:"code" xorm:"code"`
}

type QMyMsgInfo struct {
	// 消息类型{dict:SYS005}
	Type *string `json:"type" xorm:"type"`
	// 消息级别{dict: SYS006}
	Level *string `json:"level" xorm:"level"`
	// 弹出消息是否必须确认{dict:STA005}
	MustConfirm *string `json:"must_confirm" xorm:"must_confirm"`
	// 确认平台{dict:SYS007}
	ConfirmTarget *string `json:"confirm_target" xorm:"confirm_target"`
	// 已读平台{dict:SYS007}
	ReadTarget *string `json:"read_target" xorm:"read_target"`
	// 消息状态{ dict:SYS008}
	State *string `json:"state" xorm:"state"`
	// 状态{dict:STA001}
	Status *string `json:"status" xorm:"status"`
	Index  *int    `json:"index" xorm:"index"`
	Size   *int    `json:"size" xorm:"size"`
}

type QPlan struct {
	// 状态{dict:STA001}
	Status *string `json:"status" xorm:"status"`
	Index  *int    `json:"index" xorm:"index"`
	Size   *int    `json:"size" xorm:"size"`
	Count  *bool   `json:"count" xorm:"count"`
}

type QPlanStep struct {
	// 状态{dict:STA001}
	Status *string `json:"status" xorm:"status"`
	Index  *int    `json:"index" xorm:"index"`
	Size   *int    `json:"size" xorm:"size"`
	Count  *bool   `json:"count" xorm:"count"`
}

type QRole struct {
	Keyword *string `json:"keyword" xorm:"keyword"`
	// 状态{dict:STA001}
	Status *string `json:"status" xorm:"status"`
	Index  *int    `json:"index" xorm:"index"`
	Size   *int    `json:"size" xorm:"size"`
	Count  *bool   `json:"count" xorm:"count"`
}

type QUser struct {
	Keyword *string `json:"keyword" xorm:"keyword"`
	// 状态{dict:STA001}
	Status *string `json:"status" xorm:"status"`
	Index  *int    `json:"index" xorm:"index"`
	Size   *int    `json:"size" xorm:"size"`
	Count  *bool   `json:"count" xorm:"count"`
}

type Roles struct {
	Total *int         `json:"total" xorm:"total"`
	Data  []beans.Role `json:"data" xorm:"data"`
}

type UpdAccount struct {
	// 账号类型 dict: SYS002
	Type string `json:"type" xorm:"type"`
	// 用户名
	Username *string `json:"username" xorm:"username"`
	// 密码
	Password *string `json:"password" xorm:"password"`
	// 是否为默认账户，默认账户可以在用户列表中可见并且维护
	Default *int `json:"default" xorm:"default"`
	// 排序
	Weight *int `json:"weight" xorm:"weight"`
	// 状态{dict:STA001}
	Status *string `json:"status" xorm:"status"`
}

type UpdDept struct {
	// 组织类型{dict:BUS001}
	Type *string `json:"type" xorm:"type"`
	// 组织代码
	Code *string `json:"code" xorm:"code"`
	// 组织名称
	Name *string `json:"name" xorm:"name"`
	// 组织头像
	Avatar *string `json:"avatar" xorm:"avatar"`
	// 父级组织ID
	Pid *string `json:"pid" xorm:"pid"`
	// 备注
	Remark *string `json:"remark" xorm:"remark"`
	// 负责人
	Leader *string `json:"leader" xorm:"leader"`
	// 排序
	Weight *int `json:"weight" xorm:"weight"`
	// 状态{dict:STA001}
	Status *string `json:"status" xorm:"status"`
}

type UpdDict struct {
	// 字典类型{dict:SYS001}
	Type *string `json:"type" xorm:"type"`
	// 字典名称
	Name *string `json:"name" xorm:"name"`
	// 备注
	Remark *string `json:"remark" xorm:"remark"`
	// 排序
	Weight *int `json:"weight" xorm:"weight"`
	// 状态{dict:STA001}
	Status *string `json:"status" xorm:"status"`
}

type UpdDictItem struct {
	// 名称
	Name *string `json:"name" xorm:"name"`
	// 值
	Value *string `json:"value" xorm:"value"`
	// 扩展
	Ext *string `json:"ext" xorm:"ext"`
	// 排序
	Weight *int `json:"weight" xorm:"weight"`
	// 状态{dict:STA001}
	Status *string `json:"status" xorm:"status"`
}

type UpdMsgTemplate struct {
	// 模板名称
	Name *string `json:"name" xorm:"name"`
	// 消息分类{dict:SYS005}
	Type *string `json:"type" xorm:"type"`
	// 消息等级{dict:SYS006}
	Level *string `json:"level" xorm:"level"`
	// 消息推送平台{dict:SYS007}
	Target *string `json:"target" xorm:"target"`
	// 消息超时时间（秒）
	Expire *int64 `json:"expire" xorm:"expire"`
	// 消息是否必须确认{dict:STA005}
	MustConfirm *string `json:"must_confirm" xorm:"must_confirm"`
	// 消息提示图片
	ImgPath *string `json:"img_path" xorm:"img_path"`
	// 备注
	Remark *string `json:"remark" xorm:"remark"`
}

type UpdPlan struct {
	// 计划名称
	Name *string `json:"name" xorm:"name"`
	// 计划开始时间
	Pstime *int64 `json:"pstime" xorm:"pstime"`
	// 计划结束时间
	Petime *int64 `json:"petime" xorm:"petime"`
	// 排序
	Weight *int `json:"weight" xorm:"weight"`
	// 状态{dict:STA001}
	Status *string `json:"status" xorm:"status"`
}

type UpdPlanStep struct {
	// 排序
	Weight *int `json:"weight" xorm:"weight"`
	// 状态{dict:STA001}
	Status *string `json:"status" xorm:"status"`
}

type UpdRole struct {
	// 角色描述
	Name *string `json:"name" xorm:"name"`
	// 角色名称
	Desc *string `json:"desc" xorm:"desc"`
	// 排序
	Weight *int `json:"weight" xorm:"weight"`
	// 状态{dict:STA001}
	Status *string `json:"status" xorm:"status"`
}

type UpdUser struct {
	User    *UpdUserInfo `json:"user" xorm:"user"`
	Account *UpdAccount  `json:"account" xorm:"account"`
	Roles   []string     `json:"roles" xorm:"roles"`
}

type UpdUserInfo struct {
	// 所属部门
	Dept *string `json:"dept" xorm:"dept"`
	// 用户类型{dict:BUS002}
	Type *string `json:"type" xorm:"type"`
	// 用户名称
	Name *string `json:"name" xorm:"name"`
	// 头像
	Avatar *string `json:"avatar" xorm:"avatar"`
	// 身份证
	IDCard *string `json:"id_card" xorm:"id_card"`
	// 出生年月
	Birth *int `json:"birth" xorm:"birth"`
	// 性别{dict:STA002}
	Sex *string `json:"sex" xorm:"sex"`
	// 移动电话
	Mobile *string `json:"mobile" xorm:"mobile"`
	// 排序
	Weight *int `json:"weight" xorm:"weight"`
	// 状态{dict:STA001}
	Status *string `json:"status" xorm:"status"`
	// 备注
	Remark *string `json:"remark" xorm:"remark"`
	// 是否为管理员{dict:STA005}
	Admin *string `json:"admin" xorm:"admin"`
}

type Users struct {
	Total *int         `json:"total" xorm:"total"`
	Data  []beans.User `json:"data" xorm:"data"`
}
