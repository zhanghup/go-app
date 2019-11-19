package initia

import (
	"github.com/go-xorm/xorm"
	"github.com/zhanghup/go-app"
	"github.com/zhanghup/go-tools"
)

type Menu struct {
	app.Menu
	children []Menu
}

func UpsertMenu(e *xorm.Engine, parent *string, menus []Menu) {
	for i, o := range menus {
		if o.Menu.Code == nil {
			panic("menu 的code字段不能为空")
		}
		o.Menu.Id = tools.ObjectString()
		o.Menu.Status = tools.Ptr().Int(1)
		o.Menu.Weight = &i
		if o.children != nil && len(o.children) > 0 {
			UpsertMenu(e, o.Menu.Id, o.children)
		}
	}
}

func InitMenu(e *xorm.Engine) {
	UpsertMenu(e, nil, []Menu{
		{
			Menu: app.Menu{
				Path:      tools.Ptr().String("home"),
				Code:      tools.Ptr().String("home"),
				Title:     tools.Ptr().String("主页"),
				Meta:      tools.Ptr().String(""),
				Name:      tools.Ptr().String("Home"),
				Alias:     tools.Ptr().String("/"),
				Icon:      tools.Ptr().String("iconfont icon-home"),
				Component: tools.Ptr().String("iconfont icon-home"),
			},
			children: []Menu{},
		},
	})
}
