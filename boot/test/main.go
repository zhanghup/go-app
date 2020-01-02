package main

import (
	"github.com/gin-gonic/gin"
	"github.com/giter/go.rice"
	"github.com/zhanghup/go-app/boot"
)

func main() {
	boot.Boot(func() (box *rice.Box, e error) {
		return rice.FindBox("conf")
	})
	boot.Gin().GET("test", func(context *gin.Context) {
		panic("----------------")
	})
	boot.Run()
}
