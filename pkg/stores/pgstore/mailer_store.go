package pgstore

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	mailertypes "github.com/marc-campbell/nicedishy/pkg/mailer/types"
	"github.com/marc-campbell/nicedishy/pkg/persistence"
	"github.com/segmentio/ksuid"
)

func (s PGStore) GetQueuedEmails(ctx context.Context) ([]*mailertypes.Email, error) {
	pg := persistence.MustGetPGSession()

	query := `select id, queued_at, sent_at, from_address,
to_address, template_id, marshalled_context
from email_notification where sent_at is null and error_at is null order by queued_at asc`
	rows, err := pg.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("getting queued emails: %w", err)
	}

	emails := []*mailertypes.Email{}
	for rows.Next() {
		email := mailertypes.Email{}
		sentAt := sql.NullTime{}
		marshalledContext := ""

		if err := rows.Scan(&email.ID, &email.QueuedAt, &sentAt, &email.FromAddress, &email.ToAddress, &email.TemplateID, &marshalledContext); err != nil {
			return nil, fmt.Errorf("scanning email: %w", err)
		}

		if sentAt.Valid {
			email.SentAt = &sentAt.Time
		}

		templateContext := map[string]interface{}{}
		if err := json.Unmarshal([]byte(marshalledContext), &templateContext); err != nil {
			return nil, fmt.Errorf("unmarshalling email context: %w", err)
		}
		email.TemplateContext = templateContext

		emails = append(emails, &email)
	}

	return emails, nil
}

func (s PGStore) MarkQueuedEmailSent(ctx context.Context, id string) error {
	pg := persistence.MustGetPGSession()

	query := `update email_notification set sent_at = now() where id = $1`
	if _, err := pg.Exec(ctx, query, id); err != nil {
		return fmt.Errorf("marking email as sent: %w", err)
	}

	return nil
}

func (s PGStore) MarkQueuedEmailError(ctx context.Context, id string) error {
	pg := persistence.MustGetPGSession()

	query := `update email_notification set error_at = now() where id = $1`
	if _, err := pg.Exec(ctx, query, id); err != nil {
		return fmt.Errorf("marking email as error: %w", err)
	}

	return nil
}

func (s PGStore) QueueEmail(ctx context.Context, fromAddress string, toAddress string, templateID string, templateContext map[string]interface{}) (*mailertypes.Email, error) {
	marshalledContext, err := json.Marshal(templateContext)
	if err != nil {
		return nil, fmt.Errorf("marshalling email context: %w", err)
	}

	id, err := ksuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("generating email id: %w", err)
	}

	queuedAt := time.Now()

	pg := persistence.MustGetPGSession()
	query := `insert into email_notification
(id, queued_at, sent_at, error_at, from_address, to_address, template_id, marshalled_context)
values
($1, $2, null, null, $3, $4, $5, $6)`

	if _, err := pg.Exec(ctx, query, id.String(), queuedAt, fromAddress, toAddress, templateID, marshalledContext); err != nil {
		return nil, fmt.Errorf("queuing email: %w", err)
	}

	return &mailertypes.Email{
		ID:              id.String(),
		FromAddress:     fromAddress,
		ToAddress:       toAddress,
		TemplateID:      templateID,
		TemplateContext: templateContext,
		QueuedAt:        queuedAt,
	}, nil
}
