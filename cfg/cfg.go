package cfg

import (
	rice "github.com/giter/go.rice"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
	"github.com/zhanghup/go-tools/tgin"
)

type config struct {
	Database txorm.Config `yaml:"database"`
	Web      tgin.Config `yaml:"web"`
}

var DB txorm.Config
var Web tgin.Config
var Config *config

func InitConfig(box *rice.Box) *config {
	if Config != nil {
		return Config
	}
	cc := new(config)
	err := tools.Conf(box, cc)
	if err != nil {
		panic(err)
	}
	Config = cc
	DB = cc.Database
	Web = cc.Web
	return Config
}
