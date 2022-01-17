package mailer

import (
	"context"
	"os"

	"github.com/hanzoai/gochimp3"
)

func SignUpForNewsletter(ctx context.Context, emailAddress string) error {
	client := gochimp3.New(os.Getenv("MAILCHIMP_API_KEY"))
	listID := "f8c85f2a58"

	list, err := client.GetList(listID, nil)
	if err != nil {
		return err
	}

	// Add subscriber
	req := &gochimp3.MemberRequest{
		EmailAddress: emailAddress,
		Status:       "subscribed",
	}

	if _, err := list.CreateMember(req); err != nil {
		return err
	}

	return nil
}
