module github.com/zhanghup/go-app

go 1.12

require (
	github.com/gin-gonic/gin v1.4.0
	github.com/graphql-go/graphql v0.7.8 // indirect
	github.com/graphql-go/handler v0.2.3 // indirect
)

replace github.com/graphql-go/graphql => github.com/zhanghup/go-graphql v0.0.0-20190403165646-199d20bbfed7

replace github.com/gin-gonic/gin => github.com/zhanghup/go-gin v0.0.0-20190818132815-0b17e4c65c1e

replace github.com/go-xorm/xorm => github.com/zhanghup/go-xorm v0.7.6
