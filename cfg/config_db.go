package cfg

import "github.com/go-xorm/xorm"

type configDB struct {
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
