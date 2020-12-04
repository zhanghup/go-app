package beans

// 组织
type Dept struct {
	Bean   `xorm:"extends"`
	Type   *string `json:"type"`   // 组织类型
	Code   *string `json:"code"`   // 组织代码
	Name   *string `json:"name"`   // 组织名称
	Avatar *string `json:"avatar"` // 组织头像
	Pid    *string `json:"pid"`    // 父级组织ID
	Remark *string `json:"remark"` // 备注
}

// 用户
type User struct {
	Bean   `xorm:"extends"`
	Dept   *string `json:"dept"`    // 所属部门
	Type   *string `json:"type"`    // 用户类型
	Name   *string `json:"name"`    // 用户名称
	Avatar *string `json:"avatar"`  // 头像
	IdCard *string `json:"id_card"` // 身份证ID
	Birth  *int64  `json:"birth"`   // 出生日期
	Sex    *string `json:"sex"`     // 人物性别
	Mobile *string `json:"mobile"`  // 联系电话
	Remark *string `json:"remark"`  // 备注
}

// 账户
type Account struct {
	Bean     `xorm:"extends"`
	Uid      *string `json:"uid"`           // 用户ID
	Type     *string `json:"type"`          // 字典
	Username *string `json:"username"`      // 用户名
	Password *string `json:"password"`      // 密码
	Salt     *string `json:"-" xorm:"salt"` // 加盐
	Admin    *int    `json:"admin"`         // 是否为管理员账户
	Default  *int    `json:"default"`       // 是否为默认账户
}

// 授权
type Token struct {
	Bean   `xorm:"extends"`
	Uid    *string `json:"uid"`    // 用户ID
	Aid    *string `json:"aid"`    // accountID
	Ops    *int64  `json:"ops"`    // 接口调用次数
	Expire *int64  `json:"expire"` // 到期时间
	Agent  *string `json:"agent"`  // User-Agent
}

// 数据字典
type Dict struct {
	Bean   `xorm:"extends"`
	Code   *string `json:"code" xorm:"unique"` // 字典编码
	Name   *string `json:"name"`               // 字典名称
	Type   *string `json:"type"`               // 字典类型
	Remark *string `json:"remark"`             // 备注
}
type DictItem struct {
	Bean     `xorm:"extends"`
	Code     *string `json:"code"`
	Name     *string `json:"name"`
	Value    *string `json:"value"`
	Ext      *string `json:"ext"`
	Disabled *int    `json:"disabled"`
}

// 权限设置
type Role struct {
	Bean `xorm:"extends"`
	Name *string `json:"name"` // 角色名称
	Desc *string `json:"desc"` // 描述
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
	Type *string `json:"type"` // 类型
	Oid  *string `json:"oid"`  // 对象ID
}

// 对象权限
type PermObject struct {
	Bean   `xorm:"extends"`
	Role   *string `json:"role"`   // 角色ID
	Object *string `json:"object"` // 表类型（user/dict等）SYS003
	Mask   *string `json:"mask"`   // 权限（C/R/U/D）等组成的字符串 SYS002
}

// 菜单
type Menu struct {
	Bean   `xorm:"extends"`
	Name   *string `json:"name"`
	Title  *string `json:"title"`
	Path   *string `json:"path"`
	Alias  *string `json:"alias"`
	Icon   *string `json:"icon"`
	Parent *string `json:"parent"`
}

// 资源
type Resource struct {
	Bean        `xorm:"extends"`
	MD5         string `json:"md5" xorm:"'md5'"`
	ContentType string `json:"content_type"`
	Name        string `json:"name"`
	Date        string `json:"date"`
	Size        int64  `json:"size"`
	FileEnd     string `json:"file_end"`
}

// 定时任务
type Cron struct {
	Bean       `xorm:"extends"`
	State      *string `json:"state"`               // 状态
	Name       *string `json:"name"`                // 任务名称
	Expression *string `json:"expression"`          // 任务表达式
	Previous   *int64  `json:"previous"`            // 上一次执行时间
	Last       *int64  `json:"last"`                // 任务持续时间（毫秒）
	Message    *string `json:"message" xorm:"text"` // 任务结果
	Result     *string `json:"result"`              // 任务结果状态
}

func sys_tables() []interface{} {
	return []interface{}{
		new(Dept),
		new(Account),
		new(Token),
		new(Dict),
		new(DictItem),
		new(Menu),
		new(Role),
		new(RoleUser),
		new(Perm),
		new(PermObject),
		new(User),
		new(Resource),
		new(Cron),
	}
}
