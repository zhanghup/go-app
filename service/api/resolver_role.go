package api

import (
	"context"
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/service/api/lib"
	"github.com/zhanghup/go-tools"
)

func (this *Resolver) RoleLoader(ctx context.Context, id string) (*beans.Role, error) {
	result := new(beans.Role)
	_, err := this.Loader(ctx).Object(new(beans.Role)).Load(id, result)
	return result, err
}

func (this queryResolver) Roles(ctx context.Context, query lib.QRole) (*lib.Roles, error) {
	roles := make([]beans.Role, 0)
	_, total, err := this.DB.SF(`
		select * from {{ table "role" }} u
		where 1 = 1
	`).Page2(query.Index, query.Size, query.Count, &roles)
	return &lib.Roles{Data: roles, Total: &total}, err
}

func (this queryResolver) Role(ctx context.Context, id string) (*beans.Role, error) {
	return this.RoleLoader(ctx, id)
}

func (this mutationResolver) RoleCreate(ctx context.Context, input lib.NewRole) (*beans.Role, error) {
	id, err := this.Create(ctx, new(beans.Role), input)
	if err != nil {
		return nil, err
	}
	return this.RoleLoader(ctx, id)
}

func (this mutationResolver) RoleUpdate(ctx context.Context, id string, input lib.UpdRole) (bool, error) {
	return this.Update(ctx, new(beans.Role), id, input)
}

func (this mutationResolver) RoleRemoves(ctx context.Context, id []string) (bool, error) {
	return this.Removes(ctx, new(beans.Role), id)
}

func (this queryResolver) RolePerms(ctx context.Context, id string, typeArg *string) ([]string, error) {
	result := make([]string, 0)
	err := this.DB.SF(`
		select oid from {{ table "perm" }} where role = :role 
		{{ if .type }} and type = :type {{ end }}
	`, map[string]interface{}{
		"role": id,
		"type": typeArg,
	}).Find(&result)
	return result, err
}

func (this mutationResolver) RolePermCreate(ctx context.Context, id string, typeArg string, perms []string) (bool, error) {
	_, err := this.DB.SF(`delete from {{ table "perm" }} where role = :id and type = :type`, map[string]interface{}{
		"type": typeArg,
		"id":   id,
	}).Execute()
	if err != nil {
		return false, err
	}
	for i, o := range perms {
		p := beans.Perm{
			Bean: beans.Bean{
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
	_, err := this.DB.SF(`delete from {{ table "role_user" }} where uid = :uid`, map[string]interface{}{
		"uid": uid,
	}).Execute()
	if err != nil {
		return false, err
	}
	for i, o := range roles {
		p := beans.RoleUser{
			Bean: beans.Bean{
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
