package ags

import (
	"github.com/zhanghup/go-tools/database/txorm"
	"xorm.io/xorm"
)

func NewUploader(db *xorm.Engine) IUploader {
	return &uploader{db: db}
}

func NewMessage(db *xorm.Engine) IMessage {
	return &message{
		db:  db,
		dbs: txorm.NewEngine(db),
	}
}
