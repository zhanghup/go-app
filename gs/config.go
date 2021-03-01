package gs

import (
	"context"
	rice "github.com/GeertJohan/go.rice"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
	"github.com/zhanghup/go-tools/tgin"
	"github.com/zhanghup/go-tools/tog"
	"github.com/zhanghup/go-tools/wx/wxmp"
	"xorm.io/xorm"
)

const (
	GIN_CONTEXT       = "gin-context"
	GIN_TOKEN         = "gin-token"
	GIN_AUTHORIZATION = "Authorization"
	GIN_USER          = "user_info"
	GIN_WXUSER        = "wxuser_info"
)

var (
	Background = context.Background()
)

type config struct {
	box      *rice.Box
	Host     string       `yaml:"host"`
	Database txorm.Config `yaml:"database"`
	Web      tgin.Config  `yaml:"web"`
	Wxmp     wxmp.Option  `yaml:"wxmp"`
}

var Config *config

func ConfigInit(box *rice.Box) *config {
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
	return Config
}

func ConfigOf(conf interface{}) error {
	if Config == nil {
		panic("配置文件尚未初始化")
	}
	err := tools.Conf(Config.box, conf)
	return err
}

func InfoBegin(name string, items ...string) {
	if len(items) > 0 {
		str := "【%s】 "
		ss := []interface{}{name}
		for _, s := range items {
			str += `- "%s" `
			ss = append(ss, s)
		}
		str += "- 模块开始初始化......................."

		tog.Info(str, ss...)
	} else {
		tog.Info("【%s】 模块开始初始化.......................", name)
	}
}

func InfoSuccess(name string, items ...string) {
	if len(items) > 0 {
		str := "【%s】 "
		ss := []interface{}{name}
		for _, s := range items {
			str += `- "%s" `
			ss = append(ss, s)
		}
		str += "- 模块初始化.......................成功"
		tog.Info(str, ss...)
	} else {
		tog.Info("【%s】 模块初始化.......................成功", name)
	}
}
func InfoError(name string, items ...string) {
	if len(items) > 0 {
		str := "【%s】 "
		ss := []interface{}{name}
		for _, s := range items {
			str += `- "%s" `
			ss = append(ss, s)
		}
		str += "- 模块初始化.......................失败"
		tog.Info(str, ss...)
	} else {
		tog.Info("【%s】 模块初始化.......................失败", name)
	}
}

func Init(db *xorm.Engine) {
	InfoBegin("数据库")
	defaultDB = db
	defaultDBA = txorm.NewEngine(db)
	InfoSuccess("数据库")

	InfoBegin("文件服务")
	defaultUploader = &uploader{db}
	InfoSuccess("文件服务")
}
