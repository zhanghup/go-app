package beans

type WxmiUser struct {
	Bean `xorm:"extends"`

	Uid     *string `json:"uid"`
	Openid  *string `json:"openid"`
	Unionid *string `json:"unionid"`
	Appid   *string `json:"appid"`
}

func wxmi_tables() []interface{} {
	return []interface{}{new(WxmiUser)}
}
