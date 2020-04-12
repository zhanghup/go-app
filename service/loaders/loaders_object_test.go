package loaders

import (
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-tools/database/toolxorm"
	"testing"
	"time"
)

func TestObj(t *testing.T) {
	e, err := toolxorm.NewXorm(toolxorm.Config{
		Driver: "mysql",
		Uri:    "root:123@/test",
		Debug:  true,
	})
	if err != nil {
		panic(err)
	}
	if err := e.Ping(); err != nil {
		panic(err)
	}
	load := NewDataLoaden(e)

	for j := 0; j < 2; j++ {
		keys := []string{"SYS0002", "SYS0001", "SYS0003", "STA0001", "SYS0004", "STA0002"}

		for _, k := range keys {
			go func(i string) {
				dict := new(beans.Dict)
				err := load.Object(dict, "select * from dict where code in :keys", nil, "Code").Load(i, dict)
				if err != nil {
					panic(err)
				}
			}(k)
		}

		ids := []string{"5e8c9a361fa8d53580284d09", "5e8c9a361fa8d535804aafe1", "5e8c9a371fa8d535801f0cfd", "5e8c9a371fa8d5358040db79", "5e8c9a371fa8d5358083d181", "5e8c9a381fa8d5358037929d"}
		for _, k := range ids {
			go func(i string) {
				dictItem := new(beans.DictItem)
				err := load.Object(dictItem, "select * from dict_item where code in :keys", nil, "Code").Load(i, dictItem)
				if err != nil {
					panic(err)
				}
			}(k)
		}
		time.Sleep(time.Millisecond * 1000)
	}

	time.Sleep(time.Second * 3)

}
