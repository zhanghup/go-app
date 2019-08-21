package api

import (
	"context"
	"github.com/zhanghup/go-app"
	"github.com/zhanghup/go-app/api/gs"
)

func (this queryResolver) Users(ctx context.Context, query *gs.QUser) (*gs.Users, error) {
	users := make([]app.User, 0)
	this.DB.Context(ctx).SF(`
		select * from {{ table "user" }} u
		where 1 = 1
	`)
	return nil, nil
}

func (this queryResolver) User(ctx context.Context, id *string) (*app.User, error) {
	panic("implement me")
}

func (this mutationResolver) CreateUser(ctx context.Context, input gs.NewUser) (*app.User, error) {
	panic("implement me")
}

func (this mutationResolver) UpdateUser(ctx context.Context, id *string, input gs.UpdUser) (*bool, error) {
	panic("implement me")
}

func (this mutationResolver) RemoveUser(ctx context.Context, id *string) (*bool, error) {
	panic("implement me")
}
