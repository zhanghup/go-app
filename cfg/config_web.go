package cfg

import "github.com/gin-gonic/gin"

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
	return this.engine
}

func (this *configWeb) Run() error {
	if this.engineRun {
		return nil
	}
	this.engineRun = true
	return this.engine.Run(":" + this.Port)
}
