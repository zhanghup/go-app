package gs

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"testing"
)

func TestName(t *testing.T) {
	a := tools.ObjectString()
	fmt.Println(tools.Password("zhang3611", *a))
}
