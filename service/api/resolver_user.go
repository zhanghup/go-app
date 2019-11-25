package api

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zhanghup/go-app/service/api/lib"

	"github.com/zhanghup/go-app"
	"github.com/zhanghup/go-tools"
)

func (this *Resolver) UserLoader(ctx context.Context, id string) (*app.User, error) {
	obj, err := this.Loader(ctx).Object(new(app.User)).Load(id)
	if err != nil {
		return nil, err
	}
	user, ok := obj.(app.User)
	if !ok {
		return nil, nil
	}
	return &user, nil
}

func (this queryResolver) Users(ctx context.Context, query lib.QUser) (*lib.Users, error) {
	users := make([]app.User, 0)
	_, total, err := this.DB.SF(`
		select * from {{ table "user" }} u
		where 1 = 1
	`).Page2(query.Index, query.Size, query.Count, &users)
	return &lib.Users{Data: users, Total: &total}, err
}

func (this queryResolver) User(ctx context.Context, id string) (*app.User, error) {
	return this.UserLoader(ctx, id)
}

func (this mutationResolver) UserCreate(ctx context.Context, input lib.NewUser) (*app.User, error) {
	id, err := this.Create(ctx, new(app.User), input)
	if err != nil {
		return nil, err
	}
	if input.Password != "" {
		pwd := input.Password
		salt := *tools.ObjectString()
		_, err := this.DB.Table(new(app.User)).Where("id = ?", id).Update(map[string]interface{}{
			"password": tools.Password(pwd, salt),
			"slat":     salt,
		})
		if err != nil {
			return nil, err
		}
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
		input.Password = tools.Password(input.Password, *user.Slat)
	}
	return this.Update(ctx, new(app.User), id, input)
}

func (this mutationResolver) UserRemoves(ctx context.Context, id []string) (bool, error) {
	return this.Removes(ctx, new(app.User), id)
}
