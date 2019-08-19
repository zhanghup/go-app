package app

import "github.com/go-xorm/xorm"

type PageResult struct {
	Total int64       `json:"total"`
	Datas interface{} `json:"datas"`
}
type PageParam struct {
	Index int  `json:"index"`
	Size  int  `json:"size"`
	Count bool `json:"count"`
}

type Bean struct {
	Id      *string `json:"id" xorm:"Varchar(32) pk"`
	Created *int64  `json:"created" xorm:"created Int(14)"`
	Updated *int64  `json:"updated" xorm:"updated  Int(14)"`
	Weight  *int    `json:"weight" xorm:"weight  Int(9)"`
	Status  *int    `json:"status" xorm:"status  Int(1)"`
}

// 数据字典
type Dict struct {
	Bean `xorm:"extends"`

	Name   *string `json:"name"`
	Remark *string `json:"remark"`

	Values []DictItem `json:"values" xorm:"-"`
}
type DictItem struct {
	Bean `xorm:"extends"`

	Code      *string `json:"code"`
	Name      *string `json:"name"`
	Value     *string `json:"value"`
	Extension *string `json:"extension"`
}

// 菜单显示
type Menu struct {
	Bean `xorm:"extends"`

	Title  *string `json:"title"`
	Name   *string `json:"name"`
	NameEn *string `json:"name_en"`
	Index  *string `json:"index"`
	Icon   *string `json:"icon"`
	Parent *string `json:"parent"`
}

// 权限设置
type Role struct {
	Bean `xorm:"extends"`

	Name *string `json:"name"`
	Desc *string `json:"desc"`
}
type RoleUser struct {
	Bean `xorm:"extends"`

	Role *string `json:"role"`
	User *string `json:"user"`
}
type Perm struct {
	Bean `xorm:"extends"`

	Type *string `json:"type"` // 类型（menu等）
	Role *string `json:"role"` // 角色ID
	Oid  *string `json:"oid"`  // 对象ID
	Mask *string `json:"mask"` // 权限
}

// 给用户
type User struct {
	Bean `xorm:"extends"`

	Type     *string `json:"type"`
	Account  *string `json:"account" xorm:"unique"`
	Password *string `json:"password"`
	Slat     *string `json:"slot"`
	Name     *string `json:"name"`
	Avatar   *string `json:"avatar"`
	IdCard   *string `json:"id_card"`
	Birth    *int64  `json:"birth"`
	Sex      *string `json:"sex"`    // 0：未知，1：男，2：女
	Mobile   *string `json:"mobile"` // 联系电话
	Admin    *int    `json:"admin"`
}

// 授权
type UserToken struct {
	Bean   `xorm:"extends"`
	User   *string `json:"user"`
	Ops    *int64  `json:"ops"`    // 接口调用次数
	Type   *string `json:"type"`   // 授权类型 [pc,wx:微信小程序，we:微信公众号]
	Expire *int64  `json:"expire"` // 到期时间
	Agent  *string `json:"agent"`  // User-Agent
}

type Resource struct {
	Bean `xorm:"extends"`

	MD5         string `json:"md5" xorm:"'md5'"`
	ContentType string `json:"content_type"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Size        int64  `json:"size"`
	Datas       []byte `json:"datas" xorm:"longblob"`
}

func Sync(e *xorm.Engine) {
	err := e.Sync2(new(Dict), new(DictItem), new(Menu), new(Role), new(RoleUser), new(Perm), new(User), new(UserToken), new(Resource))
	if err != nil {
		panic(err)
	}
}
