package dishy

import (
	"context"
	"time"
)

func GetFourHourStart(ctx context.Context, timezoneOffset int, when time.Time) (*time.Time, error) {
	when = when.Truncate(time.Hour)
	when = when.Add(-(time.Duration(timezoneOffset) * time.Minute))

	for {
		if when.Hour()%4 == 0 {
			when = when.Add(-time.Duration(timezoneOffset) * time.Second)
			return &when, nil
		}
		when = when.Add(-time.Hour)
	}
}

func GetDayStart(ctx context.Context, timezoneOffset int, when time.Time) (*time.Time, error) {
	when = when.Truncate(time.Hour)
	when = when.Add(time.Duration(timezoneOffset) * time.Second)

	for {
		if when.UTC().Hour() == 0 {
			when = when.Add(-time.Duration(timezoneOffset) * time.Second)
			return &when, nil
		}
		when = when.Add(-time.Hour)
	}
}
