package application

import (
	"github.com/zhanghup/go-app/libs/wx/qiye/api"
	"github.com/zhanghup/go-tools"
	"time"
)

type Option struct {
	CorpId     string `ini:"corpid"`
	CorpSecret string `ini:"corpsecret"`
}

type application struct {
	op    Option
	token tools.IMap
}

func newApplication(opt Option) *application {
	app := application{op: opt, token: tools.NewCache()}
	return &app
}

func (this *application) access_token() (string, error) {
	o := this.token.Get("token")
	if o != nil {
		return o.(string), nil
	}

	res, err := api.GetAccessToken(this.op.CorpId, this.op.CorpSecret)
	if err != nil {
		return "", err
	}
	this.token.Set2("token", res.AccessToken, time.Now().Unix()+int64(res.ExpiresIn))
	return res.AccessToken, nil
}
