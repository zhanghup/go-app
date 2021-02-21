package cfg

import (
	rice "github.com/GeertJohan/go.rice"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
	"github.com/zhanghup/go-tools/tgin"
	"github.com/zhanghup/go-tools/wx/wxmp"
)

type config struct {
	box      *rice.Box
	Host     string       `yaml:"host"`
	Database txorm.Config `yaml:"database"`
	Web      tgin.Config  `yaml:"web"`
	Wxmp     wxmp.Option  `yaml:"wxmp"`
}

var DB txorm.Config
var Web tgin.Config
var Wxmp wxmp.Option
var Config *config

func InitConfig(box *rice.Box) *config {
	if Config != nil {
		return Config
	}
	cc := new(config)
	cc.box = box
	err := tools.Conf(box, cc)
	if err != nil {
		panic(err)
	}
	Config = cc
	DB = cc.Database
	Web = cc.Web
	Wxmp = cc.Wxmp
	return Config
}

func ConfigOf(conf interface{}) error {
	if Config == nil {
		panic("配置文件尚未初始化")
	}
	err := tools.Conf(Config.box, conf)
	return err
}
