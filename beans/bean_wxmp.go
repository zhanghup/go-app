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

type WxmpOrder struct {
	Bean          `xorm:"extends"`
	Uid           *string `json:"uid" xorm:"index(uid)"`
	Openid        *string `json:"openid" xorm:"index(openid)"`
	Otype         *string `json:"otype" xorm:"index(type_id)"`
	Oid           *string `json:"oid" xorm:"index(type_id)"`
	State         *string `json:"state"`
	Commit        *int64  `json:"commit"`
	Price         *int    `json:"price"`
	PriceUser     *int    `json:"price_user"`
	PrepayId      *string `json:"prepay_id"`
	TransactionId *string `json:"transaction_id"`
	PayTime       *int64  `json:"pay_time"`
}

func wxmp_tables() []interface{} {
	return []interface{}{
		new(WxmpUser),
		new(WxmpOrder),
	}
}
