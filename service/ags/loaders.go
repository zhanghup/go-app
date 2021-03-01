package ags

import (
	"context"
	"github.com/zhanghup/go-app/beans"
)

func (this *Resolver) MsgInfoLoader(ctx context.Context, id string) (*beans.MsgInfo, error) {
	result := new(beans.MsgInfo)
	err := this.Loader(ctx).Object(result, "select * from msg_info where id in :keys", nil, "Id", "").Load(id, result)
	return result, err
}
