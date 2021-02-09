package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/cfg"
	"github.com/zhanghup/go-app/service/ca"
	"github.com/zhanghup/go-app/service/directive"
	"github.com/zhanghup/go-app/service/event"
	"github.com/zhanghup/go-tools"
)

func (r *mutationResolver) Login(ctx context.Context, username string, password string) (string, error) {
	acc := beans.Account{}
	// 1. 找到账户
	ok, err := r.DB.Where("username = ? and status = '1' and type = 'password'", username).Get(&acc)
	if err != nil {
		return "", err
	}
	if !ok {
		return "", errors.New("账户不存在")
	}

	// 2. 验证密码
	flag := false
	if acc.Salt == nil {
		if *acc.Password == password {
			flag = true
		}
	} else {
		if *acc.Password == tools.Crypto.Password(password, *acc.Salt) {
			flag = true
		}
	}
	if !flag {
		return "", errors.New("用户名或者密码错误")
	}

	// 3. 找到用户
	user := beans.User{}
	ok, err = r.DB.Where("id = ? and status = '1'", acc.Uid).Get(&user)
	if err != nil {
		return "", err
	}
	if !ok {
		return "", errors.New("用户不存在")
	}

	tok, err := r.Token(ctx, *user.Id, *acc.Id)
	if err != nil {
		return "", err
	} else {
		go event.UserLogin(acc, user)
	}
	return tok, nil
}

func (r *mutationResolver) LoginWxmp(ctx context.Context, code string) (string, error) {
	res, err := resty.New().R().Get(tools.StrTmp(`https://api.weixin.qq.com/sns/jscode2session?appid={{.appid}}&secret={{.appsecret}}&js_code={{.code}}&grant_type=authorization_code`, map[string]interface{}{
		"appid":     cfg.Config.Wxmp.Appid,
		"appsecret": cfg.Config.Wxmp.Appsecret,
		"code":      code,
	}).String())
	if err != nil {
		return "", err
	}

	restru := struct {
		Openid     string `json:"openid"`
		SessionKey string `json:"session_key"`
		Unionid    string `json:"unionid"`
		Errcode    int    `json:"errcode"`
		Errmsg     string `json:"errmsg"`
	}{}
	err = json.Unmarshal(res.Body(), &restru)
	if err != nil {
		return "", err
	}


	return "", nil

}

func (r *mutationResolver) Logout(ctx context.Context) (bool, error) {
	tok, err := r.Gin(ctx).Cookie(directive.GIN_TOKEN)
	if err != nil || tok == "" {
		tok = r.Gin(ctx).GetHeader(directive.GIN_AUTHORIZATION)
	}

	if tok == "" {
		return false, nil
	}
	_, err = r.DB.Table(beans.Token{}).Where("id = ?", tok).Update(map[string]interface{}{"status": 0})
	if err != nil {
		return false, err
	}
	ca.UserCache.RemoveByToken(tok)

	return true, err
}

func (r *queryResolver) LoginStatus(ctx context.Context) (bool, error) {
	_, err := directive.WebAuthFunc(r.DB, r.Gin(ctx))
	return err == nil, nil
}
