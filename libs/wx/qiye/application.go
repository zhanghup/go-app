package qiye

import (
	"fmt"
	"github.com/zhanghup/go-app/cfg"
)

type application_api struct {
	corpId  string
	secret  string
	agentid string
}

func newApplicationApi(agentid string) *application_api {

	if !cfg.WxQyApp(agentid).Enable {
		panic(fmt.Sprintf("config.ini - wxqy-app[%s].enable 未启用", agentid))
	}
	app := application_api{corpId: cfg.WxQy().Corpid, secret: cfg.WxQyApp(agentid).Secret, agentid: agentid}
	return &app
}

func (this *application_api) access_token() (string, error) {
	return newAccessToken().getToken(this.corpId, this.secret)
}
