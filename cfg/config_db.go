package cfg

import "github.com/go-xorm/xorm"

type configDB struct {
	Enable  bool         `ini:"enable"`
	Type    string       `ini:"type"`
	Uri     string       `ini:"uri"`
	ShowSql bool         `ini:"show_sql"`
	engine  *xorm.Engine `ini:"-"`
}

func (this *configDB) Engine() *xorm.Engine {
	if this.engine != nil {
		return this.engine
	}
	var err error
	this.engine, err = xorm.NewEngine(this.Type, this.Uri)
	if err != nil {
		panic(err)
	}
	if this.ShowSql {
		this.engine.ShowSQL(true)
	}
	return this.engine
}

// 关系型数据库配置
func db(flag ...bool) *configDB {
	if my.DB == nil {
		panic("config.ini - [database] - 配置文件数据库信息尚未初始化完成")
	}
	if (len(flag) == 0 || flag[0]) && !my.DB.Enable {
		panic("config.ini - [database].enable 未启用")
	}
	return my.DB
}
func DB() *configDB {
	return db()
}
func DBEnable() bool {
	return db(false).Enable
}
