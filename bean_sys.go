package app

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

	Code   *string `json:"code" xorm:"unique"`
	Name   *string `json:"name"`
	Remark *string `json:"remark"`
}
type DictItem struct {
	Bean `xorm:"extends"`

	Code      *string `json:"code"`
	Name      *string `json:"name"`
	Value     *string `json:"value"`
	Extension *string `json:"extension"`
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

// 用户
type User struct {
	Bean `xorm:"extends"`

	Type     *string `json:"type"` // D0001 用户类型
	Account  *string `json:"account" xorm:"unique"`
	Password *string `json:"password"`
	Slat     *string `json:"-" xorm:"slat"`
	Name     *string `json:"name"`
	Avatar   *string `json:"avatar"`
	ICard    *string `json:"i_card"`
	Birth    *int64  `json:"birth"`
	Sex      *int    `json:"sex"`    // 0：未知，1：男，2：女
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

type Cron struct {
	Bean `xorm:"extends"`

	Code       *string `json:"code" xorm:"index"`   // 编码
	Name       *string `json:"name"`                // 任务名称
	Expression *string `json:"expression"`          // 任务表达式
	Previous   *int64  `json:"previous"`            // 上一次执行时间
	Last       *int64  `json:"last"`                // 任务持续时间（秒）
	Message    *string `json:"message" xorm:"text"` // 任务结果
}

func sys_tables() []interface{} {
	return []interface{}{new(Dict), new(DictItem), new(Menu), new(Role), new(RoleUser), new(Perm), new(User), new(UserToken), new(Resource), new(Cron)}
}