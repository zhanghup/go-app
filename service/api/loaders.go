package api

import (
	"context"
	"github.com/zhanghup/go-app/beans"
)

func (this *Resolver) DictLoader(ctx context.Context, id string) (*beans.Dict, error) {
	result := new(beans.Dict)
	err := this.Loader(ctx).Object(result, "select * from dict where id in :keys", nil, "Id", "").Load(id, result)
	return result, err
}

func (this *Resolver) DictItemLoader(ctx context.Context, id string) (*beans.DictItem, error) {
	result := new(beans.DictItem)
	err := this.Loader(ctx).Object(result, "select * from dict_item where id in :keys", nil, "Id", "").Load(id, result)
	return result, err
}

func (this *Resolver) RoleLoader(ctx context.Context, id string) (*beans.Role, error) {
	result := new(beans.Role)
	err := this.Loader(ctx).Object(result, "select * from role where id in :keys", nil, "Id", "").Load(id, result)
	return result, err
}
func (this *Resolver) DeptLoader(ctx context.Context, id string) (*beans.Dept, error) {
	result := new(beans.Dept)
	err := this.Loader(ctx).Object(result, "select * from dept where id in :keys", nil, "Id", "").Load(id, result)
	return result, err
}

func (this *Resolver) UserLoader(ctx context.Context, id string) (*beans.User, error) {
	result := new(beans.User)
	err := this.Loader(ctx).Object(result, "select * from user where id in :keys", nil, "Id", "").Load(id, result)
	return result, err
}

func (this *Resolver) CronLoader(ctx context.Context, id string) (*beans.Cron, error) {
	result := new(beans.Cron)
	err := this.Loader(ctx).Object(result, "select * from cron where id in :keys", nil, "Id", "").Load(id, result)
	return result, err
}

func (this *Resolver) MsgInfoLoader(ctx context.Context, id string) (*beans.MsgInfo, error) {
	result := new(beans.MsgInfo)
	err := this.Loader(ctx).Object(result, "select * from msg_info where id in :keys", nil, "Id", "").Load(id, result)
	return result, err
}

func (this *Resolver) MsgTemplateLoader(ctx context.Context, id string) (*beans.MsgTemplate, error) {
	result := new(beans.MsgTemplate)
	err := this.Loader(ctx).Object(result, "select * from msg_template where id in :keys", nil, "Id", "").Load(id, result)
	if result.Id == nil {
		return nil, err
	}
	return result, err
}

func (this *Resolver) AccountLoader(ctx context.Context, id string) (*beans.Account, error) {
	result := new(beans.Account)
	err := this.Loader(ctx).Object(result, "select * from account where id in :keys", nil, "Id", "").Load(id, result)
	if result.Id == nil {
		return nil, err
	}
	return result, err
}

func (this *Resolver) AccountDefaultLoader(ctx context.Context, uid string) (*beans.Account, error) {
	result := new(beans.Account)
	err := this.Loader(ctx).Object(result, "select * from account where uid in :keys and `default` = 1", nil, "Uid", "").Load(uid, result)
	if result.Id == nil {
		return nil, err
	}
	return result, err
}

func (this *Resolver) PlanLoader(ctx context.Context, id string) (*beans.Plan, error) {
	result := new(beans.Plan)
	err := this.Loader(ctx).Object(result, "select * from plan where id in :keys", nil, "Id", "").Load(id, result)
	if result.Id == nil {
		return nil, err
	}
	return result, err
}

func (this *Resolver) PlanStepLoader(ctx context.Context, id string) (*beans.PlanStep, error) {
	result := new(beans.PlanStep)
	err := this.Loader(ctx).Object(result, "select * from plan_step where id in :keys", nil, "Id", "").Load(id, result)
	if result.Id == nil {
		return nil, err
	}
	return result, err
}
