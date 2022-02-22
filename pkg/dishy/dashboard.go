package dishy

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// UpdateGrafanaDashboard will update the dashboard for a dishy
func UpdateGrafanaDashboard(ctx context.Context, id string, name string) error {
	folderID, err := getFolder(ctx, id)
	if err != nil {
		return fmt.Errorf("get folder: %v", err)
	}

	marshaledDashboard := fmt.Sprintf(defaultDashboard,
		id,
		"Default Dashboard",
		os.Getenv("GRAFANA_DATASOURCE_UID"), id,
		os.Getenv("GRAFANA_DATASOURCE_UID"), id,
		os.Getenv("GRAFANA_DATASOURCE_UID"), id,
		os.Getenv("GRAFANA_DATASOURCE_UID"), id,
		os.Getenv("GRAFANA_DATASOURCE_UID"), id,
		os.Getenv("GRAFANA_DATASOURCE_UID"), id,
		os.Getenv("GRAFANA_DATASOURCE_UID"), id,
		os.Getenv("GRAFANA_DATASOURCE_UID"), id,
		os.Getenv("GRAFANA_DATASOURCE_UID"), id,
		os.Getenv("GRAFANA_DATASOURCE_UID"), id,
		os.Getenv("GRAFANA_DATASOURCE_UID"), id,
		os.Getenv("GRAFANA_DATASOURCE_UID"), id)

	updateDashboardRequest := map[string]interface{}{
		"dashboard": "__REPLACEME__",
		"folderId":  folderID,
		"overwrite": true,
	}

	requestBody, err := json.Marshal(updateDashboardRequest)
	if err != nil {
		return fmt.Errorf("marshal request: %w", err)
	}

	fullRequestBody := strings.ReplaceAll(string(requestBody), `"__REPLACEME__"`, marshaledDashboard)

	req, err := http.NewRequest("POST", "http://grafana:3000/api/dashboards/db", bytes.NewBuffer([]byte(fullRequestBody)))
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

	return nil
}

// CreateGrafanaDashboard will create the default dashboard in grafana for this device
// and return the ID of the grafana dashboard
func CreateGrafanaDashboard(ctx context.Context, id string, name string) error {
	// create a folder for this device
	folderID, err := createFolder(ctx, id)
	if err != nil {
		return fmt.Errorf("create folder: %v", err)
	}

	marshaledDashboard := fmt.Sprintf(defaultDashboard,
		id,
		"Default Dashboard",
		os.Getenv("GRAFANA_DATASOURCE_UID"), id,
		os.Getenv("GRAFANA_DATASOURCE_UID"), id,
		os.Getenv("GRAFANA_DATASOURCE_UID"), id,
		os.Getenv("GRAFANA_DATASOURCE_UID"), id,
		os.Getenv("GRAFANA_DATASOURCE_UID"), id,
		os.Getenv("GRAFANA_DATASOURCE_UID"), id,
		os.Getenv("GRAFANA_DATASOURCE_UID"), id,
		os.Getenv("GRAFANA_DATASOURCE_UID"), id,
		os.Getenv("GRAFANA_DATASOURCE_UID"), id,
		os.Getenv("GRAFANA_DATASOURCE_UID"), id,
		os.Getenv("GRAFANA_DATASOURCE_UID"), id,
		os.Getenv("GRAFANA_DATASOURCE_UID"), id)

	createDashboardRequest := map[string]interface{}{
		"dashboard": "__REPLACEME__",
		"folderId":  folderID,
		"overwrite": false,
	}

	requestBody, err := json.Marshal(createDashboardRequest)
	if err != nil {
		return fmt.Errorf("marshal request: %w", err)
	}

	fullRequestBody := strings.ReplaceAll(string(requestBody), `"__REPLACEME__"`, marshaledDashboard)

	req, err := http.NewRequest("POST", "http://grafana:3000/api/dashboards/db", bytes.NewBuffer([]byte(fullRequestBody)))
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

	return nil
}

func getFolder(ctx context.Context, folderName string) (int, error) {
	req, err := http.NewRequest("GET", "http://grafana:3000/api/folders", nil)
	if err != nil {
		return 0, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("GRAFANA_API_KEY"))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, fmt.Errorf("do request: %w", err)
	}

	body, _ := ioutil.ReadAll(resp.Body)

	listFoldersResponse := []map[string]interface{}{}
	if err := json.Unmarshal(body, &listFoldersResponse); err != nil {
		return 0, fmt.Errorf("unmarshal response: %w", err)
	}

	for _, folder := range listFoldersResponse {
		if folder["title"].(string) == folderName {
			return int(folder["id"].(float64)), nil
		}
	}

	return 0, fmt.Errorf("folder not found")
}

func createFolder(ctx context.Context, folderName string) (int, error) {
	createFolderRequest := map[string]interface{}{
		"title": folderName,
	}
	requestBody, err := json.Marshal(createFolderRequest)
	if err != nil {
		return 0, fmt.Errorf("marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", "http://grafana:3000/api/folders", bytes.NewBuffer(requestBody))
	if err != nil {
		return 0, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("GRAFANA_API_KEY"))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	createFolderResponse := map[string]interface{}{}
	if err := json.Unmarshal(body, &createFolderResponse); err != nil {
		return 0, fmt.Errorf("unmarshal response: %w", err)
	}

	folderID := createFolderResponse["id"].(float64)
	return int(folderID), nil
}
