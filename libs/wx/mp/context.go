package mp

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zhanghup/go-app/cfg"
	"github.com/zhanghup/go-tools"
	"io/ioutil"
	"net/http"
	"time"
)

type IContext interface {

}

type Error struct {
	Errcode int `json:"errcode"`
	Errmsg string `json:"errmsg"`
}

type context struct {
	appid string
	appsecret string
	cache tools.IMap
}

func NewContext() IContext{
	if !cfg.WxmpEnable(){
		panic("[wxmp] - 微信公众号配置没有初始化")
	}
	c := new(context)
	c.cache = tools.NewCache()
	c.appid = cfg.Wxmp().Appid
	c.appsecret = cfg.Wxmp().AppSecret
	return c
}

func (this *context) token() error {
	if this.cache.Contain("access_token"){
		return nil
	}

	res,err := http.Get(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s",this.appid,this.appsecret))
	if err != nil{
		return err
	}
	data,err := ioutil.ReadAll(res.Body)
	if err != nil{
		return err
	}
	tok := struct {
		Error
		AccessToken string `json:"access_token"`
		ExpiresIn int `json:"expires_in"`
	}{}

	err = json.Unmarshal(data,&tok)
	if err != nil{
		return err
	}
	if tok.Errcode != 0{
		return errors.New(fmt.Sprintf("%d: %s",tok.Errcode,tok.Errmsg))
	}
	this.cache.Set2("access_token",tok.AccessToken,time.Now().Unix() + int64(tok.ExpiresIn))
	return nil
}

func (this *context) request(url string,result interface{}) error {
	http.Get(url)
	return nil
}
