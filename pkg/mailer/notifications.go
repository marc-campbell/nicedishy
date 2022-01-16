package mailer

import (
	"context"

	"github.com/keighl/postmark"
)

func SendSoftwareVersionChanged(ctx context.Context, emailAddress string, version string) error {
	email := postmark.Email{
		From:     "notifications@nicedishy.com",
		To:       emailAddress,
		Subject:  "Woot. You have some new bits on your dishy to try out",
		HtmlBody: "...",
		TextBody: "There is a new version of firmware on your dishy",
	}

	return sendEmail(ctx, email)
}
