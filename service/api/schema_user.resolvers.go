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
	id, err := r.Create(ctx, new(beans.User), input.User)
	if err != nil {
		return "", err
	}

	input.Account.Default = tools.Ptr.Int(1)
	_, err = r.AccountCreate(ctx, *input.Account)
	if err != nil {
		return "", err
	}

	//// 用户新增推送
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

	// 更新用户
	ok, err := r.Update(ctx, new(beans.User), id, input.User)
	if err != nil {
		return false, err
	}

	// 更新账户
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
	ok, err := r.Removes(ctx, new(beans.User), ids)
	if err != nil || !ok {
		return false, err
	}

	// 删除用户下所有的账户
	err = r.DBS.TS(func(sess *txorm.Session) error {
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
			{{ if .role }} 
				join role_user ru on ru.uid = u.id and ru.role = :role
			{{ end }}
		where 1 = 1
		{{ if .keyword }} and u.name like concat("%",:keyword,"%") {{ end }}
		{{ if .status }} and u.status = :status {{ end }}
		
	`, map[string]interface{}{
		"keyword": query.Keyword,
		"ctx":     ctx,
		"role":    query.Role,
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
