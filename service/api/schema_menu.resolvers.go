package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/zhanghup/go-app/gs"

	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/service/api/source"
	"github.com/zhanghup/go-tools"
)

func (r *mutationResolver) MenuCreate(ctx context.Context, input source.NewMenu) (string, error) {
	return r.Create(ctx, &beans.Menu{}, input)
}

func (r *mutationResolver) MenuUpdate(ctx context.Context, id string, input source.UpdMenu) (bool, error) {
	return r.Update(ctx, &beans.Menu{}, id, input)
}

func (r *mutationResolver) MenuReload(ctx context.Context, menus []source.MenuLocal) (bool, error) {
	err := gs.Sess(ctx).SF("delete from menu").Exec()
	if err != nil {
		return false, err
	}

	weight := 0
	insert := func(pid *string, m source.MenuLocal) error {
		weight += 1
		err := gs.Sess(ctx).Insert(beans.Menu{
			Bean: beans.Bean{
				Id:     m.ID,
				Status: tools.PtrOfString("1"),
				Weight: &weight,
			},
			Name:   m.Name,
			Title:  m.Title,
			Icon:   m.Icon,
			Parent: pid,
			Type:   m.Type,
		})
		return err
	}

	for _, m := range menus {
		err := insert(nil, m)
		if err != nil {
			return false, err
		}
		for _, mm := range m.Children {
			err := insert(m.ID, mm)
			if err != nil {
				return false, err
			}
			for _, mmm := range mm.Children {
				err := insert(mm.ID, mmm)
				if err != nil {
					return false, err
				}
				for _, mmmm := range mmm.Children {
					err := insert(mmm.ID, mmmm)
					if err != nil {
						return false, err
					}
				}
			}
		}
	}
	return true, nil
}

func (r *queryResolver) Menus(ctx context.Context, query source.QMenu) ([]beans.Menu, error) {
	plans := make([]beans.Menu, 0)
	err := gs.DBS().SF(`
		select 
			p.* 
		from 
			menu p 
		{{ if .no_admin }}
			join (
				select menu.id id from menu where type = '0'
				union 
				select perm.oid id from role_user  
				join perm on perm.role = role_user.role and role_user.uid = :uid and type = 'menu'
			) s on p.id = s.id
		{{ end }}
		where 1 = 1 
			{{ if .status }} and p.status = :status {{ end }}
		`,
		map[string]interface{}{
			"uid":      r.Me(ctx).Id,
			"no_admin": !r.Me(ctx).Admin,
			"status":   query.Status,
		}).Order("weight").Find(&plans)
	return plans, err
}

func (r *queryResolver) Menu(ctx context.Context, id string) (*beans.Menu, error) {
	return r.Resolver.MenuLoader(ctx, id)
}
