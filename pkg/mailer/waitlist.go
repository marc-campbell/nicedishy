package mailer

import (
	"context"
	"fmt"

	"github.com/keighl/postmark"
)

func SendWelcomeWaitlist(ctx context.Context, emailAddress string) error {
	email := postmark.Email{
		From:     "notifications@nicedishy.com",
		To:       emailAddress,
		Subject:  "You're on the HiceDishy list!",
		TextBody: `You are on the NiceDishy list. Stay tuned and we'll email you as soon as possible when it's time to connect your device!`,
	}

	return sendEmail(ctx, email)
}

func SendInternalWaitlistSignup(ctx context.Context, emailAddress string) error {
	email := postmark.Email{
		From:     "system-alerts@nicedishy.com",
		To:       "marc@nicedishy.com",
		Subject:  "Someone just signed up for the NiceDishy waitlist",
		TextBody: fmt.Sprintf("%s just signed up for the NiceDishy waitlist", emailAddress),
	}

	return sendEmail(ctx, email)
}
