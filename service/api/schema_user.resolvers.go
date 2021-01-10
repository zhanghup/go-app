package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"github.com/zhanghup/go-tools/database/txorm"

	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/service/api/source"
	"github.com/zhanghup/go-app/service/event"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/tog"
)

func (r *mutationResolver) UserCreate(ctx context.Context, input source.NewUser) (string, error) {
	uid := ""

	{ // 添加用户
		user := new(beans.User)
		if input.User.Name != nil {
			name := *input.User.Name
			user.Py = tools.Ptr.String(tools.Pin.Py(name))
			user.Py = tools.Ptr.String(tools.Pin.Pinyin(name))
		}
		id, err := r.Create(ctx, user, input.User)
		if err != nil {
			return "", err
		}
		uid = id
	}

	{ // 添加账户
		input.Account.Default = tools.Ptr.Int(1)
		input.Account.UID = &uid
		_, err := r.AccountCreate(ctx, *input.Account)
		if err != nil {
			return "", err
		}
	}

	{ // 角色添加
		if len(input.Roles) > 0 {
			for _, str := range input.Roles {
				_, err := r.Create(ctx, &beans.RoleUser{}, map[string]interface{}{"role": str, "uid": uid})
				if err != nil {
					return "", err
				}
			}
		}
	}

	go func() { // 用户新增事件推送
		user, err := r.UserLoader(ctx, uid)
		if err != nil {
			tog.Error(err.Error())
			return
		}
		if user == nil {
			tog.Error("用户不存在")
			return
		}
		event.UserCreate(*user)
	}()
	return uid, nil
}

func (r *mutationResolver) UserUpdate(ctx context.Context, id string, input source.UpdUser) (bool, error) {
	var user *beans.User
	{ // 读取库存用户
		u, err := r.UserLoader(ctx, id)
		if err != nil {
			return false, err
		}
		if u == nil {
			return false, errors.New("用户不存在")
		}
		user = u
	}

	{ // 更新用户
		upduser := beans.User{}
		if input.User.Name != nil {
			name := *input.User.Name
			upduser.Py = tools.Ptr.String(tools.Pin.Py(name))
			upduser.Pinyin = tools.Ptr.String(tools.Pin.Pinyin(name))
		}
		ok, err := r.Update(ctx, &upduser, id, input.User)
		if err != nil {
			return false, err
		}
		if !ok {
			return false, errors.New("用户更新失败")
		}
	}

	{ // 更新账户
		acc, err := r.AccountDefaultLoader(ctx, id)
		if err != nil {
			return false, err
		}
		if acc == nil {
			_, err := r.AccountCreate(ctx, source.NewAccount{
				UID:      &id,
				Type:     input.Account.Type,
				Username: input.Account.Username,
				Password: input.Account.Password,
				Default:  tools.Ptr.Int(1),
			})
			if err != nil {
				return false, err
			}
		} else {
			input.Account.Default = tools.Ptr.Int(1)

			_, err = r.AccountUpdate(ctx, *acc.Id, *input.Account)
			if err != nil {
				return false, err
			}
		}
	}

	{ // 更新角色
		err := r.Sess(ctx).SF(`delete from role_user where uid = :uid`, map[string]interface{}{"uid": id}).Exec()
		if err != nil {
			return false, err
		}
		if len(input.Roles) > 0 {
			for _, str := range input.Roles {
				_, err := r.Create(ctx, &beans.RoleUser{}, map[string]interface{}{"role": str, "uid": id})
				if err != nil {
					return false, err
				}
			}
		}
	}

	// 用户更新推送
	go func() {
		event.UserUpdate(*user)
	}()

	return true, nil
}

func (r *mutationResolver) UserRemoves(ctx context.Context, ids []string) (bool, error) {
	sess := r.Sess(ctx)

	if tools.Str.Contains(ids, "root") {
		return false, errors.New("root用户无法删除")
	}

	users := make([]beans.User, 0)
	{ // 查找当前需要删除的用户
		err := r.DBS(ctx).E().In("id", ids).Find(&users)
		if err != nil {
			return false, err
		}
	}

	{ // 删除用户
		ok, err := r.Removes(ctx, new(beans.User), ids)
		if err != nil || !ok {
			return false, err
		}
	}

	{ // 删除用户下所有的账户
		err := sess.SF(`delete from account where uid in :ids`, map[string]interface{}{"ids": ids}).Exec()
		if err != nil {
			return false, err
		}
	}

	{ // 删除用户下所有的角色
		err := sess.SF(`delete from role_user where uid in :ids`, map[string]interface{}{"ids": ids}).Exec()
		if err != nil {
			return false, err
		}
	}

	go func() {
		for _, user := range users {
			event.UserRemove(user)
		}
	}()

	return true, nil
}

func (r *mutationResolver) UserWithRole(ctx context.Context, uid string, roles []string) (bool, error) {
	sess := r.Sess(ctx)
	err := sess.TS(func(sess txorm.ISession) error {
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

func (r *queryResolver) Users(ctx context.Context, query source.QUser) (*source.Users, error) {
	users := make([]beans.User, 0)
	total, err := r.DBS(ctx).SF(`
		select 
			u.* 
		from 
			user u
		join with_role_user on u.id = with_role_user.id
		where 1 = 1
		{{ if .keyword }} 
			and (
				1 = 0
				or u.name like concat('%',:keyword,'%')
				or u.py like concat('%',:keyword,'%')
				or u.pinyin like concat('%',:keyword,'%')
			) 
		{{ end }}
		{{ if .status }} and u.status = :status {{ end }}
		{{ if .dept }} and u.dept = :dept {{ end }}
		{{ if .type }} and u.type = :type {{ end }}
		{{ if .sn }} and u.sn = :sn {{ end }}
		{{ if .sex }} and u.sex = :sex {{ end }}
		{{ if .admin }} and u.admin = :admin {{ end }}
		{{ if .name }} and u.name like concat('%',:name,'%') {{ end }}
		
	`, map[string]interface{}{
		"keyword": query.Keyword,
		"status":  query.Status,
		"dept":    query.Dept,
		"type":    query.Type,
		"sn":      query.Sn,
		"sex":     query.Sex,
		"admin":   query.Admin,
		"name":    query.Name,
	}).With("with_role_user").Page2(query.Index, query.Size, query.Count, &users)
	return &source.Users{Data: users, Total: &total}, err
}

func (r *queryResolver) User(ctx context.Context, id string) (*beans.User, error) {
	return r.UserLoader(ctx, id)
}

func (r *userResolver) ODept(ctx context.Context, obj *beans.User) (*beans.Dept, error) {
	if obj.Dept == nil {
		return nil, nil
	}
	return r.DeptLoader(ctx, *obj.Dept)
}

func (r *userResolver) OAccount(ctx context.Context, obj *beans.User) (*beans.Account, error) {
	return r.AccountDefaultLoader(ctx, *obj.Id)
}

func (r *userResolver) ORoles(ctx context.Context, obj *beans.User) ([]beans.Role, error) {
	if obj.Id == nil {
		return nil, nil
	}
	c := make([]beans.Role, 0)
	err := r.Loader(ctx).Slice(struct {
		beans.Role `xorm:"extends"`
		Uid        *string `xorm:"uid"`
	}{}, "select role.*,role_user.uid from role_user join role on role.id = role_user.role where role_user.uid in :keys ", nil, "Uid", "Role").Load(*obj.Id, &c)
	return c, err
}

// User returns source.UserResolver implementation.
func (r *Resolver) User() source.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
