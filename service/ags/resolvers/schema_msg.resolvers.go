package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/service/ags/source"
	"github.com/zhanghup/go-app/service/event"
)

func (r *mutationResolver) MessageConfirm(ctx context.Context, target source.MessageEnum, id string, input source.NewMessageConfirm) (bool, error) {
	msg, err := r.MsgInfoLoader(ctx, id)
	if err != nil {
		return false, err
	}
	if msg == nil {
		return false, errors.New("消息id不存在,id:" + id)
	}

	_, err = r.DB.Table(beans.MsgInfo{}).Where("id = ?", id).Update(map[string]interface{}{
		"confirm_remark": input.Remark,
	})
	if err == nil {
		go event.MsgNew(*r.Me(ctx).Info.User.Id, event.MsgTarget(target), event.MsgActionConfirm, []beans.MsgInfo{*msg})
	}

	return err == nil, err
}

func (r *subscriptionResolver) Message(ctx context.Context, target source.MessageEnum) (<-chan *source.Message, error) {
	datas := make(chan *source.Message, 100)

	fn := func(msg event.MsgInfo) {
		action := source.MessageEnum(msg.Action)
		datas <- &source.Message{
			Target:   action,
			Messages: msg.Messages,
		}
	}

	event.MsgNewSubscribe(*r.Me(ctx).Info.User.Id, event.MsgTarget(target), fn)
	go event.MsgNewUnSubscribeWithContext(ctx, *r.Me(ctx).Info.User.Id, event.MsgTarget(target), fn)

	return datas, nil
}
