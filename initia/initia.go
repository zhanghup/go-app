package initia

import "xorm.io/xorm"

func InitAction(db *xorm.Engine) {
	NewDict(db).InitDict()
	InitUser(db)
}
