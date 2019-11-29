package cfg

import (
	"github.com/gin-gonic/gin"
	"log"
	"runtime/debug"
)

type configWeb struct {
	Enable    bool        `ini:"enable"`
	Port      string      `ini:"port"`
	engine    *gin.Engine `ini:"-"`
	engineRun bool        `ini:"-"`
}

func (this *configWeb) Engine() *gin.Engine {
	if this.engine != nil {
		return this.engine
	}
	this.engine = gin.Default()
	this.engine.Use(func(c *gin.Context) {
		defer func() {
			if p := recover(); p != nil {
				log.Println(string(debug.Stack()))
			}
		}()
		c.Next()
	})
	return this.engine
}

func (this *configWeb) Run() error {
	if this.engineRun {
		return nil
	}
	this.engineRun = true
	return this.engine.Run(":" + this.Port)
}

// api服务配置
func web(flag ...bool) *configWeb {
	if my.DB == nil {
		panic("config.ini - [web] - 配置文件web信息尚未初始化完成")
	}
	if (len(flag) == 0 || flag[0]) && !my.Web.Enable {
		panic("config.ini - [web].enable 未启用")
	}
	return my.Web
}
func Web() *configWeb {
	return web()
}

func WebEnable() bool {
	return web(false).Enable
}
