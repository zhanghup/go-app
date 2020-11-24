package job

import (
	"github.com/pkg/errors"
	"github.com/robfig/cron/v3"
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
	"github.com/zhanghup/go-tools/tog"
	"time"
	"xorm.io/xorm"
)

type jobitem struct {
	id    cron.EntryID
	name  string
	spec  string
	fn    func() error // 为了可以手动停止，然后执行
	flag  bool
	isRun bool

	cron beans.Cron
}

var job *Jobs

type Jobs struct {
	job  *cron.Cron
	db   *xorm.Engine
	dbs  *txorm.Engine
	data tools.ICache // 存定时任务
}

func InitJobs(db *xorm.Engine) error {
	if job != nil {
		return nil
	}
	job = &Jobs{
		job:  cron.New(cron.WithSeconds()),
		db:   db,
		dbs:  txorm.NewEngine(db),
		data: tools.CacheCreate(),
	}
	l := make([]beans.Cron, 0)
	err := db.Find(&l)
	if err != nil {
		return err
	}
	for _, o := range l {
		if o.Id == nil {
			continue
		}
		ji := jobitem{cron: o}
		if o.Name != nil {
			ji.name = *o.Name
		}
		if o.Expression != nil {
			ji.spec = *o.Expression
		}
		job.data.Set(*o.Id, ji)
	}
	return nil
}

/*
	name	任务名称
	spec 	任务执行时间
	action 	执行方法
	flag 	是否立即执行
*/
func AddJob(name, spec string, action func() error, flag ...bool) error {

	id := tools.Crypto.MD5([]byte(name))
	if job.data.Get(id) == nil {
		// 纯粹的新任务，在数据库中没有任何记录
		model := beans.Cron{
			Bean: beans.Bean{
				Id:     &id,
				Status: tools.Ptr.String("1"),
			},
			Name:       &name,
			Expression: &spec,
			State:      tools.Ptr.String("start"),
		}
		_, err := job.db.Insert(&model)
		if err != nil {
			return err
		}
		f := false
		if len(flag) > 0 && flag[0] {
			f = true
		}

		job.data.Set(id, jobitem{0, name, spec, action, f, true, model})

		err = addJob(id)
		if err != nil {
			return err
		}

		if f {
			Run(id)()
		}
	} else {
		ji := job.data.Get(id).(jobitem)
		if ji.fn == nil {
			// 刚启动的任务
			ji.fn = action
			job.data.Set(id, ji)

			_, err := job.db.Table(beans.Cron{}).Where("id = ?", id).Update(map[string]interface{}{
				"expression": spec,
			})
			if err != nil {
				return err
			}

			if ji.cron.State != nil && *ji.cron.State == "1" {
				err := addJob(id)
				if err != nil {
					return err
				}
			}

			if ji.flag {
				Run(id)()
			}
		} else {
			err := addJob(id)
			if err != nil {
				return err
			}
		}

	}

	job.job.Start()
	return nil
}

func CheckJob(id string) error {
	_, err := checkJob(id)
	return err
}

func checkJob(id string) (jobitem, error) {
	fno := job.data.Get(id)
	if fno == nil {
		return jobitem{}, errors.New("不存在的任务id，已删除该任务")
	}
	ji := fno.(jobitem)

	return ji, nil
}

func Stop(id string) error {
	ji, err := checkJob(id)
	if err != nil {
		return err
	}
	if !ji.isRun {
		return nil
	}
	_, err = job.db.Table(beans.Cron{}).Where("id = ?", id).Update(map[string]interface{}{
		"state": 0,
	})
	if err != nil {
		return err
	}
	job.job.Remove(ji.id)
	ji.isRun = false
	job.data.Set(id, ji)
	return nil
}

func Start(id string) error {
	ji, err := checkJob(id)
	if err != nil {
		return err
	}
	if ji.isRun {
		return nil
	}
	_, err = job.db.Table(beans.Cron{}).Where("id = ?", id).Update(map[string]interface{}{
		"state": 1,
	})
	if err != nil {
		return err
	}
	return AddJob(ji.name, ji.spec, ji.fn, ji.flag)
}

func Restart(id string) error {
	err := Stop(id)
	if err != nil {
		return err
	}
	return Start(id)
}

func Run(id string) func() {
	ji := job.data.Get(id).(jobitem)

	return func() {
		defer func() {
			if obj := recover(); obj != nil {
				tog.Error("job run panic ", obj)
			}
		}()

		l1 := time.Now().UnixNano()
		err := ji.fn()
		l2 := time.Now().UnixNano()

		last := int64(float64(l2-l1) * 1000 / float64(time.Second))
		message := "ok"
		result := "1"
		if err != nil {
			message = err.Error()
			result = "0"
		}

		model := beans.Cron{}
		ok, err := job.db.Where("id = ?", id).Get(&model)
		if err != nil {
			tog.Error(err.Error())
			return
		}
		if !ok {
			return
		}

		model.Result = &result
		model.Last = &last
		model.Message = &message
		model.Previous = model.Updated

		_, err = job.db.Table(beans.Cron{}).Where("id = ?", id).Cols("last", "message", "previous", "updated", "status").Update(model)
		if err != nil {
			tog.Error(err.Error())
			return
		}
		lg := beans.CronLog{
			Bean: beans.Bean{
				Id:     tools.Ptr.Uid(),
				Status: tools.Ptr.String("1"),
			},
			Name:       &ji.name,
			Expression: &ji.spec,
			Cron:       &id,
			Message:    &message,
			Start:      tools.Ptr.Int64(l1 / int64(time.Second)),
			End:        tools.Ptr.Int64(l2 / int64(time.Second)),
		}
		_, err = job.db.Insert(lg)
		if err != nil {
			tog.Error(err.Error())
			return
		}
	}
}

func addJob(id string) error {
	ji := job.data.Get(id).(jobitem)
	job.job.Remove(ji.id)
	entityId, err := job.job.AddFunc(ji.spec, Run(id))
	if err != nil {
		return err
	}

	ji.id = entityId
	ji.isRun = true
	job.data.Set(id, ji)
	return nil
}
