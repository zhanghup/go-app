package job_test

import (
	"fmt"
	rice "github.com/giter/go.rice"
	"github.com/zhanghup/go-app/cfg"
	"github.com/zhanghup/go-app/service/job"
	"github.com/zhanghup/go-tools/database/txorm"
	"testing"
	"time"
)

func TestJob(t *testing.T) {
	box, err := rice.FindBox("conf")
	if err != nil {
		panic(err)
	}
	cfg.InitConfig(box)
	e, err := txorm.NewXorm(cfg.DB)
	if err != nil {
		panic(err)
	}

	_ = job.AddJob(e, "hello", "* * * * * *", func() error {
		fmt.Println("11111111111")
		for i := 1; i < 100000000; i++{}
		return nil
	})
	_ = job.AddJob(e, "hello2", "*/5 * * * * *", func() error {
		fmt.Println("222222222222")
		for i := 1; i < 100000000; i++{}
		return nil
	})

	for {
		time.Sleep(time.Second)
	}
}
