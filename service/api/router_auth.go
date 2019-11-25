package api

import (
	"errors"
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

		c.Set("uid", *token.Uid)
		c.SetCookie(gs.GIN_TOKEN, *token.Id, 2*60*60, "/", "", false, true)
		c.Next()
	}
}
