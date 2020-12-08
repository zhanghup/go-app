package main

import (
	"github.com/gin-gonic/gin"
	rice "github.com/giter/go.rice"
	"github.com/zhanghup/go-app/boot"
	"github.com/zhanghup/go-app/service/ags"
	"github.com/zhanghup/go-app/service/api"
	"xorm.io/xorm"
)

func main() {
	box, err := rice.FindBox("conf")
	if err != nil {
		panic(err)
	}
	_ = boot.Boot(box,"测试系统").
		//SyncTables().
		//Init().
		//InitTestData().
		XormInited().
		//JobsInit().
		//JobsInitMessages().
		//Jobs("测试消息推送", "0/5 * * * * * ", func(db *xorm.Engine) error {
		//	tpl := beans.MsgTemplate{}
		//	ok, err := db.Where("code = ?", "system").Get(&tpl)
		//	if err != nil {
		//		return err
		//	}
		//	if !ok {
		//		return errors.New("消息模板不存在")
		//	}
		//
		//	return ags.NewMessage(db).NewMessage(tpl, "root", "root", "user", "root", "天气不错 - "+tools.Ti.HMS(), "今天天气好晴朗，处处好风光")
		//}).
		//Jobs("河东域名同步", "0 * * * * * ", func() error {
		//	_, items, ok := ca.DictCache.Get("AUT001")
		//	if !ok {
		//		return errors.New("无同步内容[0]")
		//	}
		//	if len(items) == 0 {
		//		return errors.New("无同步内容[1]")
		//	}
		//
		//	return nil
		//}).
		Router(func(g *gin.Engine, db *xorm.Engine) {
			ags.Gin(g.Group(""), g.Group(""), db)
			api.Gin(g.Group(""), db)
		}).
		StartRouter()
}
