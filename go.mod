module github.com/zhanghup/go-app

go 1.12

require (
	github.com/gin-gonic/gin v1.4.0
	github.com/go-xorm/xorm v0.7.6
	github.com/graph-gophers/graphql-go v0.0.0-20190724201507-010347b5f9e6
)

replace github.com/graph-gophers/graphql-go => github.com/zhanghup/go-graphql v0.0.0-20190724201507-010347b5f9e6

replace github.com/gin-gonic/gin => github.com/zhanghup/go-gin v0.0.0-20190818132815-0b17e4c65c1e

replace github.com/go-xorm/xorm => github.com/zhanghup/go-xorm v0.7.6
