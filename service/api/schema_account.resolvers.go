package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"github.com/zhanghup/go-tools/database/txorm"

	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/service/api/source"
	"github.com/zhanghup/go-tools"
)

func (r *mutationResolver) AccountCreate(ctx context.Context, input source.NewAccount) (string, error) {
	acc := new(beans.Account)
	acc.Salt = tools.Ptr.Uid()

	if input.Type == "password" {
		if input.Username == nil || input.Password == nil {
			return "", errors.New("用户名密码不能为空")
		}
		password := tools.Crypto.Password(*input.Password, *acc.Salt)
		input.Password = &password
	}

	acc.Uid = &input.UID
	acc.Type = &input.Type
	acc.Username = input.Username
	acc.Password = input.Password
	acc.Admin = input.Admin
	acc.Default = input.Default

	session := r.DBS.NewSession(ctx)
	if input.Default != nil && *input.Default == 1 {
		err := session.TS(func(sess *txorm.Session) error {
			return sess.SF("update account set `default` = 0 where uid = :uid", map[string]interface{}{"uid": input.UID}).Exec()
		})
		if err != nil {
			return "", err
		}
	}

	return r.Create(session.Context(), acc, nil)
}

func (r *mutationResolver) AccountUpdate(ctx context.Context, id string, input source.UpdAccount) (bool, error) {
	acc, err := r.AccountLoader(ctx, id)
	if err != nil {
		return false, err
	}
	if acc == nil {
		return false, errors.New("账户不存在")
	}

	session := r.DBS.NewSession(ctx)
	if acc.Salt == nil {
		acc.Salt = tools.Ptr.Uid()
		err = session.SF("update account set salt = :salt", map[string]interface{}{"salt": acc.Salt}).Exec()
		if err != nil {
			return false, err
		}
	}

	if input.Type == "password" {
		if input.Username == nil || input.Password == nil {
			return false, errors.New("用户名密码不能为空")
		}

		if len(*input.Password) != 136 {
			password := tools.Crypto.Password(*input.Password, *acc.Salt)
			input.Password = &password
		}

	}

	if input.Default != nil && *input.Default == 1 {
		err := session.TS(func(sess *txorm.Session) error {
			return sess.SF("update account set `default` = 0 where uid = :uid", map[string]interface{}{"uid": acc.Uid}).Exec()
		})
		if err != nil {
			return false, err
		}
	}

	return r.Update(session.Context(), acc, id, input)
}

func (r *mutationResolver) AccountRemoves(ctx context.Context, ids []string) (bool, error) {
	return r.Removes(ctx, new(beans.Account), ids)
}

func (r *queryResolver) Accounts(ctx context.Context, query source.QAccount) (*source.Accounts, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Account(ctx context.Context, id string) (*beans.Account, error) {
	return r.AccountLoader(ctx, id)
}