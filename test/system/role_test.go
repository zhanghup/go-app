package system

import (
	"github.com/zhanghup/go-app/test"
	"testing"
)

func TestRoleTemplate(t *testing.T) {
	test.NewMt().MainTest(t, "role", map[string]interface{}{
		"name":   "test_role",
		"desc":   "测试角色",
		"weight": 0,
		"status": 1,
	})
}
