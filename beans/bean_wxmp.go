package beans

type WxmpUser struct {
	Bean      `xorm:"extends"`
	Uid       *string `json:"uid"`
	Mobile    *string `json:"mobile"`
	Appid     *string `json:"appid"`
	Unionid   *string `json:"unionid"`
	Openid    *string `json:"openid"`
	Nickname  *string `json:"nickname"`
	AvatarUrl *string `json:"avatar_url"`
	Gender    *string `json:"gender"`
	Country   *string `json:"country"`
	Province  *string `json:"province"`
	City      *string `json:"city"`
	Language  *string `json:"language"`
}

func wxmp_tables() []interface{} {
	return []interface{}{
		new(WxmpUser),
	}
}
