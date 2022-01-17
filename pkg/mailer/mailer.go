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

func sendTemplatedEmail(ctx context.Context, email postmark.TemplatedEmail) (*postmark.EmailResponse, error) {
	client := postmark.NewClient(os.Getenv("POSTMARK_SERVER_TOKEN"), os.Getenv("POSTMARK_ACCOUNT_TOKEN"))

	// postmark doesn't accept empty (but valid) model context
	if len(email.TemplateModel) == 0 {
		email.TemplateModel = map[string]interface{}{
			"placedolder": "something",
		}
	}

	resp, err := client.SendTemplatedEmail(email)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
