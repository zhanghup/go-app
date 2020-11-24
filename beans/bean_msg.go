package beans

// 消息 --> 部门
// 消息 --> 群组
// 消息 --> 用户

// 消息模板
type MsgTemplate struct {
	Bean        `xorm:"extends"`
	Name        *string `json:"name"`
	Type        *string `json:"type"`         // 消息分类 - dict
	Level       *string `json:"level"`        // 消息等级 [0:严重、1:重要、2:次要、3:普通]
	Target      *string `json:"target"`       // 推送目标，多个使用逗号分隔 [web,app,mini,sms...] - dict
	Expire      *int64  `json:"timeout"`      // 消息超时时间（秒）
	MustConfirm *string `json:"must_confirm"` // 弹出消息是否必须确认 - dict
	Alert       *string `json:"alert"`        // 是否弹出 - dict
	ImgPath     *string `json:"img_path"`     // 消息提示图片
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
	Bean         `xorm:"extends"`
	Receiver     *string `json:"receiver" xorm:"index"` // 消息接收者
	ReceiverName *string `json:"receiver_name"`         // 接收者名称
	Template     *string `json:"template"`              // 消息模板
	Type         *string `json:"type"`                  // 消息分类 - dict
	Level        *string `json:"level"`                 // 消息等级 [0:严重、1:重要、2:次要、3:普通]
	Target       *string `json:"target"`                // 推送目标，多个使用逗号分隔 [web,app,mini,sms...] - dict
	Timeout      *int64  `json:"timeout"`               // 消息超时时间
	Alert        *string `json:"alert"`                 // 是否弹出 - dict
	MustConfirm  *string `json:"must_confirm"`          // 弹出消息是否必须确认 - dict
	Otype        *string `json:"otype"`                 // 消息对象
	Oid          *string `json:"oid"`                   // 消息对象id
	Title        *string `json:"title"`                 // 消息标题
	Content      *string `json:"msg_content"`           // 消息体
	ImgPath      *string `json:"img_path"`              // 消息提示图片
}

// 消息体
type MsgInfo struct {
	Bean          `xorm:"extends"`
	Event         *string `json:"event"`                 // 事件id
	Receiver      *string `json:"receiver" xorm:"index"` // 消息接收者
	ReceiverName  *string `json:"receiver_name"`         // 接收者名称
	Template      *string `json:"template"`              // 消息模板
	Type          *string `json:"type"`                  // 消息分类 - dict
	Level         *string `json:"level"`                 // 消息等级 [0:严重、1:重要、2:次要、3:普通]
	Target        *string `json:"target"`                // 推送目标，多个使用逗号分隔 [web,app,mini,sms...] - dict
	Timeout       *int64  `json:"timeout"`               // 消息超时时间
	Alert         *string `json:"alert"`                 // 是否弹出 - dict
	MustConfirm   *string `json:"must_confirm"`          // 弹出消息是否必须确认 - dict
	ConfirmTarget *string `json:"confirm_target"`        // 确认平台 [web,app,mini,sms...] - dict
	ReadTarget    *string `json:"read_target"`           // 已读平台 [web,app,mini,sms...] - dict
	State         *string `json:"state"`                 // 消息状态 [0:已读、1:未读、2:已读过期、3:未读过期、4:已确认] dict
	SendTime      *int64  `json:"send_time"`             // 消息发送时间
	ReadTime      *int64  `json:"read_time"`             // 消息阅读时间
	ConfirmTime   *int64  `json:"confirm_time"`          // 消息确认时间
	Otype         *string `json:"otype"`                 // 消息对象
	Oid           *string `json:"oid"`                   // 消息对象id
	Title         *string `json:"title"`                 // 消息标题
	Content       *string `json:"content"`               // 消息体
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