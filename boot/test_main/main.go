package main

import (
	"fmt"
	rice "github.com/giter/go.rice"
	"github.com/zhanghup/go-app/boot"
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
		Jobs("测试", "* * * * * * ", func() error {
			fmt.Println("22222222222222222222")
			return nil
		}).
		RouterFile().
		RouterAuth().
		RouterApi().
		StartRouter()
}
