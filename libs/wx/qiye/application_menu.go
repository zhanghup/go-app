package qiye

import (
	"github.com/zhanghup/go-tools"
)

type Button struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Key  string `json:"key"`
	Url  string `json:"url"`

	SubButton []Button `json:"sub_button"`
}

func (this *application_api) Menus() ([]Button, error) {
	token, err := this.access_token()
	if err != nil {
		return nil, err
	}
	data := struct {
		Button []Button `json:"button"`
		Error
	}{}
	err = tools.Http().GetI("https://qyapi.weixin.qq.com/cgi-bin/menu/get?access_token={{.access_token}}&agentid={{.agentid}}", map[string]interface{}{
		"access_token": token,
		"agentid":      this.agentid,
	}, &data)
	return data.Button, err
}

func (this *application_api) MenuCreate(btn []Button) error {
	token, err := this.access_token()
	if err != nil {
		return err
	}

	str, err := tools.Str().Template("https://qyapi.weixin.qq.com/cgi-bin/menu/create?access_token={{.access_token}}&agentid={{.agentid}}", map[string]interface{}{
		"access_token": token,
		"agentid":      this.agentid,
	}, nil)
	if err != nil {
		return err
	}
	result := Error{}
	err = tools.Http().PostI(str, map[string]interface{}{
		"button": btn,
	}, &result)
	if err != nil {
		return err
	}

	return result.Error()
}

func (this *application_api) MenuRemove() error {
	token, err := this.access_token()
	if err != nil {
		return err
	}
	result := Error{}
	err = tools.Http().GetI("https://qyapi.weixin.qq.com/cgi-bin/menu/delete?access_token={{.access_token}}&agentid={{.agentid}}", map[string]interface{}{
		"access_token": token,
		"agentid":      this.agentid,
	}, &result)
	if err != nil {
		return err
	}
	return result.Error()
}
