package ctx_test

import (
	rice "github.com/giter/go.rice"
	"github.com/zhanghup/go-app/ctx"
	"testing"
)

func TestIni(t *testing.T) {
	box, err := rice.FindBox("conf")
	if err != nil {
		panic(err)
	}
	ctx.InitConfig(box)

}

func TestInitConf(t *testing.T) {
	ctx.InitConfFile()
}
