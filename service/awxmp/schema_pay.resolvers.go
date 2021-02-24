package awxmp

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/zhanghup/go-app/gs"
)

func (r *mutationResolver) PayError(ctx context.Context, id string, typeArg string) (bool, error) {
	err := gs.Sess(ctx).SF(`
		update 
			wxmp_order 
		set 
			state = 4,
			updated = unix_timestamp(now()) 
		where oid = ? and otype = ? 
	`, id, typeArg).Exec()
	if err != nil {
		return false, err
	}
	return r.PayCancelAction(ctx, id, typeArg)
}

func (r *mutationResolver) PayCancel(ctx context.Context, id string, typeArg string) (bool, error) {
	err := gs.Sess(ctx).SF(`
		update 
			wxmp_order 
		set 
			state = 2,
			updated = unix_timestamp(now()) 
		where oid = ? and otype = ? 
	`, id, typeArg).Exec()
	if err != nil {
		return false, err
	}
	return r.PayCancelAction(ctx, id, typeArg)
}

func (r *mutationResolver) PaySuccess(ctx context.Context, id string, typeArg string) (bool, error) {
	err := gs.Sess(ctx).SF(`
		update 
			wxmp_order 
		set 
			state = 1,
			updated = unix_timestamp(now()) 
		where oid = ? and otype = ? 
	`, id, typeArg).Exec()
	return err == nil, err
}
