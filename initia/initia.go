package initia

import "github.com/go-xorm/xorm"

func InitAction(e *xorm.Engine) {
	initDict(e)
	// initMenu(e)
	initUser(e)
}
