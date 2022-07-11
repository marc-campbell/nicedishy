package mailer

import (
	"context"
	"time"

	"github.com/marc-campbell/nicedishy/pkg/logger"
	"github.com/marc-campbell/nicedishy/pkg/persistence"
	"github.com/marc-campbell/nicedishy/pkg/stores"
)

const DeviceDisconnectedTemplateID = int64(27137544)

// QueueDisconnectedDeviceEmails queues emails for disconnected devices
func QueueDisconnectedDeviceEmails() {
	for {
		pg := persistence.MustGetPGSession()

		now := time.Now()
		query := `delete from dishy_disconnected_queue where send_at < $1 returning dishy_id`
		rows, err := pg.Query(context.Background(), query, now)
		if err != nil {
			logger.Error(err)
			panic(err)
		}
		defer rows.Close()

		dishyIDs := []string{}
		for rows.Next() {
			var dishyID string
			if err := rows.Scan(&dishyID); err != nil {
				logger.Error(err)
				panic(err)
			}
			dishyIDs = append(dishyIDs, dishyID)
		}
		rows.Close()

		for _, dishyID := range dishyIDs {
			query = `select gu.email_address, d.name, d.last_metric_at from google_user gu
		inner join dishy d on d.user_id = gu.id
		where d.id = $1`
			row := pg.QueryRow(context.Background(), query, dishyID)

			var email, dishyName string
			var lastMetricAt time.Time
			if err := row.Scan(&email, &dishyName, &lastMetricAt); err != nil {
				logger.Error(err)
				continue
				// panic(err)
			}

			model := map[string]interface{}{
				"dishyName":    dishyName,
				"lastMetricAt": lastMetricAt.Format(time.RFC1123),
			}
			_, err = stores.GetStore().QueueEmail(context.TODO(), "notifications@nicedishy.com", email, DeviceDisconnectedTemplateID, model)
			if err != nil {
				logger.Error(err)
				panic(err)
			}
		}

		time.Sleep(time.Second * 15)
	}
}
