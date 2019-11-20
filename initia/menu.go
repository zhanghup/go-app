package initia

import (
	"github.com/go-xorm/xorm"
	"github.com/zhanghup/go-tools"
)

type MenuObj struct {
	Id        string  `json:"id" xorm:"Varchar(32) pk"`
	Created   int64   `json:"created" xorm:"created Int(14)"`
	Updated   int64   `json:"updated" xorm:"updated  Int(14)"`
	Weight    int     `json:"weight" xorm:"weight  Int(9)"`
	Status    int     `json:"status" xorm:"status  Int(1)"`
	Code      string  `json:"code"`
	Title     string  `json:"title"`
	Meta      string  `json:"meta"`
	Name      string  `json:"name"`
	Path      string  `json:"path"`
	Alias     string  `json:"alias"`
	Icon      string  `json:"icon"`
	Component string  `json:"component"`
	Parent    *string `json:"parent"`
}

type Menu struct {
	Menu     MenuObj
	children []Menu
}

func upsertMenu(e *xorm.Engine, parent *string, menus []Menu) {
	for i, o := range menus {
		if o.Menu.Code == "" {
			panic("menu 的code字段不能为空")
		}
		o.Menu.Id = *tools.ObjectString()
		o.Menu.Status = 1
		o.Menu.Weight = i
		o.Menu.Parent = parent
		if o.children != nil && len(o.children) > 0 {
			upsertMenu(e, &o.Menu.Id, o.children)
		}
	}
}

func initMenu(e *xorm.Engine) {
	_, err := e.SF(`delete from {{ table "menu" }}`).Execute()
	if err != nil {
		panic(err)
	}
	upsertMenu(e, nil, []Menu{
		{
			Menu: MenuObj{Path: "home", Code: "home", Title: "主页", Meta: "", Name: "Home", Alias: "/", Icon: "iconfont icon-home", Component: "",},
		},
		{
			Menu: MenuObj{Path: "system", Code: "system", Title: "系统设置", Meta: "", Name: "System", Icon: "iconfont icon-system", Component: "",},
			children: []Menu{
				{Menu: MenuObj{Path: "system", Code: "system", Title: "用户管理", Meta: "", Name: "System", Icon: "iconfont icon-system", Component: ""}},
				{Menu: MenuObj{Path: "system", Code: "system", Title: "角色管理", Meta: "", Name: "System", Icon: "iconfont icon-system", Component: ""}},
				{Menu: MenuObj{Path: "system", Code: "system", Title: "菜单管理", Meta: "", Name: "System", Icon: "iconfont icon-system", Component: ""}},
			},
		},
		{
			Menu: MenuObj{Path: "system", Code: "system", Title: "日志审计", Meta: "", Name: "System", Icon: "iconfont icon-system", Component: "",},
			children: []Menu{
				{Menu: MenuObj{Path: "system", Code: "system", Title: "在线人员", Meta: "", Name: "System", Icon: "iconfont icon-system", Component: ""}},
				{Menu: MenuObj{Path: "system", Code: "system", Title: "登录日志", Meta: "", Name: "System", Icon: "iconfont icon-system", Component: ""}},
				{Menu: MenuObj{Path: "system", Code: "system", Title: "操作日志", Meta: "", Name: "System", Icon: "iconfont icon-system", Component: ""}},
			},
		},
		{
			Menu: MenuObj{Path: "system", Code: "system", Title: "平台监控", Meta: "", Name: "System", Icon: "iconfont icon-system", Component: "",},
			children: []Menu{
				{Menu: MenuObj{Path: "system", Code: "system", Title: "数据库监控", Meta: "", Name: "System", Icon: "iconfont icon-system", Component: ""}},
				{Menu: MenuObj{Path: "system", Code: "system", Title: "服务监控", Meta: "", Name: "System", Icon: "iconfont icon-system", Component: ""}},
			},
		},
	})
}
