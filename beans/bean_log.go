package beans

type CronLog struct {
	Bean `xorm:"extends"`

	Cron    *string `json:"cron" xorm:"index"`   // 作业id
	Start   *int64  `json:"start" `              // 开始时间
	End     *int64  `json:"end"`                 // 结束时间
	Message *string `json:"message" xorm:"text"` // 错误信息
}

type OperateLog struct {
	Bean `xorm:"extends"`

	Type   *string `json:"type"`   // 操作的表名
	Opt    *string `json:"opt"`    // 操作的动作 增/删/改/查
	Oid    *string `json:"oid"`    // 操作对象的id
	Uid    *string `json:"uid"`    // 操作的人员
	Uname  *string `json:"uname"`  // 操作员名称
	State  *int    `json:"state"`  // 操作的结果 0：失败，1：成功，-1：拒绝
	Msg    *string `json:"msg"`    // 异常信息
	Object *string `json:"object"` // 提交对象
}

func log_tables() []interface{} {
	return []interface{}{new(CronLog), new(OperateLog)}
}
