package cfg_test

import (
	rice "github.com/giter/go.rice"
	"github.com/zhanghup/go-app/cfg"
	"testing"
)

func TestIni(t *testing.T) {
	box, err := rice.FindBox("conf")
	if err != nil {
		panic(err)
	}
	cfg.InitConfig(box)

}

func TestInitConf(t *testing.T) {
	cfg.InitConfFile()
}
