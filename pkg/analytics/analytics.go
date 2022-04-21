package analytics

import (
	"log"
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
		log.Fatal(err)
	}
	defer client.Close()
	posthogClient = &client
}

func TrackEvent(actorID string, event string) error {
	if posthogClient == nil {
		return nil
	}

	c := *posthogClient
	err := c.Enqueue(posthog.Capture{
		DistinctId: actorID,
		Event:      event,
	})
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}
