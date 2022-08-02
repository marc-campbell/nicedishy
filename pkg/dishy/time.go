package dishy

import (
	"context"
	"time"
)

func GetFourHourStart(ctx context.Context, timezoneOffset int, when time.Time) (*time.Time, error) {
	when = when.Truncate(time.Hour)

	// when = when.UTC().Add(-(time.Duration(timezoneOffset) * time.Second))

	for {
		if when.Hour()%4 == 0 {
			// when = when.Add(-time.Duration(timezoneOffset) * time.Second)
			res, err := time.ParseInLocation("2006-01-02 15:04:05", when.Format("2006-01-02 15:04:05"), time.UTC)
			if err != nil {
				return nil, err
			}
			return &res, nil
		}
		when = when.Add(-time.Hour)
	}
}

func GetDayStart(ctx context.Context, timezoneOffset int, when time.Time) (*time.Time, error) {
	when = when.Truncate(time.Hour)
	// when = when.Add(time.Duration(timezoneOffset) * time.Second)

	for {
		if when.UTC().Hour() == 0 {
			// when = when.Add(-time.Duration(timezoneOffset) * time.Second)
			return &when, nil
		}
		when = when.Add(-time.Hour)
	}
}
