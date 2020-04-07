package cfg

import (
	rice "github.com/giter/go.rice"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/toolxorm"
	"github.com/zhanghup/go-tools/toolgin"
)

type config struct {
	Database toolxorm.Config `yaml:"database"`
	Web      toolgin.Config `yaml:"web"`
}

var DB toolxorm.Config
var Web toolgin.Config
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
