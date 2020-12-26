package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/service/api/source"
	"github.com/zhanghup/go-app/service/event"
	"github.com/zhanghup/go-tools"
)

func (r *mutationResolver) MsgTemplateUpdate(ctx context.Context, id string, input source.UpdMsgTemplate) (bool, error) {
	return r.Update(ctx, beans.MsgTemplate{}, id, input)
}

func (r *queryResolver) MsgTemplates(ctx context.Context, query source.QMsgTemplate) ([]beans.MsgTemplate, error) {
	tpls := make([]beans.MsgTemplate, 0)
	err := r.DBS(ctx).SF(`
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

func (r *queryResolver) MsgInfos(ctx context.Context, query source.QMsgInfo) ([]beans.MsgInfo, error) {
	infos := make([]beans.MsgInfo, 0)
	_, err := r.DBS(ctx).SF(`
		select info.* from msg_info info
		where 1 = 1
		{{ if .receiver }} and info.receiver = :receiver {{ end }}
		{{ if .type }} and info.type = :type {{ end }}
		{{ if .level }} and info.level = :level {{ end }}
		{{ if .target }} and info.target = :target {{ end }}
		{{ if .confirm_target }} and info.confirm_target = :level {{ end }}
		{{ if .read_target }} and info.read_target = :level {{ end }}
		{{ if .status }} and info.status = :status {{ end }}
		order by info.created desc
	`, map[string]interface{}{
		"receiver":       query.Receiver,
		"type":           query.Type,
		"level":          query.Level,
		"target":         query.Target,
		"confirm_target": query.ConfirmTarget,
		"read_target":    query.ReadTarget,
		"status":         query.Status,
	}).Page2(query.Index, query.Size, tools.Ptr.Bool(false), &infos)
	return infos, err
}

func (r *queryResolver) MsgHistorys(ctx context.Context, query source.QMsgHistory) ([]beans.MsgHistory, error) {
	infos := make([]beans.MsgHistory, 0)
	_, err := r.DBS(ctx).SF(`
		select his.* from msg_history his
		where 1 = 1
		{{ if .info }} and his.info = :info {{ end }}
	`, map[string]interface{}{
		"info": query.Info,
	}).Order("-his.created").Page2(query.Index, query.Size, tools.Ptr.Bool(false), &infos)
	return infos, err
}

func (r *subscriptionResolver) Message(ctx context.Context) (<-chan *source.Message, error) {
	datas := make(chan *source.Message, 100)

	fn := func(tpl beans.MsgTemplate, msg beans.MsgInfo) {
		datas <- &source.Message{
			Message:  &msg,
			Template: &tpl,
		}
	}

	event.MsgNewSubscribe(*r.Me(ctx).Info.User.Id, event.MsgTargetWeb, fn)
	go event.MsgNewUnSubscribeWithContext(ctx, *r.Me(ctx).Info.User.Id, event.MsgTargetWeb, fn)

	return datas, nil
}
