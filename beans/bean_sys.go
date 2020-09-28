package beans

// 授权
type UserToken struct {
	Bean   `xorm:"extends"`
	Uid    *string `json:"uid"`
	Ops    *int64  `json:"ops"`    // 接口调用次数
	Type   *string `json:"type"`   // 授权类型 [pc]
	Expire *int64  `json:"expire"` // 到期时间
	Agent  *string `json:"agent"`  // User-Agent
}

// 数据字典
type Dict struct {
	Bean `xorm:"extends"`

	Code   *string `json:"code" xorm:"unique"`
	Name   *string `json:"name"`
	Type   *string `json:"type"`
	Remark *string `json:"remark"`
}
type DictItem struct {
	Bean `xorm:"extends"`

	Code      *string `json:"code"`
	Name      *string `json:"name"`
	Value     *string `json:"value"`
	Disable   *int    `json:"disable"`
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
	Uid  *string `json:"uid"`
}

// 数据权限
type Perm struct {
	Bean `xorm:"extends"`

	Role *string `json:"role"` // 角色ID
	Type *string `json:"type"` // 类型（menu等）
	Oid  *string `json:"oid"`  // 对象ID
}

// 对象权限
type PermObject struct {
	Bean `xorm:"extends"`

	Role   *string `json:"role"`   // 角色ID
	Object *string `json:"object"` // 表类型（user/dict等）SYS003
	Mask   *string `json:"mask"`   // 权限（C/R/U/D）等组成的字符串 SYS002
}

// 用户
type User struct {
	Bean `xorm:"extends"`

	Type     *string `json:"type"` // 字典SYS001 用户类型
	Account  *string `json:"account" xorm:"unique"`
	Password *string `json:"password"`
	Salt     *string `json:"-" xorm:"salt"`
	Name     *string `json:"name"`
	Avatar   *string `json:"avatar"`
	ICard    *string `json:"i_card"`
	Birth    *int64  `json:"birth"`
	Sex      *int    `json:"sex"`    // 字典STA002 人物性别
	Mobile   *string `json:"mobile"` // 联系电话
	Admin    *int    `json:"admin"`
	Remark   *string `json:"remark"`
}

// 菜单
type Menu struct {
	Bean `xorm:"extends"`

	Code      *string `json:"code"`
	Title     *string `json:"title"`
	Meta      *string `json:"meta"`
	Name      *string `json:"name"`
	Path      *string `json:"path"`
	Alias     *string `json:"alias"`
	Icon      *string `json:"icon"`
	Component *string `json:"component"`

	Parent *string `json:"parent"`
}

// 资源
type Resource struct {
	Bean `xorm:"extends"`

	MD5         string `json:"md5" xorm:"'md5'"`
	ContentType string `json:"content_type"`
	Name        string `json:"name"`
	Size        int64  `json:"size"`
	FileEnd     string `json:"file_end"`
}

// 定时任务
type Cron struct {
	Bean     `xorm:"extends"`
	BeanDict `xorm:"extends"`

	State      *int     `json:"state"`               // 是否启动定时任务
	Name       *string  `json:"name"`                // 任务名称
	Expression *string  `json:"expression"`          // 任务表达式
	Previous   *int64   `json:"previous"`            // 上一次执行时间
	Last       *float64 `json:"last"`                // 任务持续时间（秒）
	Message    *string  `json:"message" xorm:"text"` // 任务结果
}

func sys_tables() []interface{} {
	return []interface{}{new(UserToken), new(Dict), new(DictItem), new(Menu), new(Role), new(RoleUser), new(Perm), new(PermObject), new(User), new(Resource), new(Cron)}
}
