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
		DistinctId: fmt.Sprintf("user:%s", userID, email),
		Properties: posthog.NewProperties().
			Set("email", email),
	})
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func IdentifyDishy(dishyID string, name string) error {
	if posthogClient == nil {
		return nil
	}

	c := *posthogClient
	err := c.Enqueue(posthog.Identify{
		DistinctId: fmt.Sprintf("dishy:%s", dishyID),
		Properties: posthog.NewProperties().
			Set("name", name),
	})
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func TrackDishyEvent(dishyID string, event string) error {
	if posthogClient == nil {
		return nil
	}

	c := *posthogClient
	err := c.Enqueue(posthog.Capture{
		DistinctId: fmt.Sprintf("dishy:%s", dishyID),
		Event:      event,
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
