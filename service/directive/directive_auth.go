package directive

import (
	"errors"
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/service/ca"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
	"github.com/zhanghup/go-tools/tgin"
	"strings"
	"time"
	"xorm.io/xorm"

	"github.com/gin-gonic/gin"
)

func WebAuth(db *xorm.Engine) gin.HandlerFunc {

	return func(c *gin.Context) {
		tgin.DoCustom(c, func(c *gin.Context) (interface{}, string) {
			res, err := WebAuthFunc(db, c)
			if err != nil {
				return res, err.Error()
			}
			return res, ""
		})
	}
}

func WebAuthFunc(db *xorm.Engine, c *gin.Context) (interface{}, error) {
	dbs := txorm.NewEngine(db)

	tok := c.GetHeader(GIN_AUTHORIZATION)

	if len(tok) == 0 {
		tok, _ = c.Cookie(GIN_TOKEN)
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
		_, err := db.Table(user.Token).Where("id = ?", user.TokenString).Update(user.Token)
		if err != nil {
			return err, errors.New("[3] 未授权")
		}
		ca.UserCache.Set(tok, user)
		c.Set("user_info", user)
		c.SetCookie(GIN_TOKEN, user.TokenString, 2*60*60, "/", "", false, true)
		return nil, nil
	}

	user = ca.User{}
	// token 验证
	{
		token := beans.Token{}
		ok, err := db.Table(token).Where("id = ?", tok).Get(&token)
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
		_, err = db.Table(token).Where("id = ?", token.Id).Update(token)
		if err != nil {
			return err.Error(), errors.New("[8] 未授权")
		}
		user.Token = token
		user.TokenString = *token.Id
	}

	// 用户验证
	{
		u := beans.User{}
		ok, err := db.Table(u).Where("id = ? and status = 1", *user.Token.Uid).Get(&u)
		if err != nil {
			return err.Error(), errors.New("[9] 未授权")
		}
		if !ok {
			return nil, errors.New("[10] 未授权")
		}
		user.User = u
	}
	// 是否为管理员用户
	{
		acc := beans.Account{}
		ok, err := db.Table(acc).Where("id = ? and status = 1", *user.Token.Aid).Get(&acc)
		if err != nil {
			return err.Error(), errors.New("[11] 未授权")
		}
		if !ok {
			return nil, errors.New("[12] 未授权")
		}
		user.Account = acc

		admin := true
		if acc.Admin == nil || *acc.Admin == 0 {
			admin = false
		}
		user.Admin = admin
	}

	if !user.Admin {
		// 读取权限列表
		myPerms := Perms{}
		{
			perms := make([]struct {
				Type string `json:"type"`
				Oid  string `json:"oid"`
			}, 0)
			err := dbs.SF(`
						select p.type,p.oid from user u 
						join role_user ru on u.id = ru.uid
						join perm p on p.role = ru.id
					`).Find(&perms)
			if err != nil {
				return err.Error(), errors.New("[11] 未授权")
			}
			// 去重
			for _, p := range perms {
				if o, ok := myPerms[p.Type]; ok {
					if !tools.Str.Contains(o, p.Oid) {
						o = append(o, p.Oid)
						myPerms[p.Type] = o
					}
				} else {
					myPerms[p.Type] = []string{p.Oid}
				}
			}
		}

		// 读取对象权限
		myPermObj := PermObjects{}
		{
			permObjects := make([]struct {
				Object string `json:"object"`
				Mask   string `json:"mask"`
			}, 0)
			err := dbs.SF(`
						select p.object,p.mask from user u 
						join role_user ru on u.id = ru.uid
						join perm_object p on p.role = ru.id
					`).Find(&permObjects)
			if err != nil {
				return err.Error(), errors.New("[12] 未授权")
			}
			for _, p := range permObjects {
				if o, ok := myPermObj[p.Object]; ok {
					myPermObj[p.Object] = o + p.Mask
				} else {
					myPermObj[p.Object] = p.Mask
				}
			}
			// myPermObj去重
			for k, v := range myPermObj {
				vs := strings.Split(v, "")
				t := map[string]bool{}
				for _, str := range vs {
					t[str] = true
				}
				str := ""
				for kk := range t {
					str += kk
				}
				myPermObj[k] = str
			}

		}

		user.Perms = myPerms
		user.PermObjects = myPermObj
	}

	c.Set("user_info", user)
	ca.UserCache.Set(user.TokenString, user)
	c.SetCookie(GIN_TOKEN, user.TokenString, 2*60*60, "/", "", false, true)
	c.Next()
	return nil, nil
}
