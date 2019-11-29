package qiye

import (
	"github.com/zhanghup/go-app/cfg"
)

type Context struct {
	CorpId string

	application map[string]*application
}

var cc *Context

func NewContext() *Context {
	if cc != nil {
		return cc
	}
	if !cfg.WxQy().Enable {
		panic("config.ini - wxqy.enable 未启用")
	}
	return &Context{CorpId: cfg.WxQy().Corpid, application: map[string]*application{}}
}

func (this *Context) Application(agentid string) *application {
	a, ok := this.application[agentid]
	if ok {
		return a
	}
	a = newApplication(agentid)
	this.application[agentid] = a
	return a
}
