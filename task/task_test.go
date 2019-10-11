package task

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"testing"
	"time"
)

func TestRunningJob(t *testing.T) {
	e, err := xorm.NewEngine("mysql", "root:123@/test?charset=utf8")
	if err != nil {
		panic(err)
	}
	e.ShowSQL(true)
	//app.Sync(e)

	fmt.Println(NewCron(e).Add("", "", "*/5 * * * * ?", func() error {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		return nil
	}))

	select {}
}
