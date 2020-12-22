package api

import (
	"context"
	"github.com/zhanghup/go-app/beans"
	"github.com/zhanghup/go-app/service/api/source"
	"github.com/zhanghup/go-tools"
)

func (r *mutationResolver) MenuUpdate(ctx context.Context, id string, input source.UpdMenu) (bool, error) {
	return r.Update(ctx, &beans.Menu{}, id, input)
}

func (r *mutationResolver) MenuReload(ctx context.Context, menus []source.MenuLocal) (bool, error) {
	err := r.Sess(ctx).SF("delete from menu").Exec()
	if err != nil {
		return false, err
	}

	insert := func(pid *string, m source.MenuLocal) error {
		err := r.Sess(ctx).Insert(beans.Menu{
			Bean: beans.Bean{
				Id:     m.ID,
				Status: tools.Ptr.String("1"),
			},
			Name:  m.Name,
			Title: m.Title,
			Path:  m.Path,
			Alias: m.Alias,
			Icon:  m.Icon,
			Parent: pid,
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
	err := r.DBS().SF(`
		select 
			p.* 
		from 
			menu p 
		where 1 = 1 
			{{ if .status }} and p.status = :status {{ end }}
		`,
		map[string]interface{}{
			"status": query.Status,
		}).Find(&plans)
	return plans, err
}

func (r *queryResolver) Menu(ctx context.Context, id string) (*beans.Menu, error) {
	return r.Resolver.MenuLoader(ctx, id)
}
