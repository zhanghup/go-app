package gs

import (
	rice "github.com/GeertJohan/go.rice"
	"github.com/zhanghup/go-app/service/event"
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

func init() {
	event.XormDefaultInitSubscribeOnce(func(db *xorm.Engine) {
		defaultDB = db
		defaultDBS = txorm.NewEngine(db)
		tog.Info("数据库初始化成功。。。")

		defaultUploader = &uploader{db}
		tog.Info("文件服务初始化成功。。。")
	})
}
