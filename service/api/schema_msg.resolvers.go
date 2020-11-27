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

func (r *mutationResolver) MsgInfoConfirm(ctx context.Context, id string, input source.NewMsgConfirm) (bool, error) {
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

func (r *mutationResolver) MsgInfoRead(ctx context.Context, id string) (bool, error) {
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

func (r *mutationResolver) MsgTemplateUpdate(ctx context.Context, id string, input source.UpdMsgTemplate) (bool, error) {
	return r.Update(ctx, beans.MsgTemplate{}, id, input)
}

func (r *queryResolver) MsgTemplates(ctx context.Context, query *source.QMsgTemplate) ([]beans.MsgTemplate, error) {
	if query == nil {
		query = &source.QMsgTemplate{}
	}
	tpls := make([]beans.MsgTemplate, 0)
	err := r.DBS.SF(`
		select 
			tpl.* 
		from 
			msg_template tpl
		where 1 = 1 
			{{ if .code }} and tpl.code like concat('%',:code,'%') {{ end }} 
			{{ if .name }} and tpl.name like concat('%',:name,'%') {{ end }} 
		order by tpl.code`,
		map[string]interface{}{
			"name": query.Name,
			"code": query.Code,
		}).Find(&tpls)
	return tpls, err
}

func (r *queryResolver) MsgTemplate(ctx context.Context, id string) (*beans.MsgTemplate, error) {
	return r.MsgTemplateLoader(ctx, id)
}

func (r *subscriptionResolver) Message(ctx context.Context) (<-chan *source.Message, error) {
	datas := make(chan *source.Message, 100)

	fn := func(action event.MsgAction, msg beans.MsgInfo) {
		datas <- &source.Message{
			Action:  action,
			Message: &msg,
		}
	}

	event.MsgNewSubscribe(*r.Me(ctx).Info.User.Id, event.MsgTargetWeb, fn)
	go event.MsgNewUnSubscribeWithContext(ctx, *r.Me(ctx).Info.User.Id, event.MsgTargetWeb, fn)

	return datas, nil
}