package directive

import (
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
	"github.com/zhanghup/go-tools/tgin"
	"strings"
	"time"
	"xorm.io/xorm"

	"github.com/gin-gonic/gin"
)

func WebAuth(db *xorm.Engine) gin.HandlerFunc {
	dbs := txorm.NewEngine(db)
	return func(c *gin.Context) {
		tgin.DoCustom(c, func(c *gin.Context) (interface{}, string) {
			tok, err := c.Cookie(GIN_TOKEN)
			if err != nil {
				return err.Error(), "[1] 未授权"
			}
			if len(tok) == 0 {
				tok = c.GetHeader("Authorization")
			}
			if len(tok) == 0 {
				return nil, "[2] 未授权"
			}
			token := beans.UserToken{}
			ok, err := db.Table(token).Where("id = ?", tok).Get(&token)
			if err != nil {
				return err.Error(), "[3] 未授权"
			}
			if !ok {
				return nil, "[4] 未授权"
			}
			if token.Status == nil || *token.Status != 1 {
				return nil, "[5] 未授权"
			}
			if time.Now().Unix() > *token.Updated+*token.Expire {
				return nil, "[6] 未授权"
			}
			*token.Ops += 1
			*token.Expire = 7200
			_, err = db.Table(token).Where("id = ?", token.Id).Update(token)
			if err != nil {
				return err.Error(), "[7] 未授权"
			}

			u := beans.User{}
			ok, err = db.Table(u).Where("id = ? and status = 1", *token.Uid).Get(&u)
			if err != nil {
				return err.Error(), "[10] 未授权"
			}
			// 是否为管理员用户
			admin := true
			if u.Admin == nil || *u.Admin == 0 {
				admin = false
			}

			if !admin {
				// 读取权限列表
				myPerms := Perms{}
				{
					perms := make([]struct {
						Type string `json:"type"`
						Oid  string `json:"oid"`
					}, 0)
					err = dbs.SF(`
						select p.type,p.oid from user u 
						join role_user ru on u.id = ru.uid
						join perm p on p.role = ru.id
					`).Find(&perms)
					if err != nil {
						return err.Error(), "[8] 未授权"
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
					err = dbs.SF(`
						select p.object,p.mask from user u 
						join role_user ru on u.id = ru.uid
						join perm_object p on p.role = ru.id
					`).Find(&permObjects)
					if err != nil {
						return err.Error(), "[9] 未授权"
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
				c.Set("perms", myPerms)
				c.Set("permobjs", myPermObj)
			}

			c.Set("uid", *token.Uid)
			c.Set("user", u)

			c.Set("admin", admin)
			c.SetCookie(GIN_TOKEN, *token.Id, 2*60*60, "/", "", false, true)
			c.Next()
			return nil, ""
		})
	}
}
