package directive

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/gs"
	"github.com/zhanghup/go-app/service/ca"
	"github.com/zhanghup/go-tools/tgin"
	"time"
)

func WebAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tgin.DoCustom(c, func(c *gin.Context) (interface{}, string) {
			res, err := WebAuthFunc(c)
			if err != nil {
				return res, err.Error()
			}
			return res, ""
		})
	}
}

func WebAuthFunc(c *gin.Context) (interface{}, error) {
	tok, _ := c.Cookie(gs.GIN_TOKEN)

	if len(tok) == 0 {
		tok = c.GetHeader(gs.GIN_AUTHORIZATION)
	}
	if len(tok) == 0 {
		return nil, errors.New("[1] 未授权")
	}

	user, ok := ca.UserCache.Get(tok)
	if ok {
		if time.Now().Unix() > *user.Token.Updated+*user.Token.Expire {
			return nil, errors.New("[2] 未授权")
		}
		*user.Token.Ops += 1
		*user.Token.Expire = 7200

		// 5秒内的token变化不记录
		if time.Now().Unix()-*user.Token.Updated > 5 {
			*user.Token.Updated = time.Now().Unix()
			_, err := gs.DB().Table(user.Token).Where("id = ?", user.TokenString).Update(user.Token)
			if err != nil {
				return err, errors.New("[3] 未授权")
			}
		}
		ca.UserCache.Set(tok, user)
		c.Set(gs.GIN_USER, user)
		c.SetCookie(gs.GIN_TOKEN, user.TokenString, 2*60*60, "/", "", false, true)
		return nil, nil
	}

	user = ca.User{}
	// token 验证
	{
		token := beans.Token{}
		ok, err := gs.DB().Table(token).Where("id = ?", tok).Get(&token)
		if err != nil {
			return err.Error(), errors.New("[4] 未授权")
		}
		if !ok {
			return nil, errors.New("[5] 未授权")
		}
		if token.Status == nil || *token.Status != "1" {
			return nil, errors.New("[6] 未授权")
		}
		if time.Now().Unix() > *token.Updated+*token.Expire {
			return nil, errors.New("[7] 未授权")
		}
		*token.Ops += 1
		*token.Expire = 7200
		_, err = gs.DB().Table(token).Where("id = ?", token.Id).Update(token)
		if err != nil {
			return err.Error(), errors.New("[8] 未授权")
		}
		user.Token = token
		user.TokenString = *token.Id
	}

	// 用户验证
	{
		u := beans.User{}
		ok, err := gs.DB().Table(u).Where("id = ? and status = 1", *user.Token.Uid).Get(&u)
		if err != nil {
			return err.Error(), errors.New("[9] 未授权")
		}
		if !ok {
			return nil, errors.New("[10] 未授权")
		}
		user.User = u
		if (u.Admin != nil && *u.Admin == "1") || *u.Id == "root" {
			user.Admin = true
		} else {
			user.Admin = false
		}
		user.Id = *u.Id
		if u.Name != nil {
			user.Name = *u.Name
		}
	}

	if !user.Admin {
		// 读取对象权限
		permObjects := make([]struct {
			Object string `json:"object"`
			Mask   string `json:"mask"`
		}, 0)
		err := gs.DBS(gs.Background).SF(`
				select p.object,p.mask from user u 
				join role_user ru on u.id = ru.uid
				join perm_object p on p.role = ru.role
				where u.id = :uid
			`, map[string]interface{}{
			"uid": user.User.Id,
		}).Find(&permObjects)
		if err != nil {
			return err.Error(), errors.New("[12] 未授权")
		}
		for i := range permObjects {
			user.EntityPermAdd(permObjects[i].Object, permObjects[i].Mask)
		}
	} else {

	}

	c.Set(gs.GIN_USER, user)
	ca.UserCache.Set(user.TokenString, user)
	c.SetCookie(gs.GIN_TOKEN, user.TokenString, 2*60*60, "/", "", false, true)
	c.Next()
	return nil, nil
}
