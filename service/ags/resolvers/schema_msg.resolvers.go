package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/zhanghup/go-app/service/ags/source"
	"github.com/zhanghup/go-app/service/event"
)

func (r *subscriptionResolver) Message(ctx context.Context, typeArg string) (<-chan *source.Message, error) {
	datas := make(chan *source.Message, 100)

	fn := func(msg event.MsgInfo) {
		action := source.MessageEnum(msg.Action)
		datas <- &source.Message{
			Action:   &action,
			Messages: msg.Messages,
		}
	}

	event.MsgNewSubscribe(*r.Me(ctx).Info.User.Id, event.MsgTarget(typeArg), fn)
	go event.MsgNewUnSubscribeWithContext(ctx, *r.Me(ctx).Info.User.Id, event.MsgTarget(typeArg), fn)

	return datas, nil
}
