//go:generate go run cmd/generator.go

package awxmp

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/gin-gonic/gin"
	"github.com/zhanghup/go-app/cfg"
	"github.com/zhanghup/go-app/service/ags"
	"github.com/zhanghup/go-app/service/awxmp/source"
	"github.com/zhanghup/go-app/service/ca"
	"github.com/zhanghup/go-app/service/directive"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
	"github.com/zhanghup/go-tools/tgql"
	"github.com/zhanghup/go-tools/tog"
	"github.com/zhanghup/go-tools/wx/wxmp"
	"reflect"
	"time"
)

func NewResolver(wxEngine wxmp.IEngine) *Resolver {
	return &Resolver{
		NewResolverTools(wxEngine),
	}
}

func Gin(g gin.IRouter, sc ...graphql.ExecutableSchema) {
	wxEngine := wxmp.NewEngine(&cfg.Wxmp)
	s := source.NewExecutableSchema(source.Config{
		Resolvers: NewResolver(wxEngine),
		Directives: source.DirectiveRoot{
		},
	})
	if len(sc) > 0 {
		s = sc[0]
	}

	ags.GinGql("/zpx/wxmp", g.Group("/", directive.WxmpAuth(ags.DefaultDB())), s, ags.DefaultDB())
	g.POST("/zpx/wxmp/pay/callback", PayCallback(wxEngine))
}

type Resolver struct {
	*ResolverTools
}

type ResolverTools struct {
	DBS    func(ctx context.Context) txorm.ISession
	Sess   func(ctx context.Context) txorm.ISession
	Loader func(ctx context.Context) tgql.Loader
	Me     func(ctx context.Context) *ca.WxmpUser
	Wxmp   wxmp.IEngine
}

func NewResolverTools(wxEngine wxmp.IEngine) *ResolverTools {
	dbs := txorm.NewEngine(ags.DefaultDB())
	return &ResolverTools{
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
		Loader: tgql.DataLoaden,
		Me:     directive.MyWxmpUser,
		Wxmp:   wxEngine,
	}
}

func (this *ResolverTools) Create(ctx context.Context, tab interface{}, obj interface{}) (string, error) {
	id := ""
	err := this.Sess(ctx).TS(func(sess txorm.ISession) error {
		tools.Rft.DeepSet(tab, func(t reflect.Type, v reflect.Value, tf reflect.StructField) bool {
			switch tf.Name {
			case "Id":
				if t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.String {
					if v.Pointer() == 0 {
						id = tools.UUID()
						v.Set(reflect.ValueOf(&id))
					}
				} else if t.Kind() == reflect.String {
					if v.String() == "" {
						v.Set(reflect.ValueOf(tools.UUID()))
					}
				}
			case "Status":
				if t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.String {
					if v.Pointer() == 0 {
						v.Set(reflect.ValueOf(tools.PtrOfString("1")))
					}
				}
			case "Weight":
				if t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Int {
					if v.Pointer() == 0 {
						v.Set(reflect.ValueOf(tools.PtrOfInt(int(time.Now().Unix() - 1610541047))))
					}
				}
			}
			return true
		})

		_, err := sess.S().Insert(tab)
		if err != nil {
			return err
		}
		if obj != nil {
			_, err = sess.S().Table(tab).Where("id = ?", id).Update(obj)
		}
		return err
	})

	return id, err
}

func (this *ResolverTools) Update(ctx context.Context, tab interface{}, id string, obj interface{}) (bool, error) {
	err := this.Sess(ctx).TS(func(sess txorm.ISession) error {
		_, err := sess.S().Table(tab).Where("id = ?", id).Update(tab)
		if err != nil {
			return err
		}
		if obj != nil {
			_, err = sess.S().Table(tab).Where("id = ?", id).AllCols().Update(obj)
		}
		return err
	})

	return err == nil, err
}

func (this *ResolverTools) Removes(ctx context.Context, table interface{}, ids []string) (bool, error) {
	err := this.Sess(ctx).TS(func(sess txorm.ISession) error {
		_, err := sess.S().Table(table).In("id", ids).Delete(table)
		return err
	})
	return err == nil, err
}
