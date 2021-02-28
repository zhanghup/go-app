package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"github.com/zhanghup/go-app/gs"

	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/service/api/source"
	"github.com/zhanghup/go-app/service/ca"
	"github.com/zhanghup/go-tools"
)

func (r *mutationResolver) AccountCreate(ctx context.Context, input source.NewAccount) (string, error) {
	acc := new(beans.Account)
	acc.Salt = tools.PtrOfUUID()
	if *input.Type == "password" {
		if input.Username == nil || input.Password == nil {
			return "", errors.New("用户名密码不能为空")
		}
		password := tools.Password(*input.Password, *acc.Salt)
		input.Password = &password
	}

	acc.Uid = input.UID
	acc.Type = input.Type
	acc.Username = input.Username
	acc.Password = input.Password
	acc.Default = input.Default
	if input.Default != nil && *input.Default == 1 {
		err := gs.Sess(ctx).SF("update account set `default` = 0 where uid = :uid", map[string]interface{}{"uid": input.UID}).Exec()
		if err != nil {
			return "", err
		}
	}

	return r.Create(ctx, acc, nil)
}

func (r *mutationResolver) AccountUpdate(ctx context.Context, id string, input source.UpdAccount) (bool, error) {
	acc, err := r.AccountLoader(ctx, id)
	if err != nil {
		return false, err
	}
	if acc == nil {
		return false, errors.New("账户不存在")
	}

	if acc.Salt == nil {
		acc.Salt = tools.PtrOfUUID()
		err := gs.Sess(ctx).SF("update account set salt = :salt", map[string]interface{}{"salt": acc.Salt}).Exec()
		if err != nil {
			return false, err
		}
	}

	if *input.Type == "password" {
		if input.Username == nil || input.Password == nil {
			return false, errors.New("用户名密码不能为空")
		}

		if len(*input.Password) != 136 {
			password := tools.Password(*input.Password, *acc.Salt)
			input.Password = &password
		}

	}

	if input.Default != nil && *input.Default == 1 {
		err := gs.Sess(ctx).SF("update account set `default` = 0 where uid = :uid", map[string]interface{}{"uid": acc.Uid}).Exec()
		if err != nil {
			return false, err
		}
	}

	// 删除登录状态缓存
	ca.UserCache.RemoveByUser(*acc.Uid)

	return r.Update(ctx, acc, id, input)
}

func (r *mutationResolver) AccountRemoves(ctx context.Context, ids []string) (bool, error) {
	uids := make([]string, 0)
	err := gs.DB().In("id", ids).Cols("uid").Find(&uids)
	if err != nil {
		return false, err
	}
	for _, s := range uids {
		// 删除登录状态缓存
		ca.UserCache.RemoveByUser(s)
	}
	return r.Removes(ctx, new(beans.Account), ids)
}

func (r *queryResolver) Accounts(ctx context.Context, query source.QAccount) (*source.Accounts, error) {
	account := make([]beans.Account, 0)
	total, err := gs.DBS().SF(`
		select * from account 
		where 1 = 1
		{{ if .uid }} and account.uid = :uid {{ end }}
		{{ if .username }} and account.username = :username {{ end }}
		{{ if .type }} and account.type = :type {{ end }}
		{{ if .status }} and account.status = :status {{ end }}
	`, map[string]interface{}{
		"uid":      query.UID,
		"username": query.Username,
		"type":     query.Type,
		"status":   query.Status,
	}).Page2(query.Index, query.Size, query.Count, &account)
	return &source.Accounts{Data: account, Total: &total}, err
}

func (r *queryResolver) Account(ctx context.Context, id string) (*beans.Account, error) {
	return r.AccountLoader(ctx, id)
}
