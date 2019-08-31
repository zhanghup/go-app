package api

import (
	"context"
	"fmt"
	"github.com/zhanghup/go-app"
	"github.com/zhanghup/go-app/api/gs"
)

func (this queryResolver) Users(ctx context.Context, query gs.QUser) (*gs.Users, error) {
	users := make([]app.User, 0)
	this.DB(ctx).SF(`
		select * from {{ table "user" }} u
		where 1 = 1
	`)
	fmt.Println(users)
	return nil, nil
}

func (this queryResolver) User(ctx context.Context, id string) (*app.User, error) {
	panic("implement me")
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
	_, err := this.DB(ctx).Table(new(app.User)).SF(`
		delete from {{ table "user" }} where id in :ids
	`, map[string]interface{}{"ids": id}).Execute()
	return err == nil, err
}
