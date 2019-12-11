package mp

import "fmt"

func (this *context) Menu() Imenu {
	return &menu{
		context: this,
	}
}

type Imenu interface {
	Create(btns []Button) error
	Delete() error
}

type menu struct {
	context *context
}

func (this *menu) error(err interface{}, fn string, i ...int) error {
	var s = ""
	switch err.(type) {
	case string:
		s = err.(string)
	case error:
		s = err.(error).Error()
	case Error:
		e := err.(Error)
		if e.Errcode == 0 {
			return nil
		}
		s = fmt.Sprintf("%d: %s", e.Errcode, e.Errmsg)
	}

	if len(i) > 0 {
		return fmt.Errorf("微信公众号 - 自定义菜单_%d - %s - %s", i[0], fn, s)
	}
	return fmt.Errorf("微信公众号 - 自定义菜单 - %s - %s", fn, s)
}

type Button struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	Key   string `json:"key"`
	Url   string `json:"url"`
	Value string `json:"value"`

	SubButton []Button `json:"sub_button"`
}

func (this *menu) Create(btns []Button) error {
	result := Error{}
	err := this.context.post("https://api.weixin.qq.com/cgi-bin/menu/create?access_token=ACCESS_TOKEN", btns, &result)
	if err != nil {
		return this.error(err, "Create", 1)
	}
	return this.error(result, "Create", 2)
}

func (this *menu) Delete() error {
	result := Error{}
	err := this.context.get("https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=ACCESS_TOKEN", nil, &result)
	if err != nil {
		return this.error(err, "Delete", 1)
	}
	return this.error(result, "Delete", 2)
}

func (this *menu) Get() ([]Button, error) {
	data := struct {
		Error
		IsMenuOPen   int `json:"is_menu_o_pen"`
		SelfmenuInfo struct {
			Button []struct {
				Name      string `json:"name"`
				Type      string `json:"type"`
				Value     string `json:"value"`
				Key       string `json:"key"`
				Url       string `json:"url"`
				SubButton struct {
					List []struct {
						Name  string `json:"name"`
						Type  string `json:"type"`
						Key   string `json:"key"`
						Value string `json:"value"`
						Url   string `json:"url"`
					} `json:"list"`
				} `json:"sub_button"`
			} `json:"button"`
		} `json:"selfmenu_info"`
	}{}
	err := this.context.get("https://api.weixin.qq.com/cgi-bin/get_current_selfmenu_info?access_token=ACCESS_TOKEN", nil, &data)
	if err != nil {
		return nil, this.error(err, "Get", 1)
	}

	err = this.error(data.Error, "Get", 2)
	if err != nil {
		return nil, err
	}

	result := make([]Button, 0)
	for _, o := range data.SelfmenuInfo.Button {
		btn := Button{
			Name:      o.Name,
			Type:      o.Type,
			Value:     o.Value,
			Key:       o.Key,
			Url:       o.Url,
			SubButton: make([]Button, 0),
		}
		if o.SubButton.List != nil && len(o.SubButton.List) > 0 {
			for _, oo := range o.SubButton.List {
				btn.SubButton = append(btn.SubButton, Button{
					Name:  o.Name,
					Type:  o.Type,
					Value: o.Value,
					Key:   o.Key,
					Url:   o.Url,
				})
			}
		}
	}
	return result, nil

}
