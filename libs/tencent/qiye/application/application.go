package application

import (
	"github.com/zhanghup/go-app/libs/tencent/qiye/api"
	"github.com/zhanghup/go-tools"
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

func (this *application) access_token() error {
	res, err := api.GetAccessToken(this.op.CorpId, this.op.CorpSecret)

}
