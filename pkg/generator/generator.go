package generator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/marc-campbell/nicedishy/pkg/handlers"
)

func GenerateAndSendData(token string, timestamp time.Time) error {
	hardwareVersion := "rev1_pre_production"
	softwareVersion := "98601479-46bd-4c5a-acbf-2d4839518ce2.uterm.release"
	endpoint := "http://localhost:30065"

	storeDataRequest := handlers.StoreDataRequest{
		When: timestamp.Format(time.RFC3339),
		Status: handlers.StoreDataStatusRequest{
			DeviceInfo: handlers.StoreDataStatusDeviceInfoRequest{
				HardwareVersion: hardwareVersion,
				SoftwareVersion: softwareVersion,
			},
		},
	}

	b, err := json.Marshal(storeDataRequest)
	if err != nil {
		return fmt.Errorf("error marshaling request: %w", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/stats", endpoint), bytes.NewBuffer(b))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Token %s", token))
	req.Header.Add("Content-type", "application/json")
	req.Header.Add("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error executing request: %w", err)
	}
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("unexpected response: %d", resp.StatusCode)
	}

	return nil
}
