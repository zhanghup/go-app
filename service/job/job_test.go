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
	job.InitJobs(e)

	_ = job.AddJob("hello1", "* * * * * *", func() error {
		fmt.Println("1")
		for i := 1; i < 100000000; i++ {
		}
		return nil
	})

	_ = job.AddJob("hello2", "*/2 * * * * *", func() error {
		fmt.Println("2")
		for i := 1; i < 100000000; i++ {
		}
		return nil
	})
	_ = job.AddJob("hello3", "*/3 * * * * *", func() error {
		fmt.Println("3")
		for i := 1; i < 100000000; i++ {
		}
		return nil
	})
	_ = job.AddJob("hello4", "*/4 * * * * *", func() error {
		fmt.Println("4")
		for i := 1; i < 100000000; i++ {
		}
		return nil
	})
	_ = job.AddJob("hello5", "*/5 * * * * *", func() error {
		fmt.Println("5")
		for i := 1; i < 100000000; i++ {
		}
		return nil
	})
	_ = job.AddJob("hello6", "*/6 * * * * *", func() error {
		fmt.Println("6")
		for i := 1; i < 100000000; i++ {
		}
		return nil
	})
	_ = job.AddJob("hello7", "*/7 * * * * *", func() error {
		fmt.Println("7")
		for i := 1; i < 100000000; i++ {
		}
		return nil
	})
	_ = job.AddJob("hello8", "*/8 * * * * *", func() error {
		fmt.Println("8")
		for i := 1; i < 100000000; i++ {
		}
		return nil
	})
	_ = job.AddJob("hello9", "*/9 * * * * *", func() error {
		fmt.Println("9")
		for i := 1; i < 100000000; i++ {
		}
		return nil
	})
	_ = job.AddJob("hello0", "*/10 * * * * *", func() error {
		fmt.Println("0")
		for i := 1; i < 100000000; i++ {
		}
		return nil
	})

	for i := 0; i < 20; i++ {
		time.Sleep(time.Second * 1)
		switch time.Now().Unix() % 10 {
		case 0:
			job.Remove("203ad5ffa1d7c650ad681fdff3965cd2")
		case 1:
			job.Remove("6e809cbda0732ac4845916a59016f954")
		case 2:
			job.Remove("7ce8be0fa3932e840f6a19c2b83e11ae")
		case 3:
			job.Remove("a75f2192bae11cb76cdcdada9332bab6")
		case 4:
			job.Remove("ebde9cc9540087b9688fdb470fa20f17")
		case 5:
			job.Remove("5726012822477f24fe999a1f7223c82a")
		case 6:
			job.Remove("2b4c43e3b1102b2e8492eebd97c06c58")
		case 7:
			job.Remove("f70109ba33bf1860001f89998581df2a")
		case 8:
			job.Remove("1c52d3d0503eca3c4a6db73af6d428f1")
		case 9:
			job.Remove("4f6d054536d6613a91472139cc60f072")
		}
	}

	fmt.Println("--------------------------------------------------------------------------------------------------------------------------")

	for i := 0; i < 20; i++ {
		time.Sleep(time.Second * 1)
		switch time.Now().Unix() % 10 {
		case 0:
			job.Restart("203ad5ffa1d7c650ad681fdff3965cd2")
		case 1:
			job.Restart("6e809cbda0732ac4845916a59016f954")
		case 2:
			job.Restart("7ce8be0fa3932e840f6a19c2b83e11ae")
		case 3:
			job.Restart("a75f2192bae11cb76cdcdada9332bab6")
		case 4:
			job.Restart("ebde9cc9540087b9688fdb470fa20f17")
		case 5:
			job.Restart("5726012822477f24fe999a1f7223c82a")
		case 6:
			job.Restart("2b4c43e3b1102b2e8492eebd97c06c58")
		case 7:
			job.Restart("f70109ba33bf1860001f89998581df2a")
		case 8:
			job.Restart("1c52d3d0503eca3c4a6db73af6d428f1")
		case 9:
			job.Restart("4f6d054536d6613a91472139cc60f072")
		}
	}
	for {
		time.Sleep(time.Second)
	}
}
