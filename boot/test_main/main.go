package main

import (
	"errors"
	rice "github.com/GeertJohan/go.rice"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/boot"
	"github.com/zhanghup/go-app/initia"
	"github.com/zhanghup/go-app/service/ags"
	"github.com/zhanghup/go-app/service/api"
	"github.com/zhanghup/go-tools"
	"xorm.io/xorm"
)

func main() {

	box, err := rice.FindBox("conf")
	if err != nil {
		panic(err)
	}
	boot.Boot(box, "测试系统", "这是一个测试系统").Jobs("测试消息推送", "0/10 * * * * * ", func(db *xorm.Engine) error {
		tpl := beans.MsgTemplate{}
		ok, err := db.Where("code = ?", "system").Get(&tpl)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("消息模板不存在")
		}

		return ags.MessageSend(tpl, "root", "root", "user", "root", "今天天气好晴朗，处处好风光", map[string]string{
			"name": tools.StrOfRand(8),
			"time": tools.Time.HMS(),
		})
	}).Router(func(g *gin.Engine, db *xorm.Engine) {
		g.GET("/", func(ctx *gin.Context) {
			ctx.Redirect(302, "zpw")
		})
		ags.GinAgs(g.Group(""), g.Group(""))
		ags.GinStatic(box, g.Group(""), "zpw")
		api.Gin(g.Group(""), db)
	}).
		Cmd(func(db *xorm.Engine) []cli.Command {
			return []cli.Command{
				{
					Name:        "test",
					Description: "初始化测试数据",
					Action: func(c *cli.Context) error {
						initia.InitTest(db)
						return nil
					},
				},
			}
		}).
		StartRouter()

}
