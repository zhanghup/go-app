package qiye

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"time"
)

type access_token_api struct {
	token tools.IMap
}

var token *access_token_api

func newAccessToken() *access_token_api {
	if token != nil {
		return token
	}
	token = &access_token_api{token: tools.NewCache()}
	return token
}

type accessToken struct {
	Error
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func (this *access_token_api) getToken(corpid, corpsecret string) (string, error) {
	key := fmt.Sprintf("%s -- %s", corpid, corpsecret)
	ak := this.token.Get(key)
	if ak != nil {
		return ak.(string), nil
	}

	res := new(accessToken)
	err := tools.Http().GetI("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid={{.corpid}}&corpsecret={{.corpsecret}}", map[string]interface{}{
		"corpid":     corpid,
		"corpsecret": corpsecret,
	}, res)

	if err != nil {
		return "", err
	}
	if res.Error.Error() != nil {
		return "", res.Error.Error()
	}
	this.token.Set2(key, res.AccessToken, time.Now().Unix()+int64(res.ExpiresIn))
	return res.AccessToken, err
}
