package directive

import (
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/toolgin"
	"strings"
	"time"
	"xorm.io/xorm"

	"github.com/gin-gonic/gin"
)

func WebAuth(db *xorm.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		toolgin.Do(c, func(c *gin.Context) (i interface{}, s string) {
			tok, err := c.Cookie(GIN_TOKEN)
			if err != nil {
				return err.Error(), "未授权【1】"
			}
			if len(tok) == 0 {
				tok = c.GetHeader("Authorization")
			}
			if len(tok) == 0 {
				return nil, "未授权【2】"
			}
			tok = tools.Crypto.DES(tok, "12345678").ECBDecrypt()
			if len(tok) == 0 {
				return nil, "未授权【3】"
			}

			ss := strings.Split(tok, "&")
			if len(ss) != 2 {
				return nil, "未授权【4】"
			}
			uid := ss[0]
			expire := tools.Str.MustInt64(ss[1])
			if expire == 0 {
				return nil, "未授权【5】"
			}

			if time.Now().Unix() > expire {
				return nil, "未授权【6】"
			}

			// 读取权限列表
			myPerms := Perms{}
			{
				perms := make([]struct {
					Type string `json:"type"`
					Oid  string `json:"oid"`
				}, 0)
				err = e.SF(`
					select p.type,p.oid from {{ table "user" }} u 
					join {{ table "role_user" }} ru on u.id = ru.uid
					join {{ table "perm" }} p on p.role = ru.id
				`).Find(&perms)
				if err != nil {
					c.Fail401("【8:未授权】", err)
					return
				}
				// 去重
				fn := func(strs []string, str string) bool {
					for _, s := range strs {
						if s == str {
							return true
						}
					}
					return false
				}

				for _, p := range perms {
					if o, ok := myPerms[p.Type]; ok {
						if !fn(o, p.Oid) {
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
				err = e.SF(`
			select p.object,p.mask from {{ table "user" }} u 
			join {{ table "role_user" }} ru on u.id = ru.uid
			join {{ table "perm_object" }} p on p.role = ru.id
		`).Find(&permObjects)
				if err != nil {
					c.Fail401("【9:未授权】", err)
					return
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

			u := beans.User{}
			ok, err = e.Table(u).Where("id = ? and status = 1", *token.Uid).Get(&u)
			if err != nil {
				c.Fail401("【10:未授权】", err)
				return
			}

			// 是否为管理员用户
			admin := true
			if u.Admin == nil || *u.Admin == 0 {
				admin = false
			}

			c.Set("uid", *token.Uid)
			c.Set("user", u)
			c.Set("perms", myPerms)
			c.Set("permobjs", myPermObj)
			c.Set("admin", admin)
			c.SetCookie(GIN_TOKEN, *token.Id, 2*60*60, "/", "", false, true)
			c.Next()

		})

	}
}
