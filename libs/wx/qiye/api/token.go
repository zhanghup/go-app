package api

import "github.com/zhanghup/go-tools"

const (
	ACCESS_TOKEN = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid={{.corpid}}&corpsecret={{.corpsecret}}"
)

type AccessToken struct {
	Errorcode   int    `json:"errorcode"`
	Errmsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func GetAccessToken(corpid, corpsecret string) (*AccessToken, error) {
	res := new(AccessToken)
	err := tools.Http().GetI(ACCESS_TOKEN, map[string]interface{}{
		"corpid":     corpid,
		"corpsecret": corpsecret,
	}, res)
	return res, err
}
