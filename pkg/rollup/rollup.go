package rollup

import (
	"context"
	"fmt"
	"time"

	"github.com/marc-campbell/nicedishy/pkg/dishy"
	"github.com/marc-campbell/nicedishy/pkg/logger"
	"github.com/marc-campbell/nicedishy/pkg/persistence"
)

type DishyWithTimeZoneOffset struct {
	DishyID        string
	TimeZoneOffset int
}

func ReindexAll() error {
	pg := persistence.MustGetPGSession()

	query := `select distinct(d.id), g.timezone_offset from dishy d inner join dishy_geo g on d.id = g.id
where g.timezone_offset is not null`
	rows, err := pg.Query(context.Background(), query)
	if err != nil {
		return fmt.Errorf("error querying for dishy: %v", err)
	}
	defer rows.Close()

	dishiesWithTimeZoneOffset := []DishyWithTimeZoneOffset{}
	for rows.Next() {
		dishyWithTimeZoneOffset := DishyWithTimeZoneOffset{}
		err := rows.Scan(&dishyWithTimeZoneOffset.DishyID, &dishyWithTimeZoneOffset.TimeZoneOffset)
		if err != nil {
			return fmt.Errorf("error scanning for dishy: %v", err)
		}

		dishiesWithTimeZoneOffset = append(dishiesWithTimeZoneOffset, dishyWithTimeZoneOffset)
	}
	rows.Close()

	for _, dishyWithTimeZoneOffset := range dishiesWithTimeZoneOffset {
		logger.Infof("reindexing dishy %s", dishyWithTimeZoneOffset.DishyID)
		min, err := time.Parse(time.RFC3339, "2022-01-01T00:01:00Z")
		if err != nil {
			return fmt.Errorf("error parsing start time: %v", err)
		}

		// hourly
		max := time.Now()
		for min.Before(max) {
			if err := ReindexSpeedHourly(context.Background(), dishyWithTimeZoneOffset.DishyID, min); err != nil {
				return fmt.Errorf("error reindexing hourly: %v", err)
			}
			if err := ReindexDataHourly(context.Background(), dishyWithTimeZoneOffset.DishyID, min); err != nil {
				return fmt.Errorf("error reindexing hourly: %v", err)
			}

			min = min.Add(time.Hour)
		}

		// daily
		min, err = time.Parse(time.RFC3339, "2022-01-01T00:01:00Z")
		if err != nil {
			return fmt.Errorf("error parsing start time: %v", err)
		}

		dayMin, err := dishy.GetDayStart(context.Background(), dishyWithTimeZoneOffset.TimeZoneOffset, min)
		if err != nil {
			return fmt.Errorf("error getting day start: %v", err)
		}
		min = *dayMin
		for min.Before(max) {
			if err := ReindexSpeedDaily(context.Background(), dishyWithTimeZoneOffset.DishyID, min); err != nil {
				return fmt.Errorf("error reindexing daily: %v", err)
			}
			if err := ReindexDataDaily(context.Background(), dishyWithTimeZoneOffset.DishyID, min); err != nil {
				return fmt.Errorf("error reindexing daily: %v", err)
			}

			min = min.Add(time.Hour * 24)
		}

		// four hour
		min, err = time.Parse(time.RFC3339, "2022-01-01T00:01:00Z")
		if err != nil {
			return fmt.Errorf("error parsing start time: %v", err)
		}
		fourHourMin, err := dishy.GetFourHourStart(context.Background(), dishyWithTimeZoneOffset.TimeZoneOffset, min)
		if err != nil {
			return fmt.Errorf("error getting four hour start: %v", err)
		}
		min = *fourHourMin
		for min.Before(max) {
			if err := ReindexSpeedFourHour(context.Background(), dishyWithTimeZoneOffset.DishyID, min); err != nil {
				return fmt.Errorf("error reindexing four hour: %v", err)
			}
			if err := ReindexDataFourHour(context.Background(), dishyWithTimeZoneOffset.DishyID, min); err != nil {
				return fmt.Errorf("error reindexing hour hour: %v", err)
			}

			min = min.Add(time.Hour * 4)
		}

	}

	return nil
}
