package beans

// 消息群组
type MsgGroup struct {
	Bean `xorm:"extends"`
	Name *string `json:"name"` // 群组名称
	Uid  *string `json:"uid"`  // 创建者
}

// 消息分类
type MsgType struct {
	Bean        `xorm:"extends"`
	Name        *string `json:"name"`
	Pid         *string `json:"pid"`
	Timeout     *int64  `json:"timeout"`      // 消息超时时间
	Alert       *int    `json:"alert"`        // 是否弹出
	MustConfirm *int    `json:"must_confirm"` // 弹出消息是否必须确认
}

// 群组-用户
type MsgGroupUser struct {
	Bean  `xorm:"extends"`
	Group *string `json:"group"`
	Uid   *string `json:"uid"`
}

// 群组-用户-消息
type Msg struct {
	Bean     `xorm:"extends"`
	Receiver *string `json:"receiver" xorm:"index"` // 消息接收者
	Type     *string `json:"msg_type_id"`           // 消息类型id
	Level    *string `json:"level"`                 // 消息等级 [严重、重要、次要、普通]
	State    *string `json:"is_read"`               // 消息状态 [已读、未读、过期、处理]
	Alert    *int    `json:"alert"`                 // 是否弹出
	Timeout  *int64  `json:"timeout"`               // 消息超时时间
	SendTime *int64  `json:"send_time"`             // 消息发送时间
	Otype    *string `json:"otype"`                 // 消息对象
	Oid      *string `json:"oid"`                   // 消息对象id
	Title    *string `json:"title"`                 // 消息标题
	Content  *string `json:"msg_content"`           // 消息体
	ImgPath  *string `json:"img_path"`              // 消息提示图片
}

func msg_tables() []interface{} {
	return []interface{}{
		new(MsgGroup),
		new(MsgType),
		new(MsgGroupUser),
		new(Msg),
	}
}
