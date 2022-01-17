package mailer

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/keighl/postmark"
	"github.com/marc-campbell/nicedishy/pkg/logger"
	"github.com/marc-campbell/nicedishy/pkg/stores"
)

// StartDaemon will watch for need-to-send emails and
//  handle delivery
func StartDaemon() {
	ctx := context.Background()

	logger.Info("starting email daemon")
	for {
		queuedEmails, err := stores.GetStore().GetQueuedEmails(context.Background())
		if err != nil {
			fmt.Printf("failed to get queued emails: %v\n", err)
			time.Sleep(time.Second * 15)
			continue
		}

		for _, queuedEmail := range queuedEmails {
			templateID, err := strconv.ParseInt(queuedEmail.TemplateID, 10, 64)
			if err != nil {
				fmt.Printf("failed to parse template id: %v\n", err)
				if err := stores.GetStore().MarkQueuedEmailError(ctx, queuedEmail.ID); err != nil {
					fmt.Printf("failed to mark queued email as error: %v\n", err)
				}
			}

			logger.Debugf("sending email from %s to %s", queuedEmail.FromAddress, queuedEmail.ToAddress)

			postmarkTemplatedEmail := postmark.TemplatedEmail{
				TemplateId:    templateID,
				TemplateModel: queuedEmail.TemplateContext,
				From:          queuedEmail.FromAddress,
				To:            queuedEmail.ToAddress,
			}

			postmarkResponse, err := sendTemplatedEmail(ctx, postmarkTemplatedEmail)
			fmt.Printf("%#v\n", postmarkResponse)
			if err != nil {
				fmt.Printf("failed to send templated email: %v\n", err)

				// mark as error
				if err := stores.GetStore().MarkQueuedEmailError(ctx, queuedEmail.ID); err != nil {
					fmt.Printf("failed to mark queued email as error: %v\n", err)
				}
			} else {
				if err := stores.GetStore().MarkQueuedEmailSent(ctx, queuedEmail.ID); err != nil {
					fmt.Printf("failed to mark queued email as sent: %v\n", err)
				}
			}
		}

		time.Sleep(time.Second * 15)
	}
}
