package main

import (
	"github.com/gin-gonic/gin"
	"github.com/giter/go.rice"
	"github.com/zhanghup/go-app/boot"
	"github.com/zhanghup/go-app/ctx"
)

func main() {
	boot.Boot(func() (box *rice.Box, e error) {
		return rice.FindBox("conf")
	})
	ctx.Web().Engine().GET("test", func(context *gin.Context) {
		panic("----------------")
	})
	boot.Run()
}
