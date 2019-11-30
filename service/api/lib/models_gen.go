// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package lib

import (
	"github.com/zhanghup/go-app/beans"
)

type Dicts struct {
	Total *int         `json:"total"`
	Data  []beans.Dict `json:"data"`
}

type IPermObj struct {
	// 对象
	Object string `json:"object"`
	// 操作权限
	Mask string `json:"mask"`
}

type NewDict struct {
	// 字典编码
	Code *string `json:"code"`
	// 字典名称
	Name *string `json:"name"`
	// 备注
	Remark *string `json:"remark"`
	// 排序
	Weight *int `json:"weight"`
	// 状态[dict:STA0001]
	Status *int `json:"status"`
}

type NewDictItem struct {
	// 字典id
	Code *string `json:"code"`
	// 名称
	Name *string `json:"name"`
	// 值
	Value *string `json:"value"`
	// 扩展
	Extension *string `json:"extension"`
	// 排序
	Weight *int `json:"weight"`
	// 状态[dict:STA0001]
	Status *int `json:"status"`
}

type NewRole struct {
	// 角色描述
	Name *string `json:"name"`
	// 角色名称
	Desc *string `json:"desc"`
	// 排序
	Weight *int `json:"weight"`
	// 状态[dict:STA0001]
	Status *int `json:"status"`
}

type NewUser struct {
	// 用户类型
	Type string `json:"type"`
	// 账户
	Account string `json:"account"`
	// 密码
	Password string `json:"password"`
	// 用户名称
	Name string `json:"name"`
	// 头像
	Avatar *string `json:"avatar"`
	// 身份证
	ICard *string `json:"i_card"`
	// 出生年月
	Birth *int `json:"birth"`
	// 性别[dict:STA0002]
	Sex *int `json:"sex"`
	// 移动电话
	Mobile *string `json:"mobile"`
	// 是否为管理员[0: 否,1: 是]
	Admin *int `json:"admin"`
	// 排序
	Weight *int `json:"weight"`
	// 状态[dict:STA0001]
	Status *int `json:"status"`
}

type PermObj struct {
	// 对象
	Object string `json:"object"`
	// 操作权限
	Mask string `json:"mask"`
}

type QDict struct {
	Index *int  `json:"index"`
	Size  *int  `json:"size"`
	Count *bool `json:"count"`
}

type QRole struct {
	Index *int  `json:"index"`
	Size  *int  `json:"size"`
	Count *bool `json:"count"`
}

type QUser struct {
	Index *int  `json:"index"`
	Size  *int  `json:"size"`
	Count *bool `json:"count"`
}

type Roles struct {
	Total *int         `json:"total"`
	Data  []beans.Role `json:"data"`
}

type UpdDict struct {
	// 字典名称
	Name *string `json:"name"`
	// 备注
	Remark *string `json:"remark"`
	// 排序
	Weight *int `json:"weight"`
	// 状态[dict:STA0001]
	Status *int `json:"status"`
}

type UpdDictItem struct {
	// 名称
	Name *string `json:"name"`
	// 值
	Value *string `json:"value"`
	// 扩展
	Extension *string `json:"extension"`
	// 排序
	Weight *int `json:"weight"`
	// 状态[dict:STA0001]
	Status *int `json:"status"`
}

type UpdRole struct {
	// 角色描述
	Name *string `json:"name"`
	// 角色名称
	Desc *string `json:"desc"`
	// 排序
	Weight *int `json:"weight"`
	// 状态[dict:STA0001]
	Status *int `json:"status"`
}

type UpdUser struct {
	// 用户类型
	Type string `json:"type"`
	// 账户
	Account string `json:"account"`
	// 密码
	Password string `json:"password"`
	// 用户名称
	Name string `json:"name"`
	// 头像
	Avatar *string `json:"avatar"`
	// 身份证
	ICard *string `json:"i_card"`
	// 出生年月
	Birth *int `json:"birth"`
	// 性别[dict:STA0002]
	Sex *int `json:"sex"`
	// 移动电话
	Mobile *string `json:"mobile"`
	// 是否为管理员[0: 否,1: 是]
	Admin *int `json:"admin"`
	// 排序
	Weight *int `json:"weight"`
	// 状态[dict:STA0001]
	Status *int `json:"status"`
}

type Users struct {
	Total *int         `json:"total"`
	Data  []beans.User `json:"data"`
}
