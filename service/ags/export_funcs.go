package ags

import (
	"github.com/gin-gonic/gin"
	"github.com/zhanghup/go-app/beans"
	"io"
	"os"
)

// ----- 消息

// 消息发送
func MessageSend(tpl beans.MsgTemplate, uid, uname, otype, oid, defaultContent string, model map[string]string) error {
	return defaultMessage.NewMessage(tpl, uid, uname, otype, oid, defaultContent, model)
}

// ----- 上传方法

func UploaderUploadIO(read io.Reader, name, contentType string) (string, error) {
	return defaultUploader.UploadIO(read, name, contentType)
}
func UploaderGetFile(id string) (*beans.Resource, *os.File, error) {
	return defaultUploader.GetFile(id)
}
func UploaderUploadWithGin() func(c *gin.Context) {
	return defaultUploader.GinUpload()
}
func UploaderGetWithGin() func(c *gin.Context) {
	return defaultUploader.GinGet()
}
func UploaderResizeWithGin() func(c *gin.Context) {
	return defaultUploader.GinResize()
}
func UploaderAllWithGin(auth gin.IRouter, any gin.IRouter) {
	defaultUploader.GinRouter(auth, any)
}
