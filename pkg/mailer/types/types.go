package types

import "time"

type Email struct {
	ID              string
	FromAddress     string
	ToAddress       string
	TemplateID      string
	TemplateContext map[string]interface{}
	QueuedAt        time.Time
	SentAt          *time.Time
}
