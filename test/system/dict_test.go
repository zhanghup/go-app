package system

import (
	"fmt"
	"github.com/zhanghup/go-app/test"
	"testing"
	"time"
)

func TestDictTemplate(t *testing.T) {
	code := fmt.Sprintf("%v", time.Now().String())
	test.NewMt().MainTest(t, "dict", map[string]interface{}{
		"code":   code,
		"name":   "用户类型",
		"remark": "测试哦那个和用户类型",
		"weight": 0,
		"status": 1,
	}, map[string]interface{}{
		"name":   "用户类型2",
		"remark": "测试哦那个和用户类型2",
		"weight": 0,
		"status": 1,
	})

	test.NewMt(test.MtQuery, test.MtGet).MainTest(t, "dict_item", map[string]interface{}{
		"code":   code,
		"name":   "用户类型",
		"value":  "1",
		"weight": 0,
		"status": 1,
	}, map[string]interface{}{
		"name":   "用户类型2",
		"value":  "2",
		"weight": 0,
		"status": 1,
	})
}
