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
	"github.com/zhanghup/go-app/gs"
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
)

func gqlschemaFmt(schema graphql.ExecutableSchema) func(c *gin.Context) {
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
			gtx := ctx.Value(gs.GIN_CONTEXT).(*gin.Context)
			_, err := directive.WebAuthFunc(gtx)
			return ctx, err
		},
	})
	srv.Use(extension.Introspection{})
	srv.AroundResponses(func(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
		// 统一建立session
		sess := gs.DBA().Session(ctx)
		ctx = context.WithValue(ctx, txorm.CONTEXT_SESSION, sess)
		res := next(ctx)
		if len(res.Errors) > 0 {
			sess.Rollback()
			sess.AutoClose()
		} else {
			sess.Commit()
			sess.AutoClose()
		}

		return res
	})

	hu := tgql.DataLoadenMiddleware(gs.DB(), srv)
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		// 统一关联gin对象
		ctx = context.WithValue(ctx, gs.GIN_CONTEXT, c)
		c.Header("Content-Type", "application/json")
		hu.ServeHTTP(c.Writer, c.Request.WithContext(ctx))
	}
}

func GinGql(gqlpath string, gqlrouter gin.IRouter, gqlSchema graphql.ExecutableSchema) {
	gqlrouter.POST(gqlpath, gqlschemaFmt(gqlSchema))
	gqlrouter.GET(gqlpath, gqlschemaFmt(gqlSchema))
	gs.GinPlayground(gqlrouter, gqlpath+"/playground1", gqlpath)
	gqlrouter.GET(gqlpath+"/playground2", func(c *gin.Context) {
		playground.Handler("标题", gqlpath)(c.Writer, c.Request)
	})
}

func GinAgs(auth, any gin.IRouter) {
	GinGql("/zpx/ags", any, source.NewExecutableSchema(source.Config{Resolvers: resolvers.NewResolver(gs.DB())}))
	gs.Uploader().GinRouter(auth.Group("/zpx/ags", directive.WebAuth()), any.Group("/zpx/ags"))

}

func GinStatic(box *rice.Box, g gin.IRouter, prefix string) {
	g.GET("/"+prefix+"/*path", func(c *gin.Context) {
		path, _ := c.Params.Get("path")
		if tools.StrContains([]string{"/", "index.html"}, path) {
			path = "index.html"
		}
		if strings.Index(path, "/") == 0 {
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
