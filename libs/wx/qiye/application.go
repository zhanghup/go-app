package qiye

import (
	"github.com/zhanghup/go-app/cfg"
)

type application_api struct {
	corpId  string
	secret  string
	agentid string
}

func newApplicationApi(agentid string) *application_api {

	app := application_api{corpId: cfg.Wxqy().Corpid, secret: cfg.WxqyApp(agentid).Secret, agentid: agentid}
	return &app
}

func (this *application_api) access_token() (string, error) {
	return newAccessToken().getToken(this.corpId, this.secret)
}
