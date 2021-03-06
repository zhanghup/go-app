package directive

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/gs"
	"github.com/zhanghup/go-app/service/ca"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/tgin"
)

func WxmpAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tgin.DoCustom(c, func(c *gin.Context) (interface{}, string) {
			res, err := WxmpAuthFunc(c)
			if err != nil {
				return res, err.Error()
			}
			return res, ""
		})
	}
}

func WxmpAuthFunc(c *gin.Context) (interface{}, error) {
	tok, _ := c.Cookie(gs.GIN_TOKEN)

	if len(tok) == 0 {
		tok = c.GetHeader(gs.GIN_AUTHORIZATION)
	}
	if len(tok) == 0 {
		return nil, errors.New("[1] 未授权")
	}

	form := struct {
		Uid        string `json:"uid"`
		SessionKey string `json:"sessionKey"`
		jwt.MapClaims
	}{}
	// token 验证
	{
		token, err := jwt.ParseWithClaims(tok, &form, func(token *jwt.Token) (interface{}, error) {
			return []byte(tools.MD5([]byte(gs.Config.Database.Uri))), nil
		})

		if err != nil {
			return nil, errors.New("[2] 未授权")
		}

		if !token.Valid {
			return nil, errors.New("[3] 未授权")
		}
	}

	user, ok := ca.WxuserCache.Get(form.Uid)
	if ok {
		ca.WxuserCache.Set(tok, user)
		c.Set(gs.GIN_WXUSER, user)
		c.Next()
		return nil, nil
	}

	wxuser := beans.WxmpUser{}
	ok, err := gs.DBS(gs.Background).SF(`select * from wxmp_user where id = ?`, form.Uid).Get(&wxuser)
	if err != nil {
		return nil, errors.New("[4] 未授权")
	}
	if !ok {
		return nil, errors.New("[5] 未授权")
	}

	// 数据格式化
	{
		user.User = &wxuser
		user.SessionKey = form.SessionKey
		if wxuser.Id != nil {
			user.Id = *wxuser.Id
		} else {
			user.Id = ""
		}

		if wxuser.Mobile != nil {
			user.Mobile = *wxuser.Mobile
		} else {
			user.Mobile = ""
		}

		if wxuser.Appid != nil {
			user.Appid = *wxuser.Appid
		} else {
			user.Appid = ""
		}

		if wxuser.Openid != nil {
			user.Openid = *wxuser.Openid
		} else {
			user.Openid = ""
		}

		if wxuser.Unionid != nil {
			user.Unionid = *wxuser.Unionid
		} else {
			user.Unionid = ""
		}

		user.TokenString = tok
	}

	c.Set(gs.GIN_WXUSER, user)
	ca.WxuserCache.Set(user.Id, user)
	c.Next()
	return nil, nil
}
