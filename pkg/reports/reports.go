package reports

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/marc-campbell/nicedishy/pkg/logger"
	"github.com/marc-campbell/nicedishy/pkg/persistence"
)

type WeeklyReportContext struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`

	FastestDownloadSpeed float64 `json:"fastestDownloadSpeed"`
	FastestUploadSpeed   float64 `json:"fastestUploadSpeed"`

	SlowestDownloadSpeedTimeBlock string `json:"slowestDownloadSpeedTimeBlock"`
}

func GenerateReports() {
	ctx := context.Background()

	for {
		dishyID, tzOffset, err := getNextDishyIDForWeeklyReport(ctx)
		if err != nil {
			logger.Errorf("error getting next dishy id for weekly report: %v", err)
			time.Sleep(time.Second * 5)
			continue
		}

		if dishyID == "" {
			logger.Info("no dishy id found for weekly report")
			time.Sleep(time.Second * 5)
			continue
		}

		if err := generateWeeklyReport(ctx, dishyID, tzOffset); err != nil {
			logger.Errorf("error generating weekly report for dishy id %s: %v", dishyID, err)
			time.Sleep(time.Second * 5)
			continue
		}

		time.Sleep(time.Minute)
	}
}

func getNextDishyIDForWeeklyReport(ctx context.Context) (string, int, error) {
	// get the next dishy that needs a report, and mark that as in progress
	// expect that there are multiple replicas os this service running
	// so we need to respect and let them all operate
	metricsDB := persistence.MustGetMetricsDBSession()

	// the query here joins on geo to just make sure we have geo
	query := `select d.id
from dishy d
inner join dishy_geo g on g.id = d.id
where
d.last_metric_at is not null and
d.id not in (select dishy_id from dishy_report_weekly where week_end > now() - interval '2 days')
order by random()`

	rows, err := metricsDB.Query(ctx, query)
	if err != nil {
		return "", 0, fmt.Errorf("error querying for dishy id: %v", err)
	}
	defer rows.Close()

	dishyIDs := []string{}
	for rows.Next() {
		dishyID := ""

		if err := rows.Scan(&dishyID); err != nil {
			return "", 0, fmt.Errorf("error scanning for dishy id: %v", err)
		}

		dishyIDs = append(dishyIDs, dishyID)
	}
	rows.Close()

	for _, dishyID := range dishyIDs {
		// check the time zone for this one)
		query = `select
g.timezone_offset from dishy_geo g
where g.id = $1 order by g.time desc limit 1`
		row := metricsDB.QueryRow(ctx, query, dishyID)
		var timezoneOffset sql.NullInt64
		if err := row.Scan(&timezoneOffset); err != nil {
			logger.Errorf("error getting timezone offset: %w", err)
			time.Sleep(time.Minute)
			continue
		}

		if !timezoneOffset.Valid {
			// if the timezone offset means that we have a new week for this dishy,
			if time.Now().UTC().Add(time.Second*time.Duration(timezoneOffset.Int64)).Weekday() == time.Sunday {
				return dishyID, int(timezoneOffset.Int64), nil
			}
		}

		time.Sleep(time.Minute)
	}

	return "", 0, nil
}

func dateRangeForWeeklyReportInTimezone(tzOffset int) (time.Time, time.Time, error) {
	justTheDate := fmt.Sprintf("%d-%02d-%02dT00:00:00Z", time.Now().Year(), time.Now().Month(), time.Now().Day())

	sundayStart, _ := time.Parse(time.RFC3339, justTheDate)

	i := 0
	for int(sundayStart.Weekday()) > int(time.Sunday) {
		sundayStart = sundayStart.AddDate(0, 0, -1)
		i++

		if i > 10 {
			return time.Now(), time.Now(), fmt.Errorf("could not find sunday start")
		}
	}

	sundayEnd := sundayStart.Add(time.Duration(7*24) * time.Hour)

	// now move to last week
	sundayStart = sundayStart.AddDate(0, 0, -7)
	sundayEnd = sundayEnd.AddDate(0, 0, -7)

	// convert to the dishy time zone
	sundayStart = sundayStart.Add(time.Second * time.Duration(tzOffset))
	sundayEnd = sundayEnd.Add(time.Second * time.Duration(tzOffset))

	return sundayStart, sundayEnd, nil
}

func markReportInProgress(ctx context.Context, dishyID string, start, end time.Time) error {
	metricsDB := persistence.MustGetMetricsDBSession()

	query := `insert into dishy_report_weekly (dishy_id, week_start, week_end, report_context, is_generating)
values
($1,
$2,
$3,
'{}', true)`

	_, err := metricsDB.Exec(ctx, query, dishyID, start, end)
	if err != nil {
		return fmt.Errorf("error inserting weekly report: %v", err)
	}

	return nil
}

func generateWeeklyReport(ctx context.Context, dishyID string, tzOffset int) error {
	fmt.Printf("generating weekly report for dishy id: %s\n", dishyID)

	start, end, err := dateRangeForWeeklyReportInTimezone(tzOffset)
	if err != nil {
		return fmt.Errorf("error getting date range for weekly report: %v", err)
	}

	if err := markReportInProgress(ctx, dishyID, start, end); err != nil {
		return fmt.Errorf("error marking report in progress: %v", err)
	}

	weeklyContext, err := generateWeeklyReportContext(ctx, dishyID, start, end)
	if err != nil {
		return fmt.Errorf("error generating weekly report context: %v", err)
	}

	fmt.Printf("weekly report context: %+v\n", weeklyContext)
	return nil
}

func generateWeeklyReportContext(ctx context.Context, dishyID string, start time.Time, end time.Time) (*WeeklyReportContext, error) {
	weeklyReportContext := WeeklyReportContext{
		Start: start,
		End:   end,
	}

	return &weeklyReportContext, nil
}
