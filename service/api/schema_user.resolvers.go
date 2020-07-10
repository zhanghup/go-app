package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/service/api/lib"
	"github.com/zhanghup/go-tools"
)

func (r *mutationResolver) UserCreate(ctx context.Context, input lib.NewUser) (*beans.User, error) {
	user := &beans.User{
		Salt: tools.Ptr.String(tools.Str.RandString(10)),
	}
	input.Password = tools.Crypto.Password(input.Password, *user.Salt)

	id, err := r.Create(ctx, user, input)
	if err != nil {
		return nil, err
	}

	return r.UserLoader(ctx, id)
}

func (r *mutationResolver) UserUpdate(ctx context.Context, id string, input lib.UpdUser) (bool, error) {
	user, err := r.UserLoader(ctx, id)
	if err != nil {
		return false, err
	}
	if input.Password == "" {
		return false, errors.New("密码不能为空")
	}
	if *user.Password != input.Password {
		input.Password = tools.Crypto.Password(input.Password, *user.Salt)
	}
	return r.Update(ctx, new(beans.User), id, input)
}

func (r *mutationResolver) UserRemoves(ctx context.Context, ids []string) (bool, error) {
	if tools.Str.Contains(ids, "root") {
		return false, errors.New("root用户无法删除")
	}
	return r.Removes(ctx, new(beans.User), ids)
}

func (r *queryResolver) Users(ctx context.Context, query lib.QUser) (*lib.Users, error) {
	users := make([]beans.User, 0)
	total, err := r.DBS.SF(`
		select * from user u
		where 1 = 1
	`).Page2(query.Index, query.Size, query.Count, &users)
	return &lib.Users{Data: users, Total: &total}, err
}

func (r *queryResolver) User(ctx context.Context, id string) (*beans.User, error) {
	return r.UserLoader(ctx, id)
}
