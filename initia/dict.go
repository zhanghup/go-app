package initia

import (
	"github.com/go-xorm/xorm"
	"github.com/zhanghup/go-app"
	"github.com/zhanghup/go-tools"
)

func InitDictCode(e *xorm.Engine, code, name, remark string, weight int) {
	dict := app.Dict{}
	ok, err := e.Table(dict).Where("code = ?", code).Get(&dict)
	if err != nil {
		panic(err)
	}
	if ok {
		return
	}

	dict.Code = &code
	dict.Name = &name
	dict.Remark = &remark
	dict.Weight = &weight
	dict.Id = tools.ObjectString()
	dict.Status = tools.Ptr().Int(1)
	e.Table(dict).Insert(dict)
}

func InitDict(e *xorm.Engine) {
	InitDictCode(e, "D0001", "用户类型", "", 1)
}
