package mailer

import (
	"context"
	"os"
	"strings"

	"github.com/hanzoai/gochimp3"
)

func SignUpForNewsletter(ctx context.Context, emailAddress string, requireOptIn bool) error {
	client := gochimp3.New(os.Getenv("MAILCHIMP_API_KEY"))
	listID := "f8c85f2a58"

	list, err := client.GetList(listID, nil)
	if err != nil {
		return err
	}

	status := "pending"
	if !requireOptIn {
		status = "subscribed"
	}

	// Add subscriber
	req := &gochimp3.MemberRequest{
		EmailAddress: emailAddress,
		Status:       status,
	}

	_, err = list.CreateMember(req)
	if err != nil {
		if strings.Contains(err.Error(), "is already a list member") {
			_, err := list.UpdateMember(emailAddress, &gochimp3.MemberRequest{
				Status: status, // TODO this could move someone from subscribed to pending
			})
			if err != nil {
				return err
			}
			return nil
		}

		return err
	}

	return nil
}
