package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 静态文件的访问
func ResourcePublic(public *gin.RouterGroup) {
	public.StaticFS("/static", http.Dir("resource/static/"))
	public.StaticFile("/favicon.ico", "resource/favicon.ico")

}
