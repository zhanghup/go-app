module github.com/zhanghup/go-app

go 1.13

require (
	github.com/99designs/gqlgen v0.10.1
	github.com/daaku/go.zipexe v1.0.1 // indirect
	github.com/gin-gonic/gin v1.5.0
	github.com/giter/go.rice v0.0.0-20171227004756-39a3aa768429
	github.com/go-sql-driver/mysql v1.4.1
	github.com/go-xorm/xorm v0.7.6
	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0 // indirect
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646 // indirect 图片压缩
	github.com/pkg/errors v0.8.1
	github.com/rcrowley/go-metrics v0.0.0-20190706150252-9beb055b7962 // gin接口请求统计
	github.com/robfig/cron v1.2.0
	github.com/stretchr/testify v1.4.0
	github.com/vektah/dataloaden v0.3.0 // indirect
	github.com/vektah/gqlparser v1.1.2
	github.com/zhanghup/go-tools v1.0.18
	github.com/zhanghup/go-wxmp v0.0.3

	gopkg.in/ini.v1 v1.51.0
)

replace github.com/gin-gonic/gin => github.com/zhanghup/go-gin v1.50.3

replace github.com/go-xorm/xorm => github.com/zhanghup/go-xorm v1.0.22
