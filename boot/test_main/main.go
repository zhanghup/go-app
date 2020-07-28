package main

import (
	rice "github.com/giter/go.rice"
	"github.com/zhanghup/go-app/boot"
)

func main() {
	box, err := rice.FindBox("conf")
	if err != nil {
		panic(err)
	}
	_ = boot.Boot(box).
		EnableXorm().
		SyncTables().
		InitDatas().
		RouterFile().
		RouterAuth().
		RouterApi().
		StartRouter()
}
