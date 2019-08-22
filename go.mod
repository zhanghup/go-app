module github.com/zhanghup/go-app

go 1.12

require (
	github.com/99designs/gqlgen v0.9.3
	github.com/gin-gonic/gin v1.4.0
	github.com/go-xorm/xorm v0.7.6
	github.com/mattn/go-sqlite3 v1.11.0
	github.com/stretchr/testify v1.4.0
	github.com/vektah/gqlparser v1.1.2
	github.com/zhanghup/go-tools v1.0.9
	golang.org/x/tools v0.0.0-20190822000311-fc82fb2afd64 // indirect
)

replace github.com/99designs/gqlgen => github.com/zhanghup/go-gqlgen v0.9.3

replace github.com/gin-gonic/gin => github.com/zhanghup/go-gin v0.0.0-20190818132815-0b17e4c65c1e

replace github.com/go-xorm/xorm => github.com/zhanghup/go-xorm v0.0.0-20190821094440-7a6035ced953
