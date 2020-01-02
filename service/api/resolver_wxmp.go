package api

import (
	"context"
	"errors"
	"github.com/zhanghup/go-wxmp"
)

func (this queryResolver) WxmpMenus(ctx context.Context) ([]wxmp.Button, error) {
	if this.wxmp == nil{
		return nil,errors.New("微信公众号未开通")
	}
	return this.wxmp.Menu().Get()
}

func (this mutationResolver) WxmpMenuCreate(ctx context.Context, input []wxmp.Button) (bool, error) {
	if this.wxmp == nil{
		return false,errors.New("微信公众号未开通")
	}
	err := this.wxmp.Menu().Create(input)
	return err == nil,err
}

func (this mutationResolver) WxmpMenuRemoves(ctx context.Context) (bool, error) {
	if this.wxmp == nil{
		return false,errors.New("微信公众号未开通")
	}
	err := this.wxmp.Menu().Delete()
	return err == nil,err
}