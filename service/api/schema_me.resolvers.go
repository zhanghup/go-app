package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"time"

	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/service/api/source"
	"github.com/zhanghup/go-app/service/event"
)

func (r *mutationResolver) MyMsgInfoConfirm(ctx context.Context, id string, input source.NewMsgConfirm) (bool, error) {
	msg, err := r.MsgInfoLoader(ctx, id)
	if err != nil {
		return false, err
	}
	if msg == nil {
		return false, errors.New("消息id不存在,id:" + id)
	}

	_, err = r.DB.Table(beans.MsgInfo{}).Where("id = ?", id).Update(map[string]interface{}{
		"confirm_remark": input.Remark,
		"confirm_time":   time.Now().Unix(),
	})
	if err == nil {
		go event.MsgNew(*r.Me(ctx).Info.User.Id, event.MsgTargetWeb, event.MsgActionConfirm, *msg)
	}

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

	_, err = r.DB.Table(beans.MsgInfo{}).Where("id = ?", id).Update(map[string]interface{}{
		"read_time": time.Now().Unix(),
	})
	if err == nil {
		go event.MsgNew(*r.Me(ctx).Info.User.Id, event.MsgTargetWeb, event.MsgActionRead, *msg)
	}

	return err == nil, err
}

func (r *myInfoResolver) ODept(ctx context.Context, obj *beans.User) (*beans.Dept, error) {
	if obj.Dept == nil {
		return nil, nil
	}
	return r.DeptLoader(ctx, *obj.Dept)
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
		Target:        query.Target,
		MustConfirm:   query.MustConfirm,
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
