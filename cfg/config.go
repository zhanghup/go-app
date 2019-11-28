package cfg

import (
	"github.com/giter/go.rice"
	"gopkg.in/ini.v1"
	"reflect"
)

type config struct {
	Web *configWeb `ini:"web"`
	DB  *configDB  `ini:"database"`
}

var my = new(config)

func DB() *configDB {
	if my.DB == nil {
		panic("【config.ini】 配置文件数据库信息尚未初始化完成")
	}
	return my.DB
}

func Web() *configWeb {
	if my.DB == nil {
		panic("【config.ini】 配置文件web信息尚未初始化完成")
	}
	return my.Web
}

func InitConfig(box *rice.Box) {
	f, err := box.Open("config.ini")
	if err != nil {
		panic("【config.ini】 config.ini 文件打开失败")
	}
	sess, err := ini.Load(f)
	if err != nil {
		panic("【config.ini】 初始化配置文件异常")
	}
	err = sess.ReflectFrom(my)
	vl := reflect.ValueOf(my).Elem()
	ty := reflect.TypeOf(my).Elem()
	for i := 0; i < vl.NumField(); i++ {
		v := vl.Field(i)
		f := ty.Field(i)
		t := v.Type()
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
		obj := reflect.New(t)
		sec := sess.Section(f.Tag.Get("ini"))
		err := sec.MapTo(obj.Interface())
		if err != nil {
			panic("【config.ini】 配置文件数据注入异常")
		}
		v.Set(obj)
	}

}
