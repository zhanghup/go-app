package resolvers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/service/ags/source"
	"github.com/zhanghup/go-app/service/directive"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
	"xorm.io/xorm"
)

type Resolver struct {
	DB  *xorm.Engine
	DBS *txorm.Engine
	Gin func(g context.Context) *gin.Context
	Me  func(ctx context.Context) directive.Me
}

func NewResolver(db *xorm.Engine) source.ResolverRoot {
	return &Resolver{
		DB:  db,
		DBS: txorm.NewEngine(db),
		Gin: func(g context.Context) *gin.Context {
			gg := g.Value(directive.GIN_CONTEXT)
			ggg := gg.(*gin.Context)
			return ggg
		},
		Me: directive.MyInfo,
	}
}

func (this *mutationResolver) Token(ctx context.Context, uid, aid string) (string, error) {
	token := new(beans.Token)
	err := this.DBS.TS(func(sess *txorm.Session) error {
		e := sess.SF(`update token set status = '0' where uid = :uid and aid = :aid`, map[string]interface{}{
			"uid": uid,
			"aid": aid,
		}).Exec()
		if e != nil {
			return e
		}
		token.Id = tools.Ptr.Uid()
		token.Status = tools.Ptr.String("1")
		token.Uid = &uid
		token.Aid = &aid
		token.Agent = tools.Ptr.String(this.Gin(ctx).Request.UserAgent())
		token.Expire = tools.Ptr.Int64(2 * 60 * 60)
		token.Ops = tools.Ptr.Int64(0)
		e = sess.Insert(token)
		if e != nil {
			return e
		}
		this.Gin(ctx).SetCookie(directive.GIN_TOKEN, *token.Id, 2*60*60, "/", "", false, true)
		return nil

	})

	return *token.Id, err
}
