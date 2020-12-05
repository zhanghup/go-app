package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/service/api/source"
	"github.com/zhanghup/go-app/service/event"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
	"github.com/zhanghup/go-tools/tog"
)

func (r *mutationResolver) RoleCreate(ctx context.Context, input source.NewRole) (bool, error) {
	sess := r.DBS.NewSession(ctx)
	_, err := r.Create(sess.Context(), new(beans.Role), input)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) RoleUpdate(ctx context.Context, id string, input source.UpdRole) (bool, error) {
	sess := r.DBS.NewSession(ctx)
	return r.Update(sess.Context(), new(beans.Role), id, input)
}

func (r *mutationResolver) RoleRemoves(ctx context.Context, ids []string) (bool, error) {
	sess := r.DBS.NewSession(ctx)
	return r.Removes(sess.Context(), new(beans.Role), ids)
}

func (r *mutationResolver) RolePermCreate(ctx context.Context, id string, typeArg string, perms []string) (bool, error) {
	sess := r.DBS.NewSession(ctx)
	err := sess.TS(func(sess *txorm.Session) error {
		err := sess.SF(`delete from perm where role = :id and type = :type`, map[string]interface{}{
			"type": typeArg,
			"id":   id,
		}).Exec()
		if err != nil {
			return err
		}

		for i, o := range perms {
			p := beans.Perm{
				Bean: beans.Bean{
					Id:     tools.Ptr.Uid(),
					Status: tools.Ptr.String("1"),
					Weight: &i,
				},
				Type: &typeArg,
				Role: &id,
				Oid:  &o,
			}
			err := sess.Insert(p)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err == nil, err
}

func (r *mutationResolver) RolePermObjCreate(ctx context.Context, id string, perms []source.IPermObj) (bool, error) {
	sess := r.DBS.NewSession(ctx)

	err := sess.TS(func(sess *txorm.Session) error {
		err := sess.SF(`delete from perm_object where role = :id`, map[string]interface{}{
			"id": id,
		}).Exec()
		if err != nil {
			return err
		}

		for i, o := range perms {
			p := beans.PermObject{
				Bean: beans.Bean{
					Id:     tools.Ptr.Uid(),
					Status: tools.Ptr.String("1"),
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
		return nil
	})

	return err == nil, err
}

func (r *mutationResolver) RoleToUser(ctx context.Context, uid string, roles []string) (bool, error) {
	sess := r.DBS.NewSession(ctx)
	err := sess.TS(func(sess *txorm.Session) error {
		err := sess.SF(`delete from role_user where uid = :uid`, map[string]interface{}{
			"uid": uid,
		}).Exec()
		if err != nil {
			return err
		}

		for i, o := range roles {
			p := beans.RoleUser{
				Bean: beans.Bean{
					Id:     tools.Ptr.Uid(),
					Status: tools.Ptr.String("1"),
					Weight: &i,
				},
				Role: &o,
				Uid:  &uid,
			}
			err := sess.Insert(p)
			if err != nil {
				return err
			}
		}
		return nil
	})

	// 用户角色变化推送
	go func() {
		user, err := r.UserLoader(ctx, uid)
		if err != nil {
			tog.Error(err.Error())
			return
		}
		if user != nil {
			tog.Error("用户不存在")
			return
		}
		event.UserRoleChange(*user)
	}()
	return err == nil, err
}

func (r *mutationResolver) RoleWithUser(ctx context.Context, role string, uids []string) (bool, error) {
	sess := r.DBS.NewSession(ctx)
	err := sess.TS(func(sess *txorm.Session) error {
		err := sess.SF(`delete from role_user where role = :role`, map[string]interface{}{
			"role": role,
		}).Exec()
		if err != nil {
			return err
		}

		for i, o := range uids {
			p := beans.RoleUser{
				Bean: beans.Bean{
					Id:     tools.Ptr.Uid(),
					Status: tools.Ptr.String("1"),
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
		return nil
	})

	return err == nil, err
}

func (r *queryResolver) Roles(ctx context.Context, query source.QRole) (*source.Roles, error) {
	roles := make([]beans.Role, 0)
	total, err := r.DBS.SF(`
		select * from role u
		where 1 = 1
		{{ if .keyword }} and u.name like concat("%",:keyword,"%") {{ end }}
		{{ if .status }} and u.status = :status {{ end }}
	`, map[string]interface{}{
		"keyword": query.Keyword,
		"status":  query.Status,
	}).Page2(query.Index, query.Size, query.Count, &roles)
	return &source.Roles{Data: roles, Total: &total}, err
}

func (r *queryResolver) Role(ctx context.Context, id string) (*beans.Role, error) {
	return r.RoleLoader(ctx, id)
}

func (r *queryResolver) RoleUsers(ctx context.Context, id string) ([]string, error) {
	res := make([]string, 0)
	err := r.DBS.SF(`select uid from role_user where role = :id`, map[string]interface{}{"id": id}).Find(&res)
	return res, err
}

func (r *queryResolver) RolePerms(ctx context.Context, id string, typeArg *string) ([]string, error) {
	result := make([]string, 0)
	err := r.DBS.SF(`
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
	err := r.DBS.SF(`
		select object,mask from perm_object where role = :role 
	`, map[string]interface{}{
		"role": id,
	}).Find(&result)
	return result, err
}
