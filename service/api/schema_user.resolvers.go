package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/service/api/source"
	"github.com/zhanghup/go-app/service/event"
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/database/txorm"
	"github.com/zhanghup/go-tools/tog"
)

func (r *mutationResolver) UserCreate(ctx context.Context, input source.NewUser) (string, error) {
	sess := r.DBS.NewSession(ctx)
	user := new(beans.User)
	if input.User["name"] != nil {
		name := input.User["name"].(string)
		user.Py = tools.Ptr.String(tools.Pin.Py(name))
		user.Py = tools.Ptr.String(tools.Pin.Pinyin(name))
	}
	id, err := r.Create(sess.Context(), user, input.User)
	if err != nil {
		return "", err
	}

	input.Account.Default = tools.Ptr.Int(1)
	input.Account.UID = &id
	_, err = r.AccountCreate(sess.Context(), *input.Account)
	if err != nil {
		return "", err
	}

	// 用户新增推送
	go func() {
		user, err := r.UserLoader(ctx, id)
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
	return id, err
}

func (r *mutationResolver) UserUpdate(ctx context.Context, id string, input source.UpdUser) (bool, error) {
	// 读取库存用户
	user, err := r.UserLoader(ctx, id)
	if err != nil {
		return false, err
	}
	if user == nil {
		return false, errors.New("用户不存在")
	}
	sess := r.DBS.NewSession(ctx)

	// 更新用户
	upduser := beans.User{}
	if input.User["name"] != nil {
		name := input.User["name"].(string)
		upduser.Py = tools.Ptr.String(tools.Pin.Py(name))
		upduser.Pinyin = tools.Ptr.String(tools.Pin.Pinyin(name))
	}
	ok, err := r.Update(sess.Context(), &upduser, id, input.User)
	if err != nil {
		return false, err
	}

	// 更新账户
	acc, err := r.AccountDefaultLoader(ctx, id)
	if err != nil {
		return false, err
	}
	if acc == nil {
		_, err := r.AccountCreate(sess.Context(), source.NewAccount{
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
		_, err = r.AccountUpdate(sess.Context(), *acc.Id, *input.Account)
		if err != nil {
			return false, err
		}
	}

	// 用户更新推送
	go func() {
		event.UserUpdate(*user)
	}()

	return ok, nil
}

func (r *mutationResolver) UserRemoves(ctx context.Context, ids []string) (bool, error) {
	if tools.Str.Contains(ids, "root") {
		return false, errors.New("root用户无法删除")
	}
	users := make([]beans.User, 0)
	err := r.DB.In("id", ids).Find(&users)
	if err != nil {
		return false, err
	}
	sess := r.DBS.NewSession(ctx)
	ok, err := r.Removes(sess.Context(), new(beans.User), ids)
	if err != nil || !ok {
		return false, err
	}

	// 删除用户下所有的账户
	err = sess.TS(func(sess *txorm.Session) error {
		return sess.SF(`delete from account where uid in :ids`, map[string]interface{}{"ids": ids}).Exec()
	})
	if err != nil {
		return false, err
	}

	go func() {
		for _, user := range users {
			event.UserRemove(user)
		}
	}()

	return true, nil
}

func (r *queryResolver) Users(ctx context.Context, query source.QUser) (*source.Users, error) {
	users := make([]beans.User, 0)
	total, err := r.DBS.SF(`
		select 
			u.* 
		from 
			user u
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
		
	`, map[string]interface{}{
		"keyword": query.Keyword,
		"ctx":     ctx,
		"status":  query.Status,
	}).Page2(query.Index, query.Size, query.Count, &users)
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

// User returns source.UserResolver implementation.
func (r *Resolver) User() source.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
