package system

import (
	"github.com/zhanghup/go-app/test"
	"github.com/zhanghup/go-tools"
	"testing"
)

func TestUserTemplate(t *testing.T) {
	test.MainTest(t, "user", map[string]interface{}{
		"type":     "a",
		"account":  tools.ObjectString(),
		"password": "Aa123456",
		"name":     "张超",
		"avatar":   "www.baidu.com",
		"i_card":   "330411199206123611",
		"birth":    1567252216,
		"sex":      1,
		"mobile":   "15115151515",
		"admin":    1,
		"weight":   0,
		"status":   1,
	})
}
