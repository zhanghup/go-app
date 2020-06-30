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

	_ = job.AddJob(e, "hello1", "* * * * * *", func() error {
		fmt.Println("1")
		for i := 1; i < 100000000; i++{}
		return nil
	})

	_ = job.AddJob(e, "hello2", "*/2 * * * * *", func() error {
		fmt.Println("2")
		for i := 1; i < 100000000; i++{}
		return nil
	})
	_ = job.AddJob(e, "hello3", "*/3 * * * * *", func() error {
		fmt.Println("3")
		for i := 1; i < 100000000; i++{}
		return nil
	})
	_ = job.AddJob(e, "hello4", "*/4 * * * * *", func() error {
		fmt.Println("4")
		for i := 1; i < 100000000; i++{}
		return nil
	})
	_ = job.AddJob(e, "hello5", "*/5 * * * * *", func() error {
		fmt.Println("5")
		for i := 1; i < 100000000; i++{}
		return nil
	})
	_ = job.AddJob(e, "hello6", "*/6 * * * * *", func() error {
		fmt.Println("6")
		for i := 1; i < 100000000; i++{}
		return nil
	})
	_ = job.AddJob(e, "hello7", "*/7 * * * * *", func() error {
		fmt.Println("7")
		for i := 1; i < 100000000; i++{}
		return nil
	})
	_ = job.AddJob(e, "hello8", "*/8 * * * * *", func() error {
		fmt.Println("8")
		for i := 1; i < 100000000; i++{}
		return nil
	})
	_ = job.AddJob(e, "hello9", "*/9 * * * * *", func() error {
		fmt.Println("9")
		for i := 1; i < 100000000; i++{}
		return nil
	})
	_ = job.AddJob(e, "hello0", "*/10 * * * * *", func() error {
		fmt.Println("0")
		for i := 1; i < 100000000; i++{}
		return nil
	})


	for {
		time.Sleep(time.Second)
	}
}
