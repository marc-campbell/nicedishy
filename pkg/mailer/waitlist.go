package mailer

import (
	"context"
	"fmt"

	"github.com/keighl/postmark"
)

func SendInternalWaitlistSignup(ctx context.Context, emailAddress string) error {
	email := postmark.Email{
		From:     "system-alerts@nicedishy.com",
		To:       "marc@nicedishy.com",
		Subject:  "Someone just signed up for the NiceDishy waitlist",
		TextBody: fmt.Sprintf("%s just signed up for the NiceDishy waitlist", emailAddress),
	}

	return sendEmail(ctx, email)
}
