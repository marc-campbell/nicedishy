package user

import (
	"context"

	"github.com/marc-campbell/nicedishy/pkg/dishy"
	"github.com/marc-campbell/nicedishy/pkg/stores"
	"github.com/marc-campbell/nicedishy/pkg/user/types"
	"github.com/pkg/errors"
)

func GetOrCreate(ctx context.Context, email string, avatarURL string) (*types.User, error) {
	existingUser, err := stores.GetStore().GetUserByEmail(ctx, email)
	if stores.GetStore().IsNotFound(err) {
		createdUser, err := stores.GetStore().CreateUser(ctx, email, avatarURL)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create user")
		}

		if _, err := dishy.CreateDefaultForUser(ctx, createdUser.ID); err != nil {
			return nil, errors.Wrap(err, "failed to create default for user")
		}

		return createdUser, nil
	} else if err != nil {
		return nil, errors.Wrap(err, "failed to get user by email")
	}

	return existingUser, nil
}
