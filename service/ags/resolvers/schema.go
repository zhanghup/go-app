package resolvers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/cfg"
	"github.com/zhanghup/go-app/service/ags/source"
	"github.com/zhanghup/go-app/service/ca"
	"github.com/zhanghup/go-app/service/directive"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
	"github.com/zhanghup/go-tools/tgql"
	"github.com/zhanghup/go-tools/tog"
	"github.com/zhanghup/go-tools/wx/wxmp"
	"xorm.io/xorm"
)

type Resolver struct {
	DBS    func(ctx context.Context) txorm.ISession
	Sess   func(ctx context.Context) txorm.ISession
	Gin    func(g context.Context) *gin.Context
	Me     func(ctx context.Context) *ca.User
	Loader func(ctx context.Context) tgql.Loader
	Wxmp   wxmp.IEngine
}

func NewResolver(db *xorm.Engine) source.ResolverRoot {
	dbs := txorm.NewEngine(db)
	r := &Resolver{
		DBS: func(ctx context.Context) txorm.ISession {
			return dbs.NewSession(true, ctx)
		},
		Sess: func(ctx context.Context) txorm.ISession {
			sess := dbs.Session(ctx)
			err := sess.Begin()
			if err != nil {
				tog.Error("【开启事务异常！！！】")
			}
			return sess
		},
		Gin: func(g context.Context) *gin.Context {
			gg := g.Value(directive.GIN_CONTEXT)
			ggg := gg.(*gin.Context)
			return ggg
		},
		Me:     directive.MyInfo,
		Loader: tgql.DataLoaden,
	}

	if cfg.Wxmp.Appid != "" {
		r.Wxmp = wxmp.NewEngine(&cfg.Wxmp)
	}

	return r
}

func (this *mutationResolver) Token(ctx context.Context, uid, aid string) (string, error) {
	token := new(beans.Token)
	e := this.Sess(ctx).SF(`update token set status = '0' where uid = :uid and aid = :aid`, map[string]interface{}{
		"uid": uid,
		"aid": aid,
	}).Exec()
	if e != nil {
		return "", e
	}
	token.Id = tools.PtrOfUUID()
	token.Status = tools.PtrOfString("1")
	token.Uid = &uid
	token.Aid = &aid
	token.Agent = tools.PtrOfString(this.Gin(ctx).Request.UserAgent())
	token.Expire = tools.PtrOfInt64(2 * 60 * 60)
	token.Ops = tools.PtrOfInt64(0)
	e = this.Sess(ctx).Insert(token)
	if e != nil {
		return "", e
	}
	this.Gin(ctx).SetCookie(directive.GIN_TOKEN, *token.Id, 2*60*60, "/", "", false, true)

	return *token.Id, nil
}
