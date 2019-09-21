module github.com/zhanghup/go-app

go 1.12

require (
	github.com/99designs/gqlgen v0.9.3
	github.com/gin-gonic/gin v1.4.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/go-xorm/xorm v0.7.6
	github.com/pkg/errors v0.8.1
	github.com/stretchr/testify v1.4.0
	github.com/vektah/dataloaden v0.3.0 // indirect
	github.com/vektah/gqlparser v1.1.2
	github.com/zhanghup/go-tools v1.0.15
)

replace github.com/99designs/gqlgen => github.com/zhanghup/go-gqlgen v0.9.3

replace github.com/gin-gonic/gin => github.com/zhanghup/go-gin v0.0.0-20190818132815-0b17e4c65c1e

replace github.com/go-xorm/xorm => github.com/zhanghup/go-xorm v1.0.18
