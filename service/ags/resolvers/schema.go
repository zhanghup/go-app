package resolvers

import (
	"context"
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/gs"
	"github.com/zhanghup/go-app/service/ags/source"
	"github.com/zhanghup/go-app/service/ca"
	"github.com/zhanghup/go-app/service/directive"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/tgql"
	"github.com/zhanghup/go-tools/wx/wxmp"
	"xorm.io/xorm"
)

type Resolver struct {
	Me     func(ctx context.Context) *ca.User
	Loader func(ctx context.Context) tgql.Loader
	Wxmp   wxmp.IEngine
}

func NewResolver(db *xorm.Engine) source.ResolverRoot {
	r := &Resolver{
		Me:     directive.MyInfo,
		Loader: tgql.DataLoaden,
	}

	if gs.Config.Wxmp.Appid != "" {
		r.Wxmp = wxmp.NewEngine(&gs.Config.Wxmp)
	}
	return r
}

func (this *mutationResolver) Token(ctx context.Context, uid, aid string) (string, error) {
	token := new(beans.Token)
	e := gs.Sess(ctx).SF(`update token set status = '0' where uid = :uid and aid = :aid`, map[string]interface{}{
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
	token.Agent = tools.PtrOfString(gs.Gin(ctx).Request.UserAgent())
	token.Expire = tools.PtrOfInt64(2 * 60 * 60)
	token.Ops = tools.PtrOfInt64(0)
	e = gs.Sess(ctx).Insert(token)
	if e != nil {
		return "", e
	}
	gs.Gin(ctx).SetCookie(gs.GIN_TOKEN, *token.Id, 2*60*60, "/", "", false, true)

	return *token.Id, nil
}
