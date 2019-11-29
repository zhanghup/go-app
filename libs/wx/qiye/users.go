package qiye

import (
	"github.com/zhanghup/go-tools"
)

type user_api struct {
	corpId string
	secret string

	token tools.IMap
}

func (this *user_api) User(uid string) error {
	token, err := this.access_token()
	if err != nil {
		return err
	}
	result := Error{}
	err = tools.Http().GetI("https://qyapi.weixin.qq.com/cgi-bin/user/get?access_token={{.token}}&userid={{.userid}}", map[string]interface{}{
		"token":  token,
		"userid": uid,
	}, &result)
	if err != nil {
		return err
	}
	return result.Error()
}

func (this *user_api) access_token() (string, error) {
	return newAccessToken().getToken(this.corpId, this.secret)
}
