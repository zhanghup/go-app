package job

import (
	"github.com/robfig/cron/v3"
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/tog"
	"time"
	"xorm.io/xorm"
)

type jobitem struct {
	id   cron.EntryID
	spec string
	fn   func() error // 为了可以手动停止，然后执行
}

var job = cron.New(cron.WithSeconds())

func AddJob(db *xorm.Engine, name, spec string, action func() error, flag ...bool) error {
	_, err := job.AddFunc(spec, run(db, name, action))
	if err != nil {
		return err
	}

	if len(flag) > 0 && flag[0] {
		run(db, name, action)()
	}

	if db != nil {
		id := tools.Crypto.MD5([]byte(name))
		model := beans.Cron{
			Bean: beans.Bean{
				Id:     &id,
				Status: tools.Ptr.Int(1),
			},
			Name:       &name,
			Expression: &spec,
		}
		ok, err := db.Table(model).Where("id = ?", id).Exist()
		if err != nil {
			return err
		}
		if !ok {
			_, err = db.Insert(model)
			if err != nil {
				return err
			}
		}

	} else {
		tog.Warn("cron waring: sql db is not supply")
	}

	job.Start()
	return nil
}

func run(db *xorm.Engine, name string, fn func() error) func() {
	return func() {
		defer func() {
			if obj := recover(); obj != nil {
				tog.Error("job run panic ", obj)
			}
		}()

		l1 := time.Now().UnixNano()
		err := fn()
		l2 := time.Now().UnixNano()

		if db != nil {
			go func() {
				last := float64(l2-l1) / float64(time.Second)
				message := "ok"
				status := 1
				if err != nil {
					message = err.Error()
					status = 0
				}

				id := tools.Crypto.MD5([]byte(name))

				model := beans.Cron{}
				ok, err := db.Where("id = ?", id).Get(&model)
				if err != nil {
					tog.Error(err.Error())
					return
				}
				if !ok {
					return
				}

				model.Status = &status
				model.Last = &last
				model.Message = &message
				model.Previous = model.Updated

				_, err = db.Table(beans.Cron{}).Where("id = ?", id).Cols("last", "message", "previous", "updated", "status").Update(model)
				if err != nil {
					tog.Error(err.Error())
					return
				}
				lg := beans.CronLog{
					Bean: beans.Bean{
						Id:     tools.Ptr.Uid(),
						Status: &status,
					},
					Cron:    &id,
					Message: &message,
					Start:   tools.Ptr.Int64(l1 / int64(time.Second)),
					End:     tools.Ptr.Int64(l2 / int64(time.Second)),
				}
				_, err = db.Insert(lg)
				if err != nil {
					tog.Error(err.Error())
					return
				}
			}()
		}
	}
}
