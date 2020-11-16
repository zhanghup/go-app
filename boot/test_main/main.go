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
		//SyncTables().
		//InitDatas().
		//JobsInit().
		//Jobs("测试", "0 * * * * * ", func() error {
		//	if time.Now().Unix()%2 == 0 {
		//		return errors.New("测试错误")
		//	}
		//	return nil
		//}).
		//Jobs("河东域名同步", "0 * * * * * ", func() error {
		//	_, items, ok := ca.DictCache.Get("AUT001")
		//	if !ok {
		//		return errors.New("无同步内容[0]")
		//	}
		//	if len(items) == 0 {
		//		return errors.New("无同步内容[1]")
		//	}
		//
		//	return nil
		//}).
		RouterFile().
		RouterAuth().
		RouterApi().
		StartRouter()
}
