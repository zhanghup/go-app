package beans

// 消息 --> 部门
// 消息 --> 群组
// 消息 --> 用户

// 消息模板
type MsgTemplate struct {
	Bean        `xorm:"extends"`
	Name        *string `json:"name"`
	Code        *string `json:"code"`         // 模板编码
	Type        *string `json:"type"`         // 消息分类 - dict
	Level       *string `json:"level"`        // 消息等级 [0:严重、1:重要、2:次要、3:普通]
	Target      *string `json:"target"`       // 推送目标，多个使用逗号分隔 [web,app,mini,sms...] - dict
	Expire      *int64  `json:"timeout"`      // 消息超时时间（秒）
	MustConfirm *string `json:"must_confirm"` // 消息是否必须确认 - dict
	ImgPath     *string `json:"img_path"`     // 消息提示图片
}

// 消息体
type MsgInfo struct {
	Bean          `xorm:"extends"`
	Receiver      *string `json:"receiver" xorm:"index"` // 消息接收者
	ReceiverName  *string `json:"receiver_name"`         // 接收者名称
	Template      *string `json:"template"`              // 消息模板
	Type          *string `json:"type"`                  // 消息分类 - dict
	Level         *string `json:"level"`                 // 消息等级 [0:严重、1:重要、2:次要、3:普通]
	Target        *string `json:"target"`                // 推送目标，多个使用逗号分隔 [web,app,mini,sms...] - dict
	Timeout       *int64  `json:"timeout"`               // 消息超时时间
	MustConfirm   *string `json:"must_confirm"`          // 弹出消息是否必须确认 - dict
	ConfirmTarget *string `json:"confirm_target"`        // 确认平台 [web,app,mini,sms...] - dict
	ReadTarget    *string `json:"read_target"`           // 已读平台 [web,app,mini,sms...] - dict
	State         *string `json:"state"`                 // 消息状态 [0:已读、1:未读、2:已读过期、3:未读过期、4:已确认] dict
	SendTime      *int64  `json:"send_time"`             // 消息发送时间
	ReadTime      *int64  `json:"read_time"`             // 消息阅读时间
	ConfirmTime   *int64  `json:"confirm_time"`          // 消息确认时间
	ConfirmRemark *string `json:"confirm_remark"`        // 消息确认备注
	Otype         *string `json:"otype"`                 // 消息对象
	Oid           *string `json:"oid"`                   // 消息对象id
	Title         *string `json:"title"`                 // 消息标题
	Content       *string `json:"content"`               // 消息体
	ImgPath       *string `json:"img_path"`              // 消息提示图片
}

func msg_tables() []interface{} {
	return []interface{}{
		new(MsgTemplate),
		new(MsgInfo),
	}
}
