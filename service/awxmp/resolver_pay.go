package awxmp

import (
	"context"
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-tools/wx/wxmp"
)

func (this *ResolverTools) Pay(ctx context.Context) (*wxmp.PayRes, error) {
	order := beans.WxmpOrder{}
	id, err := this.Create(ctx, order, nil)
	if err != nil {
		return nil, err
	}
	return this.Wxmp.Pay(&wxmp.PayOption{
		OutTradeNo: id,
		NotifyUrl:  "",
	})
}
