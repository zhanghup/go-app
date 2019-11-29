package qiye

import (
	"fmt"
	rice "github.com/giter/go.rice"
	"github.com/zhanghup/go-app/cfg"
	"github.com/zhanghup/go-tools"
	"testing"
)

var ap *application

func TestAccessToken(t *testing.T) {
	fmt.Println(ap.access_token())
}

func TestMenuCreate(t *testing.T) {
	err := ap.MenuCreate([]Button{
		{Type: "click", Name: "今日金曲", Key: "V1001_TODAY_MUSIC"},
		{Name: "菜单", SubButton: []Button{
			{Type: "view", Name: "搜索", Url: "http://www.soso.com/", Key: "V1001_TODAY_MUSIC"},
			{Type: "click", Name: "赞一下我们", Url: "V1001_GOOD", Key: "V1001_TODAY_MUSIC"},
		}},
	})
	if err != nil {
		panic(err)
	}
}

func TestMenus(t *testing.T) {
	btns, err := ap.Menus()
	if err != nil {
		panic(err)
	}
	tools.Str().JSONStringPrintln(btns)
}

func TestMenuRemove(t *testing.T) {
	err := ap.MenuRemove()
	if err != nil {
		panic(err)
	}
}

func init() {
	box, err := rice.FindBox("conf")
	if err != nil {
		panic(err)
	}
	cfg.InitConfig(box)
	ap = newApplication("1000002")
}
