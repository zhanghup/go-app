package loaders

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"testing"
	"time"

	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-tools/database/toolxorm"
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

	for j := 0; j < 10; j++ {
		keys := []string{"SYS0002", "SYS0001", "SYS0003", "STA0001", "SYS0004", "STA0002"}

		for _, k := range keys {
			go func(i string) {
				dict := new(beans.Dict)
				err := load.Object(dict, "select * from dict where code in :keys", nil, "Code", "").Load(i, dict)
				fmt.Println(tools.Str.JSONString(dict, true), i)
				if err != nil {
					panic(err)
				}
			}(k)
		}

		for _, k := range keys {
			go func(i string) {
				dictItem := make([]beans.DictItem, 0)
				err := load.Slice(struct {
					beans.DictItem `xorm:"extends"`
					DictCode       string `json:"dict_code"`
				}{}, "select di.*,d.code dict_code from dict_item di join dict d on d.id = di.code where d.code in :keys", nil, "DictCode", "DictItem").Load(i, &dictItem)
				if err != nil {
					panic(err)
				}
			}(k)
		}
		time.Sleep(time.Millisecond * 10)
	}

	time.Sleep(time.Second * 3)

}
