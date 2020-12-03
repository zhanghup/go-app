package beans

// 计划
type Plan struct {
	Bean   `xorm:"extends"`
	Name   *string `json:"name"`   // 计划名称
	Puid   *string `json:"puid"`   // 发起人
	Puname *string `json:"puname"` // 发起人名称
	Pstime *int64  `json:"pstime"` // 计划开始时间
	Petime *int64  `json:"petime"` // 计划结束时间
}

// 计划节点
type PlanStep struct {
	Bean  `xorm:"extends"`
	Plan  *string `json:"plan"`  // 计划id
	Pid   *string `json:"pid"`   // 上一个节点id => 生成一颗节点树
	Name  *string `json:"name"`  // 节点名称
	Stime *int64  `json:"stime"` // 节点开始时间
	Etime *int64  `json:"etime"` // 节点结束时间
}

// 计划节点
type PlanRole struct {
	Bean `xorm:"extends"`
	Plan *string `json:"plan"` // 计划id
	Step *string `json:"step"` // 节点id
	Role *string `json:"role"` // 角色
}

// 计划节点表单
type PlanForm struct {
	Bean     `xorm:"extends"`
	Plan     *string `json:"plan"`     // 计划id
	Step     *string `json:"step"`     // 节点id
	Type     *string `json:"role"`     // 属性类型 dict:PLA001
	Default  *string `json:"default"`  // 默认值
	Required *string `json:"required"` // 是否必填 dict: STA006
	Tfield   *string `json:"tfield"`   // 回填到计划的表单中的某个字段
}

// 任务
type Assign struct {
	Bean   `xorm:"extends"`
	Plan   *string `json:"plan"`   // 计划id
	Pname  *string `json:"pname"`  // 计划名称
	Puid   *string `json:"puid"`   // 发起人
	Puname *string `json:"puname"` // 发起人名称
	Pstime *int64  `json:"pstime"` // 计划开始时间
	Petime *int64  `json:"petime"` // 计划结束时间
	Step   *string `json:"step"`   // 当前节点id
	Uid    *string `json:"uid"`    // 当前提交人
	Commit *int64  `json:"commit"` // 提交时间
}

// 任务节点
type AssignStep struct {
	Bean   `xorm:"extends"`
	Plan   *string `json:"plan"`   // 计划id
	Step   *string `json:"step"`   // 节点id
	Pid    *string `json:"pid"`    // 上一个节点的任务
	Uid    *string `json:"uid"`    // 当前提交人
	Commit *int64  `json:"commit"` // 提交时间
}

// 任务表单
type AssignForm struct {
	Bean   `xorm:"extends"`
	Plan   *string `json:"plan"`   // 计划id
	Assign *string `json:"assign"` // 任务id
	Fid    *string `json:"fid"`    // 表单字段id
	Value  *string `json:"value"`  // 字段值
}
