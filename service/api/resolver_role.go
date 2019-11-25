package api

import (
	"context"
	"github.com/zhanghup/go-app"
	"github.com/zhanghup/go-app/service/api/lib"
	"github.com/zhanghup/go-tools"
)

func (this *Resolver) RoleLoader(ctx context.Context, id string) (*app.Role, error) {
	result := new(app.Role)
	_, err := this.Loader(ctx).Object(new(app.Role)).Load(id, result)
	return result, err
}

func (this queryResolver) Roles(ctx context.Context, query lib.QRole) (*lib.Roles, error) {
	roles := make([]app.Role, 0)
	_, total, err := this.DB.SF(`
		select * from {{ table "role" }} u
		where 1 = 1
	`).Page2(query.Index, query.Size, query.Count, &roles)
	return &lib.Roles{Data: roles, Total: &total}, err
}

func (this queryResolver) Role(ctx context.Context, id string) (*app.Role, error) {
	return this.RoleLoader(ctx, id)
}

func (this mutationResolver) RoleCreate(ctx context.Context, input lib.NewRole) (*app.Role, error) {
	id, err := this.Create(ctx, new(app.Role), input)
	if err != nil {
		return nil, err
	}
	return this.RoleLoader(ctx, id)
}

func (this mutationResolver) RoleUpdate(ctx context.Context, id string, input lib.UpdRole) (bool, error) {
	return this.Update(ctx, new(app.Role), id, input)
}

func (this mutationResolver) RoleRemoves(ctx context.Context, id []string) (bool, error) {
	return this.Removes(ctx, new(app.Role), id)
}

func (this mutationResolver) RolePermCreate(ctx context.Context, id string, typeArg string, perms []string) (bool, error) {
	_, err := this.DB.SF(`delete * from {{ table "perm" }} where role = :id and type = :type`, map[string]interface{}{
		"type": typeArg,
	}).Execute()
	if err != nil {
		return false, err
	}
	for i, o := range perms {
		p := app.Perm{
			Bean: app.Bean{
				Id:     tools.ObjectString(),
				Status: tools.Ptr().Int(1),
				Weight: &i,
			},
			Type: &typeArg,
			Role: &id,
			Oid:  &o,
		}
		_, err := this.DB.Insert(p)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

func (this mutationResolver) RoleToUser(ctx context.Context, uid string, roles []string) (bool, error) {
	_, err := this.DB.SF(`delete * from {{ table "role_user" }} where uid = :uid`, map[string]interface{}{
		"uid": uid,
	}).Execute()
	if err != nil {
		return false, err
	}
	for i, o := range roles {
		p := app.RoleUser{
			Bean: app.Bean{
				Id:     tools.ObjectString(),
				Status: tools.Ptr().Int(1),
				Weight: &i,
			},
			Role: &o,
			Uid:  &uid,
		}
		_, err := this.DB.Insert(p)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}
