package beans

type WxqyUser struct {
	Bean   `xorm:"extend"`
	Uid    *string `json:"uid"`    // user表中的id
	Corpid *string `json:"corpid"` // 企业微信的corpid

	Userid           *string `json:"userid"`            // 成员UserID。对应管理端的帐号，企业内必须唯一。不区分大小写，长度为1~64个字节。只能由数字、字母和“_-@.”四种字符组成，且第一个字符必须是数字或字母。
	Name             *string `json:"name"`              // 成员名称。长度为1~64个utf8字符
	Alias            *string `json:"alias"`             // 成员别名。长度1~32个utf8字符
	Mobile           *string `json:"mobile"`            // 手机号码。企业内必须唯一，mobile/email二者不能同时为空
	Position         *string `json:"position"`          // 职务信息。长度为0~128个字符
	Gender           *string `json:"gender"`            // 性别。1表示男性，2表示女性
	Email            *string `json:"email"`             // 邮箱。长度6~64个字节，且为有效的email格式。企业内必须唯一，mobile/email二者不能同时为空
	Telephone        *string `json:"telephone"`         // 座机。32字节以内，由纯数字或’-‘号组成。
	Enable           *int    `json:"enable"`            // 启用/禁用成员。1表示启用成员，0表示禁用成员
	AvatarMediaid    *string `json:"avatar_mediaid"`    // 成员头像的mediaid，通过素材管理接口上传图片获得的mediaid
	Address          *string `json:"address"`           // 地址。长度最大128个字符
	ExternalPosition *string `json:"external_position"` // 对外职务，如果设置了该值，则以此作为对外展示的职务，否则以position来展示。长度12个汉字内
}

type WxqyButton struct {
	Bean    `xorm:"extend"`
	Corpid  *string `json:"corpid"`  // 企业微信的corpid
	Agentid *string `json:"agentid"` // 企业应用的id，整型。可在应用的设置页面查看

	Type   string `json:"type"`   // 菜单的响应动作类型 - WXQY001
	Name   string `json:"name"`   // 菜单的名字。不能为空，主菜单不能超过16字节，子菜单不能超过40字节。
	Key    string `json:"key"`    // 菜单KEY值，用于消息接口推送，不超过128字节
	Url    string `json:"url"`    // 网页链接，成员点击菜单可打开链接，不超过1024字节。为了提高安全性，建议使用https的url
	Parent string `json:"parent"` // 父菜单id
}
