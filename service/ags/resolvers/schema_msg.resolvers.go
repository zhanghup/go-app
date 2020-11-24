package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/service/event"
)

func (r *subscriptionResolver) MessageNew(ctx context.Context, typeArg string) (<-chan []beans.MsgInfo, error) {
	datas := make(chan []beans.MsgInfo, 100)
	event.MsgNewSubscribe(*r.Me(ctx).Info.User.Id, typeArg, func(msg []beans.MsgInfo) {
		datas <- msg
	})
	return datas, nil
}

func (r *subscriptionResolver) MessageRead(ctx context.Context, typeArg string) (<-chan []beans.MsgInfo, error) {
	datas := make(chan []beans.MsgInfo, 100)
	event.MsgReadSubscribe(*r.Me(ctx).Info.User.Id, typeArg, func(msg []beans.MsgInfo) {
		datas <- msg
	})
	return datas, nil
}

func (r *subscriptionResolver) MessageConfirm(ctx context.Context, typeArg string) (<-chan []beans.MsgInfo, error) {
	datas := make(chan []beans.MsgInfo, 100)
	event.MsgConfirmSubscribe(*r.Me(ctx).Info.User.Id, typeArg, func(msg []beans.MsgInfo) {
		datas <- msg
	})
	return datas, nil
}
