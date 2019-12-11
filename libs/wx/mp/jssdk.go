package mp

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"net/url"
)

type UserInfo struct {
	Openid string `json:"openid"`
	Nickname string `json:"nickname"`
	Sex int `json:"sex"`
	Province string `json:"province"`
	City string `json:"city"`
	Country string `json:"country"`
	Headimgurl string `json:"headimgurl"`
	Privilege []string `json:"privilege"`
	Unionid string `json:"unionid"`
}

func (this *context) JSSDK() IJssdk{
	return &jssdk{
		context: this,
	}
}

type IJssdk interface {
	AuthUrl(uri,scope,state string) (string,error)
	Auth(code string) (string,string,error)
	AuthUserInfo(code string) (*UserInfo,error)
}

type jssdk struct {
	context *context
}

func (this *jssdk) error(err interface{},fn string,i ... int) error{
	var s = ""
	switch err.(type) {
	case string:
		s = err.(string)
	case error:
		s = err.(error).Error()
	case Error:
		e := err.(Error)
		if e.Errcode == 0{
			return nil
		}
		s = fmt.Sprintf("%d: %s",e.Errcode,e.Errmsg)
	}

	if len(i) > 0{
		return fmt.Errorf("微信公众号 - 网页开发_%d - %s - %s",i[0],fn,s)
	}
	return fmt.Errorf("微信公众号 - 网页开发 - %s - %s",fn,s)
}

func (this *jssdk) Auth(code string) (string,string,error) {
	tok := struct {
		Error
		AccessToken string `json:"access_token"`
		ExpiresIn int `json:"expires_in"`
		Openid string `json:"openid"`
	}{}

	err := tools.Http().GetI("https://api.weixin.qq.com/sns/oauth2/access_token?appid={{.appid}}&secret={{.secret}}&code={{.code}}&grant_type=authorization_code", map[string]interface{}{
		"appid":this.context.appid,
		"secret":this.context.appsecret,
		"code": code,
	},&tok)

	if err != nil{
		return "","",this.error(err,"Auth",1)
	}

	return tok.Openid,tok.AccessToken,this.error(tok.Error,"Auth",2)
}
func (this *jssdk) AuthUserInfo(code string) (*UserInfo,error){
	r := struct {
		*UserInfo
		Error
	}{}

	openid,token,err := this.Auth(code)
	if err != nil{
		return nil,err
	}

	err = tools.Http().GetI("https://api.weixin.qq.com/sns/userinfo?access_token={{.token}}&openid={{.openid}}&lang=zh_CN", map[string]interface{}{
		"token": token,
		"openid": openid,
	},&r)

	if err != nil{
		return nil,this.error(err,"AuthUserInfo",1)
	}

	return r.UserInfo,this.error(r.Error,"AuthUserInfo",2)

}

func (this *jssdk) AuthUrl(uri,scope,state string) (string,error){
	uri = url.QueryEscape(uri)

	if !tools.Str().StrContains([]string{"snsapi_base","snsapi_userinfo"},scope){
		return "",this.error("授权类型错误，应为“snsapi_base” 或者 “snsapi_userinfo”","AuthUrl")
	}

	return fmt.Sprintf("https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s#wechat_redirect",this.context.appid,uri,scope,state),nil
}

