package api

import (
	"context"
	"github.com/zhanghup/go-app"
	"github.com/zhanghup/go-app/api/gs"
)
func (this queryResolver) Roles(ctx context.Context, query gs.QRole) (*gs.Roles, error) {
	panic("implement me")
}

func (this queryResolver) Role(ctx context.Context, id string) (*app.Role, error) {
	panic("implement me")
}

func (this mutationResolver) RoleCreate(ctx context.Context, input gs.NewRole) (*app.Role, error) {
	panic("implement me")
}

func (this mutationResolver) RoleUpdate(ctx context.Context, id string, input gs.UpdRole) (bool, error) {
	panic("implement me")
}

func (this mutationResolver) RoleRemoves(ctx context.Context, id []string) (bool, error) {
	panic("implement me")
}

