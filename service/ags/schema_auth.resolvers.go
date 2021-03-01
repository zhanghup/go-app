package ags

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/gs"
	"github.com/zhanghup/go-app/service/ca"
	"github.com/zhanghup/go-app/service/directive"
	"github.com/zhanghup/go-app/service/event"
	"github.com/zhanghup/go-tools"
)

func (r *mutationResolver) Login(ctx context.Context, username string, password string) (string, error) {
	acc := beans.Account{}
	// 1. 找到账户
	ok, err := gs.Sess(ctx).SF("select * from account where username = ? and status = '1' and type = 'password'", username).Get(&acc)
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
		if *acc.Password == tools.Password(password, *acc.Salt) {
			flag = true
		}
	}
	if !flag {
		return "", errors.New("用户名或者密码错误")
	}

	// 3. 找到用户
	user := beans.User{}
	ok, err = gs.Sess(ctx).SF("select * from user where id = ? and status = '1'", acc.Uid).Get(&user)
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
	if r.Wxmp == nil {
		return "", errors.New("微信接口未初始化!!!")
	}

	res, err := r.Wxmp.Code2Session(code)
	if err != nil {
		return "", err
	}

	user := beans.WxmpUser{}

	ok, err := gs.Sess(ctx).SF(`select * from wxmp_user where appid = ? and openid = ?`, gs.Config.Wxmp.Appid, res.Openid).Get(&user)
	if err != nil {
		return "", err
	}
	if !ok {
		user = beans.WxmpUser{
			Bean: beans.Bean{
				Id:     tools.PtrOfUUID(),
				Status: tools.PtrOfString("1"),
			},
			Appid:   &gs.Config.Wxmp.Appid,
			Openid:  &res.Openid,
			Unionid: &res.Unionid,
		}
		err := gs.Sess(ctx).Insert(user)
		if err != nil {
			return "", err
		}
		go event.WxmpUserCreate(user)
	}

	// Create the token
	claims := jwt.MapClaims{
		"uid":        *user.Id,
		"sessionKey": res.SessionKey,
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	t, err := token.SignedString([]byte(tools.MD5([]byte(gs.Config.Database.Uri))))

	if err != nil {
		return "", err
	} else {
		go event.WxmpUserLogin(user)
	}

	return t, nil
}

func (r *mutationResolver) Logout(ctx context.Context) (bool, error) {
	tok, err := gs.Gin(ctx).Cookie(gs.GIN_TOKEN)
	if err != nil || tok == "" {
		tok = gs.Gin(ctx).GetHeader(gs.GIN_AUTHORIZATION)
	}

	if tok == "" {
		return false, nil
	}
	_, err = gs.Sess(ctx).S().Table(beans.Token{}).Where("id = ?", tok).Update(map[string]interface{}{"status": 0})
	if err != nil {
		return false, err
	}
	ca.UserCache.RemoveByToken(tok)

	return true, err
}

func (r *queryResolver) LoginStatus(ctx context.Context) (bool, error) {
	_, err := directive.WebAuthFunc(gs.Gin(ctx))
	return err == nil, nil
}
