package mailer

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/keighl/postmark"
)

const (
	SendSoftwareVersionChangedTemplateID = 27132960
	WeeklyReportTemplateID               = 28522168
)

type SoftwareVersionChangedContext struct {
	NewFirmware      string `json:"newFirmware"`
	PreviousFirmware string `json:"previousFirmware"`
}

func GetSoftwareVersionChangedModel(ctx SoftwareVersionChangedContext) map[string]interface{} {
	return map[string]interface{}{
		"newFirmware":      ctx.NewFirmware,
		"previousFirmware": ctx.PreviousFirmware,
	}
}

func SendSoftwareVersionChanged(ctx context.Context, emailAddress string, mailContext SoftwareVersionChangedContext) error {
	model, err := json.Marshal(mailContext)
	if err != nil {
		return fmt.Errorf("marshal mail context: %v", err)
	}
	marsheledModel := map[string]interface{}{}
	if err := json.Unmarshal(model, &marsheledModel); err != nil {
		return fmt.Errorf("unmarshal mail context: %v", err)
	}

	email := postmark.TemplatedEmail{
		TemplateId:    SendSoftwareVersionChangedTemplateID,
		From:          "notifications@nicedishy.com",
		To:            emailAddress,
		TemplateModel: marsheledModel,
	}

	_, err = sendTemplatedEmail(ctx, email)
	if err != nil {
		return fmt.Errorf("send templated email: %v", err)
	}

	return nil
}
