package beans

// 消息 --> 部门
// 消息 --> 群组
// 消息 --> 用户

// 消息模板
type MsgTemplate struct {
	Bean        `xorm:"extends"`
	Name        *string `json:"name"`
	Type        *string `json:"type"`         // 消息分类 - dict
	Level       *string `json:"level"`        // 消息等级 [严重、重要、次要、普通] - dict
	Target      *string `json:"target"`       // 推送目标，多个使用逗号分隔 [web,app,mini,sms...] - dict
	Timeout     *int64  `json:"timeout"`      // 消息超时时间
	MustConfirm *string `json:"must_confirm"` // 弹出消息是否必须确认 - dict
	Alert       *string `json:"alert"`        // 是否弹出 - dict
}

// 模板 - 群组
type MsgTemplateGroup struct {
	Bean     `xorm:"extends"`
	Group    *string `json:"group"`
	Template *string `json:"template"`
}

// 消息群组
type MsgGroup struct {
	Bean `xorm:"extends"`
	Name *string `json:"name"` // 群组名称
	Uid  *string `json:"uid"`  // 创建者
}

// 群组-用户
type MsgGroupUser struct {
	Bean  `xorm:"extends"`
	Group *string `json:"group"`
	Uid   *string `json:"uid"`
}

// 消息事件
type MsgEvent struct {
	Bean          `xorm:"extends"`
	Receiver      *string `json:"receiver" xorm:"index"` // 消息接收者
	ReceiverName  *string `json:"receiver_name"`         // 接收者名称
	Template      *string `json:"template"`              // 消息模板
	Type          *string `json:"type"`                  // 消息分类 - dict
	Level         *string `json:"level"`                 // 消息等级 [严重、重要、次要、普通] - dict
	Target        *string `json:"target"`                // 推送目标，多个使用逗号分隔 [web,app,mini,sms...] - dict
	Timeout       *int64  `json:"timeout"`               // 消息超时时间
	Alert         *string `json:"alert"`                 // 是否弹出 - dict
	MustConfirm   *string `json:"must_confirm"`          // 弹出消息是否必须确认 - dict
	ConfirmTarget *string `json:"confirm_target"`        // 确认平台 [web,app,mini,sms...] - dict
	State         *string `json:"is_read"`               // 消息状态 [已读、未读、未读并且已过期、已读并且已过期、已处理] - dict
	IsSend        *string `json:"is_send"`               // 消息发送状态 [已发送、未发送]
	Otype         *string `json:"otype"`                 // 消息对象
	Oid           *string `json:"oid"`                   // 消息对象id
	Title         *string `json:"title"`                 // 消息标题
	Content       *string `json:"msg_content"`           // 消息体
	ImgPath       *string `json:"img_path"`              // 消息提示图片
}

// 消息体
type MsgInfo struct {
	Bean          `xorm:"extends"`
	Event         *string `json:"event"`                 // 事件id
	Receiver      *string `json:"receiver" xorm:"index"` // 消息接收者
	ReceiverName  *string `json:"receiver_name"`         // 接收者名称
	Template      *string `json:"template"`              // 消息模板
	Type          *string `json:"type"`                  // 消息分类 - dict
	Level         *string `json:"level"`                 // 消息等级 [严重、重要、次要、普通]
	Target        *string `json:"target"`                // 推送目标，多个使用逗号分隔 [web,app,mini,sms...] - dict
	Timeout       *int64  `json:"timeout"`               // 消息超时时间
	MustConfirm   *string `json:"must_confirm"`          // 弹出消息是否必须确认 - dict
	ConfirmTarget *string `json:"confirm_target"`        // 确认平台 [web,app,mini,sms...] - dict
	State         *string `json:"is_read"`               // 消息状态 [已读、未读、已过期、已处理]
	SendTime      *int64  `json:"send_time"`             // 消息发送时间
	Otype         *string `json:"otype"`                 // 消息对象
	Oid           *string `json:"oid"`                   // 消息对象id
	Title         *string `json:"title"`                 // 消息标题
	Content       *string `json:"msg_content"`           // 消息体
	ImgPath       *string `json:"img_path"`              // 消息提示图片
}

func msg_tables() []interface{} {
	return []interface{}{
		new(MsgTemplate),
		new(MsgGroup),
		new(MsgGroupUser),
		new(MsgEvent),
		new(MsgInfo),
		new(MsgTemplateGroup),
	}
}
