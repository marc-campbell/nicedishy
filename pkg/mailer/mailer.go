package mailer

import (
	"context"
	"os"

	"github.com/keighl/postmark"
)

func sendEmail(ctx context.Context, email postmark.Email) error {
	client := postmark.NewClient(os.Getenv("POSTMARK_SERVER_TOKEN"), os.Getenv("POSTMARK_ACCOUNT_TOKEN"))

	_, err := client.SendEmail(email)
	if err != nil {
		return err
	}

	return nil
}

func SendSoftwareVersionChanged(ctx context.Context, emailAddress string, version string) error {
	email := postmark.Email{
		From:     "no-reply@nicedishy.com",
		To:       emailAddress,
		Subject:  "Woot. You have some new bits on your dishy to try out",
		HtmlBody: "...",
		TextBody: "There is a new version of firmware on your dishy",
	}

	return sendEmail(ctx, email)
}
