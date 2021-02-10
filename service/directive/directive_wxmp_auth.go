package directive

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/cfg"
	"github.com/zhanghup/go-app/service/ca"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
	"github.com/zhanghup/go-tools/tgin"
	"xorm.io/xorm"

	"github.com/gin-gonic/gin"
)

func WxmpAuth(db *xorm.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		tgin.DoCustom(c, func(c *gin.Context) (interface{}, string) {
			res, err := WxmpAuthFunc(db, c)
			if err != nil {
				return res, err.Error()
			}
			return res, ""
		})
	}
}

func WxmpAuthFunc(db *xorm.Engine, c *gin.Context) (interface{}, error) {
	dbs := txorm.NewEngine(db)

	tok, _ := c.Cookie(GIN_TOKEN)

	if len(tok) == 0 {
		tok = c.GetHeader(GIN_AUTHORIZATION)
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
			return []byte(tools.Crypto.MD5([]byte(cfg.DB.Uri))), nil
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
		c.Set(GIN_WXUSER, user)
		c.Next()
		return nil, nil
	}

	wxuser := beans.WxmpUser{}
	ok, err := dbs.SF(`select * from wxmp_user where id = ?`, form.Uid).Get(&wxuser)
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

	c.Set(GIN_WXUSER, wxuser)
	ca.WxuserCache.Set(user.Id, user)
	c.Next()
	return nil, nil
}
