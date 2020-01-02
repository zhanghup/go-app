package ctx

import (
	"github.com/giter/go.rice"
	"gopkg.in/ini.v1"
	"reflect"
	"regexp"
)

type config struct {
	Web     *configWeb               `ini:"web"`
	DB      *configDB                `ini:"database"`
	Wxqy    *configWxQy              `ini:"wxqy"`
	WxqyApp map[string]configWxQyApp `ini-map:"wxqy-app"`
	Wxmp    *configWxmp              `ini:"wxmp"`
	Wxmi    *configWxmi              `ini:"wxmi"`
}

var my = &config{
	WxqyApp: map[string]configWxQyApp{"": {}}, // 若是map类型的配置项数据，必须先初始化一条数据，不然反射不到
}

func InitConfig(box *rice.Box) {
	f, err := box.Open("config.ini")
	if err != nil {
		panic("config.ini - 配置文件文件打开失败")
	}
	sess, err := ini.Load(f)
	if err != nil {
		panic("config.ini - 初始化配置文件异常")
	}

	r, err := regexp.Compile(`(.*?)\s"(.*?)"`)
	if err != nil {
		panic("config.ini - 初始化配置文件异常")
	}
	{
		vl := reflect.ValueOf(my).Elem()
		ty := reflect.TypeOf(my).Elem()

		for i := 0; i < vl.NumField(); i++ {
			v := vl.Field(i)
			f := ty.Field(i)
			t := v.Type()
			if t.Kind() == reflect.Ptr {
				t = t.Elem()
			}

			if len(f.Tag.Get("ini")) != 0 {
				obj := reflect.New(t)
				sec := sess.Section(f.Tag.Get("ini"))
				err := sec.MapTo(obj.Interface())
				if err != nil {
					panic("config.ini - 配置文件数据注入异常【1】")
				}

				// 配置文件对象
				v.Set(obj)

			} else if len(f.Tag.Get("ini-map")) != 0 {
				obj := reflect.MakeMap(t)
				secs := sess.Sections()
				for _, o := range secs {
					cfs := r.FindStringSubmatch(o.Name())
					if len(cfs) != 3 {
						continue
					}
					key := cfs[1]
					value := cfs[2]

					if f.Tag.Get("ini-map") == key {
						// 配置文件对象
						rg := v.MapRange()
						if rg.Next() {
							oo := reflect.New(rg.Value().Type())
							err := o.MapTo(oo.Interface())
							if err != nil {
								panic("config.ini - 配置文件数据注入异常【3】")
							}

							obj.SetMapIndex(reflect.ValueOf(value), oo.Elem())
						}

					}
				}
				v.Set(obj)
			}
		}
	}
}
