package api

import (
	"errors"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"github.com/zhanghup/go-app"
	"github.com/zhanghup/go-app/service/auth"
	"github.com/zhanghup/go-app/service/gs"
)

func userAuth(e *xorm.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		tok, err := c.Cookie(gs.GIN_TOKEN)
		if err != nil {
			c.Fail(errors.New("【1:未授权】"), nil, 401)
			c.Abort()
			return
		}
		if len(tok) == 0 {
			tok = c.GetHeader("Authorization")
		}
		token := app.UserToken{}
		ok, err := e.Table(token).Where("id = ?", tok).Get(&token)
		if err != nil {
			c.Fail(errors.New("【2:未授权】"), nil, 401)
			c.Abort()
			return
		}
		if !ok {
			c.Fail(errors.New("【3:未授权】"), nil, 401)
			c.Abort()
			return
		}
		if token.Status == nil || *token.Status != 1 {
			c.Fail(errors.New("【4:未授权】"), nil, 401)
			c.Abort()
			return
		}
		if time.Now().Unix() > *token.Created+*token.Expire {
			c.Fail(errors.New("【5:未授权】"), nil, 401)
			c.Abort()
			return
		}
		if token.Type == nil || auth.TokenType(*token.Type) != auth.TokenPc {
			c.Fail(errors.New("【6:未授权】"), nil, 401)
			c.Abort()
			return
		}

		*token.Ops += 1
		*token.Expire = 7200
		_, err = e.Table(token).Where("id = ?", token.Id).Update(token)
		if err != nil {
			c.Fail(errors.New("【7:未授权】"), nil, 401)
			c.Abort()
			return
		}

		// 读取权限列表
		myPerms := gs.Perms{}
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
				c.Fail(errors.New("【8:未授权】"), nil, 401)
				c.Abort()
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
		myPermObj := gs.PermObjects{}
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
				c.Fail(errors.New("【9:未授权】"), nil, 401)
				c.Abort()
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

		u := app.User{}
		ok, err = e.Table(u).Where("id = ? and status = 1", *token.Uid).Get(&u)
		if err != nil {
			c.Fail(errors.New("【10:未授权】"), nil, 401)
			c.Abort()
			return
		}

		// 是否为管理员用户
		admin := true
		if u.Admin == nil || *u.Admin == 0 {
			admin = false
		}

		c.Set("uid", *token.Uid)
		c.Set("perms", myPerms)
		c.Set("permobjs", myPermObj)
		c.Set("admin", admin)
		c.SetCookie(gs.GIN_TOKEN, *token.Id, 2*60*60, "/", "", false, true)
		c.Next()
	}
}
