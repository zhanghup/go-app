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
	total, err := this.DBS.SF(`
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

func (this queryResolver) RolePermObjects(ctx context.Context, id string) ([]lib.PermObj, error) {
	result := make([]lib.PermObj, 0)
	err := this.DBS.SF(`
		select object,mask from {{ table "perm_object" }} where role = :role 
	`, map[string]interface{}{
		"role": id,
	}).Find(&result)
	return result, err
}

func (this queryResolver) RolePerms(ctx context.Context, id string, typeArg *string) ([]string, error) {
	result := make([]string, 0)
	err := this.DBS.SF(`
		select oid from {{ table "perm" }} where role = :role 
		{{ if .type }} and type = :type {{ end }}
	`, map[string]interface{}{
		"role": id,
		"type": typeArg,
	}).Find(&result)
	return result, err
}

func (this mutationResolver) RolePermCreate(ctx context.Context, id string, typeArg string, perms []string) (bool, error) {
	err := this.DBS.SF(`delete from {{ table "perm" }} where role = :id and type = :type`, map[string]interface{}{
		"type": typeArg,
		"id":   id,
	}).Exec()
	if err != nil {
		return false, err
	}
	for i, o := range perms {
		p := beans.Perm{
			Bean: beans.Bean{
				Id:     tools.Ptr.Uid(),
				Status: tools.Ptr.Int(1),
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

func (this mutationResolver) RolePermObjCreate(ctx context.Context, id string, perms []lib.IPermObj) (result bool, err error) {
	sess := this.DBS.NewSession(ctx)

	err = sess.SF(`delete from {{ table "perm_object" }} where role = :id`, map[string]interface{}{
		"id": id,
	}).Exec()
	if err != nil {
		return
	}

	for i, o := range perms {
		p := beans.PermObject{
			Bean: beans.Bean{
				Id:     tools.Ptr.Uid(),
				Status: tools.Ptr.Int(1),
				Weight: &i,
			},
			Role:   &id,
			Object: &o.Object,
			Mask:   &o.Mask,
		}
		err = sess.Insert(p)
		if err != nil {
			return
		}
	}
	return true, nil
}

func (this mutationResolver) RoleToUser(ctx context.Context, uid string, roles []string) (bool, error) {
	sess := this.DBS.NewSession(ctx)
	err := sess.SF(`delete from {{ table "role_user" }} where uid = :uid`, map[string]interface{}{
		"uid": uid,
	}).Exec()
	if err != nil {
		return false, err
	}
	for i, o := range roles {
		p := beans.RoleUser{
			Bean: beans.Bean{
				Id:     tools.Ptr.Uid(),
				Status: tools.Ptr.Int(1),
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
