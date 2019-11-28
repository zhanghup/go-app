package main

import (
	"github.com/giter/go.rice"
	"github.com/zhanghup/go-app/boot"
)

func main() {
	box, err := rice.FindBox("conf")
	if err != nil {
		panic(err)
	}
	boot.Boot(box)
	boot.Run()
}
