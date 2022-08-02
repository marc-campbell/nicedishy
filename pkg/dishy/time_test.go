package dishy

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetDayStart(t *testing.T) {
	reindexIn, err := time.Parse(time.RFC3339, "2022-01-01T00:01:00Z")
	if err != nil {
		panic(err)
	}
	reindexWant, err := time.Parse(time.RFC3339, "2021-12-31T7:00:00Z")
	if err != nil {
		panic(err)
	}
	noonPDTIn, err := time.Parse(time.RFC3339, "2020-04-01T12:00:00-07:00")
	if err != nil {
		panic(err)
	}
	noonPDTWant, err := time.Parse(time.RFC3339, "2020-04-01T00:00:00-07:00")
	if err != nil {
		panic(err)
	}

	eighteenThirtyFourESTIn, err := time.Parse(time.RFC3339, "2020-04-19T18:34:00-05:00")
	if err != nil {
		panic(err)
	}
	eighteenThirtyFourESTWant, err := time.Parse(time.RFC3339, "2020-04-19T00:00:00-05:00")
	if err != nil {
		panic(err)
	}

	type args struct {
		timezoneOffset int
		when           time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "noon pacific daylight time",
			args: args{
				timezoneOffset: -25200,
				when:           noonPDTIn,
			},
			want: noonPDTWant,
		},
		{
			name: "18:34 est",
			args: args{
				timezoneOffset: -18000,
				when:           eighteenThirtyFourESTIn,
			},
			want: eighteenThirtyFourESTWant,
		},
		{
			name: "reindex",
			args: args{
				timezoneOffset: -25200,
				when:           reindexIn,
			},
			want: reindexWant,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := require.New(t)

			got, err := GetDayStart(context.TODO(), tt.args.timezoneOffset, tt.args.when)
			req.NoError(err)

			assert.Equal(t, tt.want, *got)
		})
	}
}

func TestGetFourHourStart(t *testing.T) {
	reindexIn, err := time.Parse(time.RFC3339, "2022-01-01T00:01:00Z")
	if err != nil {
		panic(err)
	}
	reindexWant, err := time.Parse(time.RFC3339, "2022-01-01T07:00:00Z")
	if err != nil {
		panic(err)
	}

	noonPDTIn, err := time.Parse(time.RFC3339, "2020-04-01T12:00:00-07:00")
	if err != nil {
		panic(err)
	}
	noonPDTWant, err := time.Parse(time.RFC3339, "2020-04-01T19:00:00Z")
	if err != nil {
		panic(err)
	}

	eighteenThirtyFourESTIn, err := time.Parse(time.RFC3339, "2020-04-19T18:34:00-05:00")
	if err != nil {
		panic(err)
	}
	eighteenThirtyFourESTWant, err := time.Parse(time.RFC3339, "2020-04-19T21:00:00Z")
	if err != nil {
		panic(err)
	}

	type args struct {
		timezoneOffset int
		when           time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "noon pacific daylight time",
			args: args{
				timezoneOffset: -25200,
				when:           noonPDTIn,
			},
			want: noonPDTWant,
		},
		{
			name: "18:34 est",
			args: args{
				timezoneOffset: -18000,
				when:           eighteenThirtyFourESTIn,
			},
			want: eighteenThirtyFourESTWant,
		},
		{
			name: "reindex",
			args: args{
				timezoneOffset: -25200,
				when:           reindexIn,
			},
			want: reindexWant,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := require.New(t)

			got, err := GetFourHourStart(context.TODO(), tt.args.timezoneOffset, tt.args.when)
			req.NoError(err)

			assert.Equal(t, &tt.want, got)

		})
	}
}
