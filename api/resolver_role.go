package api

import (
	"context"
	"github.com/zhanghup/go-app"
	"github.com/zhanghup/go-app/api/gs"
)

func (this *Resolver) RoleLoader(ctx context.Context, id string) (*app.Role, error) {
	obj, err := this.Loader(ctx).Object(new(app.Role)).Load(id)
	if err != nil {
		return nil, err
	}
	role, ok := obj.(app.Role)
	if !ok {
		return nil, nil
	}
	return &role, nil
}

func (this queryResolver) Roles(ctx context.Context, query gs.QRole) (*gs.Roles, error) {
	roles := make([]*app.Role, 0)
	_, total, err := this.DB.SF(`
		select * from {{ table "role" }} u
		where 1 = 1
	`).Page2(query.Index, query.Size, query.Count, &roles)
	return &gs.Roles{Data: roles, Total: &total}, err
}

func (this queryResolver) Role(ctx context.Context, id string) (*app.Role, error) {
	return this.RoleLoader(ctx, id)
}

func (this mutationResolver) RoleCreate(ctx context.Context, input gs.NewRole) (*app.Role, error) {
	id, err := this.Create(ctx, new(app.Role), input)
	if err != nil {
		return nil, err
	}
	return this.RoleLoader(ctx, id)
}

func (this mutationResolver) RoleUpdate(ctx context.Context, id string, input gs.UpdRole) (bool, error) {
	return this.Update(ctx, new(app.Role), id, input)
}

func (this mutationResolver) RoleRemoves(ctx context.Context, id []string) (bool, error) {
	return this.Removes(ctx, new(app.Role), id)
}
