package awxmp

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/cfg"
	"github.com/zhanghup/go-app/service/event"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/tog"
	"github.com/zhanghup/go-tools/wx/wxmp"
	"io"
	"strings"
	"time"
)

type PayOption struct {
	Price       int
	Description string
	Otype       string
	Oid         string

	TimeExpire *int64
	Attach     *string
	Currency   *string // 支付货币
	GoodsTag   *string
}

func (this *ResolverTools) Pay(ctx context.Context, opt *PayOption) (*wxmp.PayRes, error) {
	me := this.Me(ctx)
	order := beans.WxmpOrder{
		Uid:    &me.Id,
		Openid: &me.Openid,
		Otype:  &opt.Otype,
		Oid:    &opt.Oid,
		State:  tools.PtrOfString("0"),
		Commit: tools.PtrOfInt64(time.Now().Unix()),
		Price:  tools.PtrOfInt(opt.Price),
	}
	id := tools.UUID()
	id = strings.ReplaceAll(id, "-", "")
	order.Id = &id
	_, err := this.Create(ctx, &order, nil)
	if err != nil {
		return nil, err
	}
	return this.Wxmp.Pay(&wxmp.PayOption{
		OutTradeNo:  id,
		NotifyUrl:   cfg.Config.Host + "/zpx/wxmp/pay/callback",
		Openid:      me.Openid,
		TotalPrice:  opt.Price,
		Description: opt.Description,
		TimeExpire:  opt.TimeExpire,
		Attach:      opt.Attach,
		Currency:    opt.Currency,
		GoodsTag:    opt.GoodsTag,
	})
}

func (this *ResolverTools) PayCancelAction(ctx context.Context, id, ty string) (bool, error) {
	id, err := this.Sess(ctx).SF(`select id from wxmp_order where otype = ? and oid = ?`, ty, id).String()
	if err != nil {
		return false, err
	}
	err = this.Wxmp.PayCancel(id)
	return err == nil, err
}

func PayCallback(c *gin.Context) {
	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		tog.Error("【微信支付】 %s", err.Error())
		return
	}
	go event.WxmpPayCallbackPush(data)
	fmt.Println(string(data), "------------------")
}
