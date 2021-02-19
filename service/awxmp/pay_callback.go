package awxmp

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhanghup/go-app/service/event"
	"github.com/zhanghup/go-tools/tog"
	"io"
)

func PayCallback(c *gin.Context) {
	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		tog.Error("【微信支付】 %s", err.Error())
		return
	}
	go event.WxmpPayCallbackPush(data)
	fmt.Println(string(data), "------------------")
}
