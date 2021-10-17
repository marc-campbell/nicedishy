package generator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/marc-campbell/nicedishy/pkg/handlers"
)

func GenerateAndSendData(token string, timestamp time.Time, uptimeSeconds int) error {
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
			DeviceState: handlers.StoreDataStatusDeviceStateRequest{
				UptimeSeconds: uptimeSeconds,
			},
			State:                 "connected",
			SNR:                   0.0 + rand.Float64()*(9.0-0.0),
			DownlinkThroughputBps: 5000.0 + rand.Float64()*(7000.0-5000.0),
			UplinkThroughputBps:   700.0 + rand.Float64()*(1200.0-700.0),
			PopPingLatencyMs:      20.0 + rand.Float64()*(90.0-20.0),
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
