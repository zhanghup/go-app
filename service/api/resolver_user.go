package api

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/service/api/lib"

	"github.com/zhanghup/go-tools"
)

func (this *Resolver) UserLoader(ctx context.Context, id string) (*beans.User, error) {
	result := new(beans.User)
	err := this.Loader(ctx).Object(result, "select * from user where id in :keys", nil, "Id", "").Load(id, result)
	return result, err
}

func (this queryResolver) Users(ctx context.Context, query lib.QUser) (*lib.Users, error) {
	users := make([]beans.User, 0)
	total, err := this.DBS.SF(`
		select * from {{ table "user" }} u
		where 1 = 1
	`).Page2(query.Index, query.Size, query.Count, &users)
	return &lib.Users{Data: users, Total: &total}, err
}

func (this queryResolver) User(ctx context.Context, id string) (*beans.User, error) {
	return this.UserLoader(ctx, id)
}

func (this mutationResolver) UserCreate(ctx context.Context, input lib.NewUser) (*beans.User, error) {
	user := &beans.User{
		Salt: tools.Ptr.String(tools.Str.RandString(10)),
	}
	input.Password = tools.Crypto.Password(input.Password, *user.Salt)

	id, err := this.Create(ctx, user, input)
	if err != nil {
		return nil, err
	}

	return this.UserLoader(ctx, id)
}

func (this mutationResolver) UserUpdate(ctx context.Context, id string, input lib.UpdUser) (bool, error) {
	user, err := this.UserLoader(ctx, id)
	if err != nil {
		return false, err
	}
	if input.Password == "" {
		return false, errors.New("密码不能为空")
	}
	if *user.Password != input.Password {
		input.Password = tools.Crypto.Password(input.Password, *user.Salt)
	}
	return this.Update(ctx, new(beans.User), id, input)
}

func (this mutationResolver) UserRemoves(ctx context.Context, id []string) (bool, error) {
	return this.Removes(ctx, new(beans.User), id)
}
