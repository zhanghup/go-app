//go:generate go run cmd/generator.go

package ags

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	rice "github.com/GeertJohan/go.rice"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"github.com/zhanghup/go-app/service/ags/resolvers"
	"github.com/zhanghup/go-app/service/ags/source"
	"github.com/zhanghup/go-app/service/directive"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
	"github.com/zhanghup/go-tools/tgql"
	"io"
	"net/http"
	"strings"
	"time"
	"xorm.io/xorm"
)

func gqlschemaFmt(db *xorm.Engine, schema graphql.ExecutableSchema) func(c *gin.Context) {
	srv := handler.New(schema)
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			Error: func(w http.ResponseWriter, r *http.Request, status int, e error) {
				if e != nil {
					// 【事务】 统一提交关闭回归事务
					val := r.Context().Value(txorm.CONTEXT_SESSION)
					sess, ok := val.(txorm.ISession)
					if ok {
						sess.Rollback()
						sess.AutoClose()
					}
				}
			},
		},
		InitFunc: func(ctx context.Context, initPayload transport.InitPayload) (context.Context, error) {
			gtx := ctx.Value(directive.GIN_CONTEXT).(*gin.Context)
			_, err := directive.WebAuthFunc(db, gtx)
			return ctx, err
		},
	})
	// 【graphql】 错误处理
	srv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		if e != nil {
			// 【事务】 统一提交关闭回归事务
			val := ctx.Value(txorm.CONTEXT_SESSION)
			sess, ok := val.(txorm.ISession)
			if ok {
				sess.Rollback()
				sess.AutoClose()
			}
		}
		err := graphql.DefaultErrorPresenter(ctx, e)

		return err
	})
	srv.Use(extension.Introspection{})

	hu := tgql.DataLoadenMiddleware(db, srv)
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		// 统一建立session
		ctx = context.WithValue(ctx, txorm.CONTEXT_SESSION, txorm.NewEngine(db).Session(ctx))
		// 统一关联gin对象
		ctx = context.WithValue(ctx, directive.GIN_CONTEXT, c)

		c.Header("Content-Type", "application/json")
		hu.ServeHTTP(c.Writer, c.Request.WithContext(ctx))

		// 【事务】 统一提交关闭事务
		val := ctx.Value(txorm.CONTEXT_SESSION)
		sess, ok := val.(txorm.ISession)
		if ok {
			sess.Commit()
			sess.AutoClose()
		}
	}
}

func GinGql(gqlpath string, gqlrouter gin.IRouter, gqlSchema graphql.ExecutableSchema, db *xorm.Engine) {
	gqlrouter.POST(gqlpath, gqlschemaFmt(db, gqlSchema))
	gqlrouter.GET(gqlpath, gqlschemaFmt(db, gqlSchema))
	GinPlayground(gqlrouter, gqlpath+"/playground1", gqlpath)
	gqlrouter.GET(gqlpath+"/playground2", func(c *gin.Context) {
		playground.Handler("标题", gqlpath)(c.Writer, c.Request)
	})
}

func GinAgs(auth, any gin.IRouter) {
	GinGql("/zpx/ags", any, source.NewExecutableSchema(source.Config{Resolvers: resolvers.NewResolver(defaultDB)}), defaultDB)
	defaultUploader.GinRouter(auth.Group("/zpx/ags", directive.WebAuth(defaultDB)), any.Group("/zpx/ags"))

}

func GinStatic(box *rice.Box, g gin.IRouter, prefix string) {
	g.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(302, prefix)
	})
	g.GET("/"+prefix+"/*path", func(c *gin.Context) {

		path, _ := c.Params.Get("path")
		if tools.StrContains([]string{"/", "index.html"}, path) {
			path = "index.html"
		}
		if strings.Index(path,"/") == 0{
			path = path[1:]
		}

		f, err := box.Open(prefix + "/" + path)

		if err != nil {
			f, err = box.Open(prefix + "/index.html")
			if err != nil {
				return
			}
		}
		if path == "index.html" {
			c.Header("Content-Type", "text/html; charset=utf-8")
			io.Copy(c.Writer, f)
			return
		} else {
			stat, err := f.Stat()
			if err == nil {
				http.ServeContent(c.Writer, c.Request, c.Request.URL.Path, stat.ModTime(), f)
				return
			}

		}
		c.String(404, "404")
	})
}
