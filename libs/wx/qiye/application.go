package qiye

import (
	"fmt"
	"github.com/zhanghup/go-app/cfg"
	"github.com/zhanghup/go-app/libs/wx/qiye/api"
	"github.com/zhanghup/go-tools"
	"time"
)

type application struct {
	corpId  string
	secret  string
	agentid string

	token tools.IMap
}

func newApplication(agentid string) *application {

	if !cfg.WxQyApp(agentid).Enable {
		panic(fmt.Sprintf("config.ini - wxqy-app[%s].enable 未启用", agentid))
	}
	app := application{corpId: cfg.WxQy().Corpid, secret: cfg.WxQyApp(agentid).Secret, agentid: agentid, token: tools.NewCache()}
	return &app
}

func (this *application) access_token() (string, error) {
	o := this.token.Get("token")
	if o != nil {
		return o.(string), nil
	}

	res, err := getAccessToken(this.corpId, this.secret)
	if err != nil {
		return "", err
	}
	this.token.Set2("token", res.AccessToken, time.Now().Unix()+int64(res.ExpiresIn))
	return res.AccessToken, nil
}
