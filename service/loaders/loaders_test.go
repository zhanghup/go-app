package loaders

import (
	"fmt"
	"testing"
	"time"

	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-tools"
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
				err := load.Object(dict, "select * from dict where code in :keys", nil, "Code").Load(i, dict)
				//fmt.Println(tools.Str.JSONString(dict, true), i)
				if err != nil {
					panic(err)
				}
			}(k)
		}

		ids := []string{"5dcd5f3a3a391110a5002080", "5df0f6a33a3911288ebc244f", "5dcd5f3a3a391110a500207d", "5dcd5f3a3a391110a5002079", "5df0f6a33a3911288ebc2449", "5e8c9a371fa8d535801f0cfd"}
		for _, k := range ids {
			go func(i string) {
				dictItem := make([]beans.DictItem, 0)
				err := load.Slice(new(beans.DictItem), "select di.* from dict_item di where di.code in :keys", nil, "Code").Load(i, &dictItem)
				fmt.Println(tools.Str.JSONString(dictItem), i)
				if err != nil {
					panic(err)
				}
			}(k)
		}
		time.Sleep(time.Millisecond * 10)
	}

	time.Sleep(time.Second * 3)

}
