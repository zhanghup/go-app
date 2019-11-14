package task

import (
	"context"
	"github.com/go-xorm/xorm"
	"github.com/robfig/cron"
	"github.com/zhanghup/go-app"
	"github.com/zhanghup/go-tools"
	"os"
	"runtime"
	"time"
)

type ICron interface {
	Add(code, name, expression string, fn func() error) error
}

type Cron struct {
	db   *xorm.Engine
	cron *cron.Cron
}

func NewCron(db *xorm.Engine) ICron {
	c := new(Cron)
	c.db = db
	c.cron = cron.New()

	c.cron.Start()
	return c
}

func (this *Cron) Add(code, name, expression string, fn func() error) error {
	cr := new(app.Cron)
	ok, err := this.db.SF(`select * from {{ table "cron" }} where code = :code`, map[string]interface{}{
		"code": code,
	}).Get(cr)
	if err != nil {
		return err
	}
	cr.Message = nil
	ctx := context.Background()
	ctx, err = this.db.Ts(ctx, func(s *xorm.Session) error {
		if !ok {
			cr.Bean = app.Bean{
				Id:     tools.ObjectString(),
				Status: tools.Ptr().Int(1),
			}
			cr.Code = &code
			cr.Name = &name
			cr.Expression = &expression
			_, err := s.Table(cr).Insert(cr)
			if err != nil {
				return err
			}
		} else {
			cr.Name = &name
			cr.Expression = &expression
			_, err := s.Table(cr).Where("id = ?", cr.Id).Update(cr)
			if err != nil {
				return err
			}
		}
		return nil
	})

	err = this.cron.AddFunc(expression, func() {
		lg := app.CronLog{
			Bean: app.Bean{
				Id:     tools.ObjectString(),
				Status: tools.Ptr().Int(1),
			},
			Cron:  cr.Id,
			Start: tools.Ptr().Int64(time.Now().Unix()),
		}
		_, err := this.db.Insert(lg)
		if err != nil {
			_, _ = os.Stdout.Write([]byte(err.Error()))
			return
		}

		ret := func(str string) {
			lg.End = tools.Ptr().Int64(time.Now().Unix())
			last := *lg.End - *lg.Start
			cr.Previous = lg.Start
			cr.Last = &last
			cr.Message = &str
			_, _ = this.db.Table(cr).Where("id = ?", cr.Id).Update(cr)
			lg.Message = cr.Message
			_, _ = this.db.Table(lg).Where("id = ?", lg.Id).Update(lg)
		}
		defer func() {
			if r := recover(); r != nil {
				const size = 64 << 10
				buf := make([]byte, size)
				buf = buf[:runtime.Stack(buf, false)]
				ret(string(buf))
			}
		}()

		err = fn()
		if err != nil {
			ret(err.Error())
		} else {
			ret("")
		}

	})
	return err
}
