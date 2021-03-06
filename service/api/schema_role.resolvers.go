package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/zhanghup/go-app/gs"

	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/service/api/source"
	"github.com/zhanghup/go-app/service/ca"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
)

func (r *mutationResolver) RoleCreate(ctx context.Context, input source.NewRole) (bool, error) {
	_, err := r.Create(ctx, new(beans.Role), input)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) RoleUpdate(ctx context.Context, id string, input source.UpdRole) (bool, error) {
	return r.Update(ctx, new(beans.Role), id, input)
}

func (r *mutationResolver) RoleRemoves(ctx context.Context, ids []string) (bool, error) {
	ok, err := r.Removes(ctx, new(beans.Role), ids)
	if err != nil {
		return false, err
	}
	if ok {
		err := gs.Sess(ctx).SF(`delete from role_user where role in :roles`, map[string]interface{}{"roles": ids}).Exec()
		if err != nil {
			return false, err
		}
		err = gs.Sess(ctx).SF(`delete from perm where role in :roles`, map[string]interface{}{"roles": ids}).Exec()
		if err != nil {
			return false, err
		}
		err = gs.Sess(ctx).SF(`delete from perm_object where role in :roles`, map[string]interface{}{"roles": ids}).Exec()
		if err != nil {
			return false, err
		}
		ca.UserCache.Clear()
	}
	return true, nil
}

func (r *mutationResolver) RolePermMenuCreate(ctx context.Context, id string, perms []string) (bool, error) {
	err := gs.Sess(ctx).SF(`delete from perm where role = :id and type = :type`, map[string]interface{}{
		"type": "menu",
		"id":   id,
	}).Exec()
	if err != nil {
		return false, err
	}

	for i, o := range perms {
		p := beans.Perm{
			Bean: beans.Bean{
				Id:     tools.PtrOfUUID(),
				Status: tools.PtrOfString("1"),
				Weight: &i,
			},
			Type: tools.PtrOfString("menu"),
			Role: &id,
			Oid:  &o,
		}
		err := gs.Sess(ctx).Insert(p)
		if err != nil {
			return false, err
		}
	}
	ca.UserCache.Clear()
	return true, nil
}

func (r *mutationResolver) RolePermObjCreate(ctx context.Context, id string, perms []source.IPermObj) (bool, error) {
	sess := gs.Sess(ctx)

	err := sess.TS(func(sess txorm.ISession) error {
		err := sess.SF(`delete from perm_object where role = :id`, map[string]interface{}{
			"id": id,
		}).Exec()
		if err != nil {
			return err
		}

		for i, o := range perms {
			p := beans.PermObject{
				Bean: beans.Bean{
					Id:     tools.PtrOfUUID(),
					Status: tools.PtrOfString("1"),
					Weight: &i,
				},
				Role:   &id,
				Object: &o.Object,
				Mask:   &o.Mask,
			}
			err = sess.Insert(p)
			if err != nil {
				return err
			}
		}
		ca.UserCache.Clear()
		return nil
	})

	return err == nil, err
}

func (r *mutationResolver) RoleWithUser(ctx context.Context, role string, uids []string) (bool, error) {
	sess := gs.Sess(ctx)
	err := sess.TS(func(sess txorm.ISession) error {
		err := sess.SF(`delete from role_user where role = :role`, map[string]interface{}{
			"role": role,
		}).Exec()
		if err != nil {
			return err
		}

		for i, o := range uids {
			p := beans.RoleUser{
				Bean: beans.Bean{
					Id:     tools.PtrOfUUID(),
					Status: tools.PtrOfString("1"),
					Weight: &i,
				},
				Role: &role,
				Uid:  &o,
			}
			err := sess.Insert(p)
			if err != nil {
				return err
			}
		}
		ca.UserCache.Clear()
		return nil
	})

	return err == nil, err
}

func (r *queryResolver) Roles(ctx context.Context, query source.QRole) (*source.Roles, error) {
	roles := make([]beans.Role, 0)
	total, err := gs.DBS(ctx).SF(`
		select u.* from role u
		join with_role on u.id = with_role.id
		where 1 = 1
		{{ if .keyword }} and u.name like concat("%",:keyword,"%") {{ end }}
		{{ if .status }} and u.status = :status {{ end }}
	`, map[string]interface{}{
		"keyword": query.Keyword,
		"status":  query.Status,
	}).With("with_role").Order("weight", "created").Page2(query.Index, query.Size, query.Count, &roles)
	return &source.Roles{Data: roles, Total: &total}, err
}

func (r *queryResolver) Role(ctx context.Context, id string) (*beans.Role, error) {
	return r.RoleLoader(ctx, id)
}

func (r *queryResolver) RoleUsers(ctx context.Context, id string) ([]string, error) {
	res := make([]string, 0)
	err := gs.DBS(ctx).SF(`select uid from role_user where role = :id`, map[string]interface{}{"id": id}).Find(&res)
	return res, err
}

func (r *queryResolver) RolePerms(ctx context.Context, id string, typeArg *string) ([]string, error) {
	result := make([]string, 0)
	err := gs.DBS(ctx).SF(`
		select oid from perm where role = :role 
		{{ if .type }} and type = :type {{ end }}
	`, map[string]interface{}{
		"role": id,
		"type": typeArg,
	}).Find(&result)
	return result, err
}

func (r *queryResolver) RolePermObjects(ctx context.Context, id string) ([]source.PermObj, error) {
	result := make([]source.PermObj, 0)
	err := gs.DBS(ctx).SF(`
		select object,mask from perm_object where role = :role 
	`, map[string]interface{}{
		"role": id,
	}).Find(&result)
	return result, err
}
