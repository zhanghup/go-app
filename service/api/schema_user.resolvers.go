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
	"github.com/zhanghup/go-tools/tog"
)

func (r *mutationResolver) UserCreate(ctx context.Context, input map[string]interface{}) (string, error) {
	id, err := r.Create(ctx, new(beans.User), input)
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

func (r *mutationResolver) UserUpdate(ctx context.Context, id string, input map[string]interface{}) (bool, error) {
	user, err := r.UserLoader(ctx, id)
	if err != nil {
		return false, err
	}
	if user == nil {
		return false, errors.New("用户不存在")
	}

	ok, err := r.Update(ctx, new(beans.User), id, input)
	if err != nil {
		return false, err
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
		select * from user u
		where 1 = 1
		{{ if .keyword }} and u.name like concat("%",:keyword,"%") {{ end }}
	`, map[string]interface{}{"keyword": query.Keyword}).Page2(query.Index, query.Size, query.Count, &users)
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

// User returns source.UserResolver implementation.
func (r *Resolver) User() source.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
