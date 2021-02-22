package awxmp

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/cfg"
	"github.com/zhanghup/go-app/service/ags"
	"github.com/zhanghup/go-app/service/event"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
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
	res, err := this.Wxmp.Pay(&wxmp.PayOption{
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
	if err != nil{
		return nil,err
	}

	return res, err
}

func (this *ResolverTools) PayCancelAction(ctx context.Context, id, ty string) (bool, error) {
	id, err := this.Sess(ctx).SF(`select id from wxmp_order where otype = ? and oid = ?`, ty, id).String()
	if err != nil {
		return false, err
	}
	err = this.Wxmp.PayCancel(id)
	return err == nil, err
}

func PayCallback(wxEngine wxmp.IEngine) func(c *gin.Context) {
	dbs := txorm.NewEngine(ags.DefaultDB())
	return func(c *gin.Context) {
		data, err := io.ReadAll(c.Request.Body)
		if err != nil {
			tog.Error("【微信支付】 %s", err.Error())
			c.JSON(200, map[string]interface{}{
				"code":    "ERROR",
				"message": err.Error(),
			})
			return
		}

		res, err := wxEngine.PayDecrypt(data)
		if err != nil {
			tog.Error("【微信支付】 %s", err.Error())
			c.JSON(200, map[string]interface{}{
				"code":    "ERROR",
				"message": err.Error(),
			})
			return
		}

		err = dbs.SF(`update wxmp_order set 
			updated = unix_timestamp(now()),
			state = :state,
			price_user = :price,
			pay_time = unix_timestamp(now()),
			transaction_id = :transaction_id,
			message = :message
			where id = :id
		`, map[string]interface{}{
			"id":             res.OutTradeNo,
			"price":          res.Amount.PayerTotal,
			"transaction_id": res.TransactionId,
			"message":        res.TradeStateDesc,
			"state": func() string {
				if res.TradeState == "SUCCESS" {
					return "3"
				} else {
					return "4"
				}
			}(),
		}).Exec()
		if err != nil {
			tog.Error("【微信支付】 %s", err.Error())
			c.JSON(200, map[string]interface{}{
				"code":    "ERROR",
				"message": err.Error(),
			})
			return
		}

		order := new(beans.WxmpOrder)
		ok, err := ags.DefaultDB().Where("id = ?", res.OutTradeNo).Get(order)
		if err != nil {
			tog.Error("【微信支付】 %s", err.Error())
			c.JSON(200, map[string]interface{}{
				"code":    "ERROR",
				"message": err.Error(),
			})
			return
		}
		if !ok {
			tog.Error("【微信支付】 订单不存在")
			c.JSON(200, map[string]interface{}{
				"code":    "ERROR",
				"message": "订单不存在",
			})
			return
		}

		go event.WxmpPayCallbackPush(*order)
		c.JSON(200, map[string]interface{}{
			"code":    "SUCCESS",
			"message": "",
		})

	}
}
