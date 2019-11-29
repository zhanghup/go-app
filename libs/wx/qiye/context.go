package qiye

import (
	"fmt"
	"github.com/zhanghup/go-app/cfg"
)

type Error struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func (this Error) Error() error {
	if this.Errcode == 0 {
		return nil
	}
	return fmt.Errorf("Error: %d, ErrorMessage: %s ", this.Errcode, this.Errmsg)
}

type Context struct {
	CorpId string

	application map[string]*application_api
}

var cc *Context

func NewContext() *Context {
	if cc != nil {
		return cc
	}
	if !cfg.WxQy().Enable {
		panic("config.ini - wxqy.enable 未启用")
	}
	return &Context{CorpId: cfg.WxQy().Corpid, application: map[string]*application_api{}}
}

func (this *Context) Application(agentid string) *application_api {
	a, ok := this.application[agentid]
	if ok {
		return a
	}
	a = newApplicationApi(agentid)
	this.application[agentid] = a
	return a
}
