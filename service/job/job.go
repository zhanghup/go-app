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

func AddJob(name, spec string, action func() error, flag ...bool) error {

	id := tools.Crypto.MD5([]byte(name))
	if job.data.Get(id) == nil {
		// 纯粹的新任务，在数据库中没有任何记录
		model := beans.Cron{
			Bean: beans.Bean{
				Id:     &id,
				Status: tools.Ptr.Int(1),
			},
			Name:       &name,
			Expression: &spec,
			State:      tools.Ptr.Int(1),
		}
		_, err := job.db.Insert(&model)
		if err != nil {
			return err
		}
		f := false
		if len(flag) > 0 && flag[0] {
			f = true
		}

		entityId, err := job.job.AddFunc(spec, run(name, action))
		if err != nil {
			return err
		}

		if f {
			run(name, action)()
		}

		job.data.Set(id, jobitem{entityId, name, spec, action, f, true, model})
	} else {
		ji := job.data.Get(id).(jobitem)
		if ji.fn == nil {
			// 刚启动的任务
			ji.fn = action
			if ji.cron.State != nil && *ji.cron.State == 1 {
				entityId, err := job.job.AddFunc(ji.spec, run(ji.name, ji.fn))
				if err != nil {
					return err
				}
				ji.id = entityId
			}
			ji.isRun = false
		} else {
			job.job.Remove(ji.id)
			entityId, err := job.job.AddFunc(ji.spec, run(ji.name, ji.fn))
			if err != nil {
				return err
			}

			ji.id = entityId
			ji.isRun = true
		}

		job.data.Set(id, ji)
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

func Remove(id string) error {
	ji, err := checkJob(id)
	if err != nil {
		return err
	}
	if !ji.isRun {
		return nil
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
	return AddJob(ji.name, ji.spec, ji.fn, ji.flag)
}

func Restart(id string) error {
	err := Remove(id)
	if err != nil {
		return err
	}
	return Start(id)
}

func run(name string, fn func() error) func() {
	return func() {
		defer func() {
			if obj := recover(); obj != nil {
				tog.Error("job run panic ", obj)
			}
		}()

		l1 := time.Now().UnixNano()
		err := fn()
		l2 := time.Now().UnixNano()

		go tools.Run(func() {
			last := float64(l2-l1) / float64(time.Second)
			message := "ok"
			status := 1
			if err != nil {
				message = err.Error()
				status = 0
			}

			id := tools.Crypto.MD5([]byte(name))

			model := beans.Cron{}
			ok, err := job.db.Where("id = ?", id).Get(&model)
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

			_, err = job.db.Table(beans.Cron{}).Where("id = ?", id).Cols("last", "message", "previous", "updated", "status").Update(model)
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
			_, err = job.db.Insert(lg)
			if err != nil {
				tog.Error(err.Error())
				return
			}
		})
	}
}
