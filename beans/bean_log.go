package beans



type CronLog struct {
	Bean `xorm:"extends"`

	Cron    *string `json:"cron" xorm:"index"` // 作业id
	Start   *int64  `json:"start" `            // 开始时间
	End     *int64  `json:"end"`               // 结束时间
	Message *string `json:"message" xorm:"text"`           // 错误信息
}

func log_tables() []interface{} {
	return []interface{}{new(CronLog)}
}

