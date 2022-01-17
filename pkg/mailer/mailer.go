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

func sendTemplatedEmail(ctx context.Context, email postmark.TemplatedEmail) error {
	client := postmark.NewClient(os.Getenv("POSTMARK_SERVER_TOKEN"), os.Getenv("POSTMARK_ACCOUNT_TOKEN"))

	_, err := client.SendTemplatedEmail(email)
	if err != nil {
		return err
	}

	return nil
}
