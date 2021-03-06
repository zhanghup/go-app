package awxmp

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/zhanghup/go-app/gs"

	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/service/awxmp/source"
	"github.com/zhanghup/go-app/service/event"
)

func (r *mutationResolver) UserRegister(ctx context.Context, input source.NewUserRegister) (bool, error) {
	me := r.Me(ctx)
	user, err := r.Wxmp.UserInfoDecrypt(me.SessionKey, input.RawData, input.EncryptedData, input.Signature, input.Iv)
	if err != nil {
		return false, err
	}
	_, err = gs.Sess(ctx).S().Table(beans.WxmpUser{}).Where("id = ?", me.Id).Update(map[string]interface{}{
		"nickname":   user.Nickname,
		"gender":     user.Gender,
		"city":       user.City,
		"province":   user.Province,
		"country":    user.Country,
		"avatar_url": user.AvatarUrl,
		"language":   user.Language,
	})
	if err != nil {
		return false, err
	}

	wxuser, err := r.Query().User(ctx)
	if err != nil {
		return false, err
	}
	go event.WxmpUserUpdate(*wxuser)
	return true, nil
}

func (r *mutationResolver) UserRegisterMobile(ctx context.Context, input source.NewUserRegisterMobile) (bool, error) {
	me := r.Me(ctx)
	mobile, err := r.Wxmp.UserMobileDecrypt(me.SessionKey, input.EncryptedData, input.Iv)
	if err != nil {
		return false, err
	}
	_, err = gs.Sess(ctx).S().Table(beans.WxmpUser{}).Where("id = ?", me.Id).Update(map[string]interface{}{
		"mobile": mobile.PhoneNumber,
	})
	if err != nil {
		return false, err
	}

	wxuser, err := r.Query().User(ctx)
	if err != nil {
		return false, err
	}
	go event.WxmpUserUpdate(*wxuser)
	return true, nil
}

func (r *queryResolver) MyInfo(ctx context.Context) (*beans.WxmpUser, error) {
	return r.Resolver.Me(ctx).User, nil
}

func (r *queryResolver) User(ctx context.Context) (*beans.WxmpUser, error) {
	wxuser := new(beans.WxmpUser)
	_, err := gs.Sess(ctx).SF(`select * from wxmp_user where id = ?`, r.Resolver.Me(ctx).Id).Get(wxuser)
	return wxuser, err
}
