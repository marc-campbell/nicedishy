package dishy

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// CreateGrafanaDashboard will create the default dashboard in grafana for this device
// and return the ID of the grafana dashboard
func CreateGrafanaDashboard(ctx context.Context, id string, name string) error {
	marshaledDashboard := fmt.Sprintf(defaultDashboard,
		id,
		name,
		os.Getenv("GRAFANA_DATASOURCE_UID"))

	req, err := http.NewRequest("POST", "http://grafana:3000/api/dashboards/db", bytes.NewBuffer([]byte(marshaledDashboard)))
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("GRAFANA_API_KEY"))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("do request: %w", err)
	}

	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return nil
}
