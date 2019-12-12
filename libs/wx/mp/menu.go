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
	Get() ([]Button, error)
}

type menu struct {
	context *context
}

func (this *menu) error(err interface{}, fn string) error {
	s := this.context.error(err)
	if len(s) == 0 {
		return nil
	}
	return fmt.Errorf("微信公众号 - 自定义菜单 - %s - %s", fn, s)
}

type Button struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	Key      string `json:"key"`
	Url      string `json:"url"`
	Value    string `json:"value"`
	Appid    string `json:"appid"`
	Pagepath string `json:"pagepath"`

	SubButton []Button `json:"sub_button"`
}

func (this *menu) Create(btns []Button) error {
	result := Error{}
	err := this.context.post("https://api.weixin.qq.com/cgi-bin/menu/create?access_token=ACCESS_TOKEN", map[string]interface{}{
		"button": btns,
	}, &result)
	if err != nil {
		return this.error(err, "Create_0")
	}
	return this.error(result, "Create_1")
}

func (this *menu) Delete() error {
	result := Error{}
	err := this.context.get("https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=ACCESS_TOKEN", nil, &result)
	if err != nil {
		return this.error(err, "Delete_0")
	}
	return this.error(result, "Delete_1")
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
				Appid     string `json:"appid"`
				Pagepath  string `json:"pagepath"`
				SubButton struct {
					List []struct {
						Name     string `json:"name"`
						Type     string `json:"type"`
						Key      string `json:"key"`
						Value    string `json:"value"`
						Url      string `json:"url"`
						Appid    string `json:"appid"`
						Pagepath string `json:"pagepath"`
					} `json:"list"`
				} `json:"sub_button"`
			} `json:"button"`
		} `json:"selfmenu_info"`
	}{}
	err := this.context.get("https://api.weixin.qq.com/cgi-bin/get_current_selfmenu_info?access_token=ACCESS_TOKEN", nil, &data)
	if err != nil {
		return nil, this.error(err, "Get_0")
	}

	err = this.error(data.Error, "Get_1")
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
			Appid:     o.Appid,
			Pagepath:  o.Pagepath,
			SubButton: make([]Button, 0),
		}
		if o.SubButton.List != nil && len(o.SubButton.List) > 0 {
			for _, oo := range o.SubButton.List {
				btn.SubButton = append(btn.SubButton, Button{
					Name:     oo.Name,
					Type:     oo.Type,
					Value:    oo.Value,
					Key:      oo.Key,
					Url:      oo.Url,
					Appid:    oo.Appid,
					Pagepath: oo.Pagepath,
				})
			}
		}
		result = append(result, btn)
	}
	return result, nil

}
