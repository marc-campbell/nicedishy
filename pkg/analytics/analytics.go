package analytics

import (
	"fmt"
	"os"

	"github.com/marc-campbell/nicedishy/pkg/logger"
	posthog "github.com/posthog/posthog-go"
)

var (
	posthogClient *posthog.Client
)

func init() {
	client, err := posthog.NewWithConfig(os.Getenv("POSTHOG_KEY"), posthog.Config{Endpoint: "https://app.posthog.com"})
	if err != nil {
		logger.Error(err)
		return
	}
	posthogClient = &client
}

func IdentifyUser(userID string, email string) error {
	if posthogClient == nil {
		return nil
	}

	c := *posthogClient
	err := c.Enqueue(posthog.Identify{
		DistinctId: fmt.Sprintf("user:%s", userID),
		Properties: posthog.NewProperties().
			Set("email", email),
	})
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func TrackUserEvent(userID string, event string) error {
	if posthogClient == nil {
		return nil
	}

	c := *posthogClient
	err := c.Enqueue(posthog.Capture{
		DistinctId: fmt.Sprintf("user:%s", userID),
		Event:      event,
	})
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}
