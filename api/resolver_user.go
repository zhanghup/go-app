package api

import (
	"context"
	"github.com/zhanghup/go-app"
	"github.com/zhanghup/go-app/api/gs"
	"github.com/zhanghup/go-tools"
)

func (this *Resolver) UserLoader(ctx context.Context, id string) (*app.User, error) {
	obj, err := this.Loader(ctx).Common(new(app.User)).Load(id)
	if err != nil {
		return nil, err
	}
	user, ok := obj.(app.User)
	if !ok {
		return nil, nil
	}
	return &user, nil
}

func (this queryResolver) Users(ctx context.Context, query gs.QUser) (*gs.Users, error) {
	users := make([]*app.User, 0)
	_, total, err := this.DB.SF(`
		select * from {{ table "user" }} u
		where 1 = 1
	`).Page2(query.Index, query.Size, query.Count, &users)
	tools.Str().JSONStringPrintln(users)
	return &gs.Users{Data: users, Total: &total}, err
}

func (this queryResolver) User(ctx context.Context, id string) (*app.User, error) {
	return this.UserLoader(ctx, id)
}

func (this mutationResolver) UserCreate(ctx context.Context, input gs.NewUser) (*app.User, error) {
	id, err := this.Create(ctx, new(app.User), input)
	if err != nil {
		return nil, err
	}
	return this.UserLoader(ctx, id)
}

func (this mutationResolver) UserUpdate(ctx context.Context, id string, input gs.UpdUser) (bool, error) {
	return this.Update(ctx, new(app.User), id, input)
}

func (this mutationResolver) UserRemoves(ctx context.Context, id []string) (bool, error) {
	return this.Removes(ctx, new(app.User), id)
}
