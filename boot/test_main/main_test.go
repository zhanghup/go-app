package main

import (
	rice "github.com/giter/go.rice"
	"github.com/zhanghup/go-app/boot"
	"github.com/zhanghup/go-tools/tog"
	"testing"
	"time"
)

func TestMyLogger(t *testing.T) {
	for i := 0; i < 1000000; i++ {
		tog.Info("ddddddddddddddddd dsjkdj")
		tog.Error("ddddddddddddddddd dsjkdj")
		tog.Warn("ddddddddddddddddd dsjkdj")

		tog.Warn("ddddddddddddddddd dsjkdj")

		tog.Error("ddddddddddddddddd dsjkdj")

		tog.InfoAsJson(map[string]interface{}{"a": 1, "b": 2})
		tog.InfoAsJson(map[string]interface{}{"a": 1, "b": 2}, true)
		time.Sleep(time.Millisecond)
	}
}

func TestRand(t *testing.T) {
	box, err := rice.FindBox("conf")
	if err != nil {
		panic(err)
	}
	_ = boot.Boot(box).SyncTables().InitTestData()
}
