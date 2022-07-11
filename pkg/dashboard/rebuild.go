package dashboard

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/marc-campbell/nicedishy/pkg/logger"
	"github.com/marc-campbell/nicedishy/pkg/persistence"
)

func RebuildAllDashboards(ctx context.Context) error {
	pg := persistence.MustGetPGSession()

	query := `select id from dishy`
	rows, err := pg.Query(ctx, query)
	if err != nil {
		return fmt.Errorf("error querying for dishy: %v", err)
	}
	defer rows.Close()

	dishyIDs := []string{}
	for rows.Next() {
		var dishyID string
		err := rows.Scan(&dishyID)
		if err != nil {
			return fmt.Errorf("error scanning for dishy: %v", err)
		}

		dishyIDs = append(dishyIDs, dishyID)
	}
	rows.Close()

	for _, dishyID := range dishyIDs {
		logger.Infof("rebuilding dashboard for dishy %s", dishyID)

		uri := fmt.Sprintf("%s/api/dishy/%s/dashboard", os.Getenv("WEB_INTERNAL_ENDPOINT"), dishyID)
		fmt.Printf("uri: %s\n", uri)
		req, err := http.NewRequest("PUT", uri, nil)
		if err != nil {
			return fmt.Errorf("error creating request: %v", err)
		}

		req.Header.Add("Authorization", os.Getenv("INTERNAL_AUTH_TOKEN"))

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return fmt.Errorf("error sending request: %v", err)
		}

		if resp.StatusCode != http.StatusNoContent {
			return fmt.Errorf("error response: %v", resp.StatusCode)
		}
	}

	return nil
}
