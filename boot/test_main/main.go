package main

import (
	rice "github.com/giter/go.rice"
	"github.com/pkg/errors"
	"github.com/zhanghup/go-app/boot"
	"time"
)

func main() {
	box, err := rice.FindBox("conf")
	if err != nil {
		panic(err)
	}
	_ = boot.Boot(box).
		SyncTables().
		InitDatas().
		JobsInit().
		Jobs("测试", "0 * * * * * ", func() error {
			if time.Now().Unix()%2 == 0 {
				return errors.New("测试错误")
			}
			return nil
		}).
		RouterFile().
		RouterAuth().
		RouterApi().
		StartRouter()
}
