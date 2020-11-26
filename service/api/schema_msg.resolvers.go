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

func (r *mutationResolver) MessageConfirm(ctx context.Context, id string, input source.NewMessageConfirm) (bool, error) {
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
		go event.MsgNew(*r.Me(ctx).Info.User.Id, event.MsgTargetWeb, event.MsgActionConfirm, []beans.MsgInfo{*msg})
	}

	return err == nil, err
}

func (r *mutationResolver) MessageRead(ctx context.Context, id string) (bool, error) {
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
		go event.MsgNew(*r.Me(ctx).Info.User.Id, event.MsgTargetWeb, event.MsgActionRead, []beans.MsgInfo{*msg})
	}

	return err == nil, err
}

func (r *subscriptionResolver) Message(ctx context.Context) (<-chan *source.Message, error) {
	datas := make(chan *source.Message, 100)

	fn := func(msg event.MsgInfo) {
		action := source.MessageEnum(msg.Action)
		datas <- &source.Message{
			Target:   action,
			Messages: msg.Messages,
		}
	}

	event.MsgNewSubscribe(*r.Me(ctx).Info.User.Id, event.MsgTargetWeb, fn)
	go event.MsgNewUnSubscribeWithContext(ctx, *r.Me(ctx).Info.User.Id, event.MsgTargetWeb, fn)

	return datas, nil
}
