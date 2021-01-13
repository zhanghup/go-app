package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/service/api/source"
	"github.com/zhanghup/go-tools"
)

func (r *mutationResolver) MyMsgInfoConfirm(ctx context.Context, id string, input source.NewMsgConfirm) (bool, error) {
	msg, err := r.MsgInfoLoader(ctx, id)
	if err != nil {
		return false, err
	}
	if msg == nil {
		return false, errors.New("消息id不存在,id:" + id)
	}

	err = r.Sess(ctx).SF(`
		update 
			msg_info 
		set 
			confirm_remark = :confirm_remark,
			confirm_time =  unix_timestamp(now()), 
			confirm_target = 'web',
			state = '4'
		where id = :id
	`,
		map[string]interface{}{
			"confirm_remark": input.Remark,
			"id":             id,
		}).Exec()

	return err == nil, err
}

func (r *mutationResolver) MyMsgInfoRead(ctx context.Context, id string) (bool, error) {
	msg, err := r.MsgInfoLoader(ctx, id)
	if err != nil {
		return false, err
	}
	if msg == nil {
		return false, errors.New("消息id不存在,id:" + id)
	}

	err = r.Sess(ctx).SF(`
		update 
			msg_info 
		set 
			read_time = unix_timestamp(now()), 
			read_target = 'web',
			state = '0'
		where id = :id
	`, map[string]interface{}{
		"id": id,
	}).Exec()

	return err == nil, err
}

func (r *myInfoResolver) ODept(ctx context.Context, obj *beans.User) (*beans.Dept, error) {
	if obj.Dept == nil {
		return nil, nil
	}
	return r.DeptLoader(ctx, *obj.Dept)
}

func (r *myInfoResolver) PermObjects(ctx context.Context, obj *beans.User) (interface{}, error) {
	return r.Me(ctx).Info.PermObjects, nil
}

func (r *queryResolver) MyInfo(ctx context.Context) (*beans.User, error) {
	user := r.Me(ctx).Info.User
	return &user, nil
}

func (r *queryResolver) MyMsgInfos(ctx context.Context, query source.QMyMsgInfo) ([]beans.MsgInfo, error) {
	return r.Query().MsgInfos(ctx, source.QMsgInfo{
		Receiver:      r.Me(ctx).Info.User.Id,
		Type:          query.Type,
		Level:         query.Level,
		Target:        tools.Ptr.String("web"),
		ConfirmTarget: query.ConfirmTarget,
		ReadTarget:    query.ReadTarget,
		State:         query.State,
		Index:         query.Index,
		Size:          query.Size,
		Status:        query.Status,
	})
}

// MyInfo returns source.MyInfoResolver implementation.
func (r *Resolver) MyInfo() source.MyInfoResolver { return &myInfoResolver{r} }

type myInfoResolver struct{ *Resolver }
