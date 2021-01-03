package beans

// 消息模板
type MsgTemplate struct {
	Bean         `xorm:"extends"`
	Name         *string `json:"name"`
	Code         *string `json:"code" xorm:"index"`         // 模板编码
	Type         *string `json:"type"`                      // 消息分类 - dict
	Level        *string `json:"level"`                     // 消息等级 [0:严重、1:重要、2:次要、3:普通]
	Target       *string `json:"target"`                    // 推送目标，多个使用逗号分隔 [web,app,mini,sms...] - dict
	ToAdmin      *string `json:"to_admin"`                  // 是否推送管理员用户
	Expire       *int64  `json:"expire"`                    // 消息超时时间（秒）
	Delay        *int64  `json:"delay"`                     // 消息延时
	Alert        *int64  `json:"alert"`                     // 消息提前提醒时间
	ImgPath      *string `json:"img_path"`                  // 消息提示图片
	Remark       *string `json:"remark"`                    // 备注
	Template     *string `json:"template" xorm:"text"`      // 模板
	TemplateCode *string `json:"template_code" xorm:"text"` // 模板
}

// 消息体
type MsgInfo struct {
	Bean          `xorm:"extends"`
	Receiver      *string `json:"receiver" xorm:"index"` // 消息接收者
	ReceiverName  *string `json:"receiver_name"`         // 接收者名称
	Template      *string `json:"template"`              // 消息模板
	Type          *string `json:"type" xorm:"index"`     // 消息分类 - dict
	Level         *string `json:"level"`                 // 消息等级 [0:严重、1:重要、2:次要、3:普通]
	Target        *string `json:"target" xorm:"target"`  // 推送目标，多个使用逗号分隔 [web,app,mini,sms...] - dict
	Timeout       *int64  `json:"timeout"`               // 消息超时时间
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
	Model         *string `json:"content_map"`           // json字符串，用于格式化content
	ImgPath       *string `json:"img_path"`              // 消息提示图片
	Remark        *string `json:"remark"`                // 备注
}

// 消息体
type MsgHistory struct {
	Bean          `xorm:"extends"`
	Info          *string `json:"info" xorm:"index"`
	Receiver      *string `json:"receiver"`             // 消息接收者
	ReceiverName  *string `json:"receiver_name"`        // 接收者名称
	Template      *string `json:"template"`             // 消息模板
	Type          *string `json:"type"`                 // 消息分类 - dict
	Level         *string `json:"level"`                // 消息等级 [0:严重、1:重要、2:次要、3:普通]
	Target        *string `json:"target" xorm:"target"` // 推送目标，多个使用逗号分隔 [web,app,mini,sms...] - dict
	Timeout       *int64  `json:"timeout"`              // 消息超时时间
	ConfirmTarget *string `json:"confirm_target"`       // 确认平台 [web,app,mini,sms...] - dict
	ReadTarget    *string `json:"read_target"`          // 已读平台 [web,app,mini,sms...] - dict
	State         *string `json:"state"`                // 消息状态 [0:已读、1:未读、2:已读过期、3:未读过期、4:已确认] dict
	SendTime      *int64  `json:"send_time"`            // 消息发送时间
	ReadTime      *int64  `json:"read_time"`            // 消息阅读时间
	ConfirmTime   *int64  `json:"confirm_time"`         // 消息确认时间
	ConfirmRemark *string `json:"confirm_remark"`       // 消息确认备注
	Otype         *string `json:"otype"`                // 消息对象
	Oid           *string `json:"oid"`                  // 消息对象id
	Title         *string `json:"title"`                // 消息标题
	Content       *string `json:"content"`              // 消息体
	Model         *string `json:"content_map"`          // json字符串，用于格式化content
	ImgPath       *string `json:"img_path"`             // 消息提示图片
	Remark        *string `json:"remark"`               // 备注
}

func msg_tables() []interface{} {
	return []interface{}{
		new(MsgTemplate),
		new(MsgInfo),
		new(MsgHistory),
	}
}
