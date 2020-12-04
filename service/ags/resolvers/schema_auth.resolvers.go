package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"github.com/zhanghup/go-app/beans"
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

func (r *mutationResolver) Logout(ctx context.Context) (bool, error) {
	tok := r.Gin(ctx).GetHeader(directive.GIN_AUTHORIZATION)

	if tok == "" {
		tokk, err := r.Gin(ctx).Cookie(directive.GIN_TOKEN)
		if err != nil {
			return false, err
		}
		tok = tokk
	}
	if tok == "" {
		return false, nil
	}
	_, err := r.DB.Table(beans.Token{}).Where("id = ?", tok).Update(map[string]interface{}{"status": 0})
	if err != nil {
		return false, err
	}
	ca.UserCache.RemoveByToken(tok)

	return err == nil, err
}

func (r *queryResolver) LoginStatus(ctx context.Context) (bool, error) {
	_,err := directive.WebAuthFunc(r.DB,r.Gin(ctx))
	return err == nil,err
	//tok := r.Gin(ctx).GetHeader(directive.GIN_AUTHORIZATION)
	//
	//if tok == "" {
	//	tokk, err := r.Gin(ctx).Cookie(directive.GIN_TOKEN)
	//	if err != nil && err != http.ErrNoCookie {
	//		return false, err
	//	}
	//	tok = tokk
	//}
	//if tok == "" {
	//	return false, nil
	//}
	//t := beans.Token{}
	//ok, err := r.DB.Where("id = ? and status = 1", tok).Get(&t)
	//if err != nil {
	//	return false, err
	//}
	//if !ok {
	//	return false, nil
	//}
	//if time.Now().Unix() > *t.Updated+*t.Expire {
	//	return false, nil
	//}
	return true, nil
}
