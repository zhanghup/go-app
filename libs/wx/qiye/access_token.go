package qiye

import (
	"fmt"
	"github.com/zhanghup/go-tools"
)

type accessToken struct {
	Errorcode   int    `json:"errorcode"`
	Errmsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func getAccessToken(corpid, corpsecret string) (*accessToken, error) {
	res := new(accessToken)
	err := tools.Http().GetI("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid={{.corpid}}&corpsecret={{.corpsecret}}", map[string]interface{}{
		"corpid":     corpid,
		"corpsecret": corpsecret,
	}, res)
	return res, err
}

type Error struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func (this Error) Error() error {
	if this.Errcode == 0 {
		return nil
	}
	return fmt.Errorf("Error: %d, ErrorMessage: %s ", this.Errcode, this.Errmsg)
}

