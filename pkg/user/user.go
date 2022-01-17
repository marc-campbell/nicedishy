package user

import (
	"context"

	"github.com/marc-campbell/nicedishy/pkg/mailer"
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

		if createdUser.IsWaitlisted {
			// send email to customer
			_, err := stores.GetStore().QueueEmail(ctx, "marc@nicedishy.com", createdUser.EmailAddress, "26743037", map[string]interface{}{})
			if err != nil {
				return nil, errors.Wrap(err, "failed to queue email")
			}

			_,, err := mailer.SignUpForNewsletter(ctx, createdUser.EmailAddress); err != nil {
				return nil, errors.Wrap(err, "failed to sign up for newsletter")
			}

			if err := mailer.SendInternalWaitlistSignup(ctx, createdUser.EmailAddress); err != nil {
				return nil, errors.Wrap(err, "failed to send waitlist signup email")
			}
		}

		return createdUser, nil
	} else if err != nil {
		return nil, errors.Wrap(err, "failed to get user by email")
	}

	return existingUser, nil
}
