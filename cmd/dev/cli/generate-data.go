package cli

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/marc-campbell/nicedishy/pkg/logger"
	"github.com/marc-campbell/nicedishy/pkg/persistence"
	"github.com/marc-campbell/nicedishy/pkg/stores"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func GenerateDataCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate-data",
		Short: "Generates data for a dishy id",
		Long:  ``,
		PreRun: func(cmd *cobra.Command, args []string) {
			viper.BindPFlags(cmd.Flags())
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			v := viper.GetViper()

			if v.GetString("log-level") == "debug" {
				logger.Info("setting log level to debug")
				logger.SetDebug()
			}

			ctx := context.TODO()

			rand.Seed(time.Now().UnixNano())

			d, err := stores.GetStore().GetDishy(ctx, v.GetString("id"))
			if err != nil {
				return err
			}

			if d == nil {
				return errors.New("dishy not found")
			}

			// find the most recent timestamp in the database, and backfill to now
			startFrom := time.Now().Add(-time.Duration(30) * 24 * time.Hour) // 30 days
			if d.LastMetricAt != nil {
				startFrom = *d.LastMetricAt
				startFrom = startFrom.Add(time.Duration(time.Minute))
			}

			nextGeoCheck := startFrom.Add(-1 * time.Duration(time.Minute))
			if d.LastGeocheckAt != nil {
				nextGeoCheck = d.LastGeocheckAt.Add(time.Hour * 30 * 24)
			}

			i := 0
			for startFrom.Before(time.Now()) {
				if nextGeoCheck.Before(startFrom) {
					if err := writeGeoCheck(startFrom, d.ID); err != nil {
						return err
					}
					nextGeoCheck = startFrom.Add(time.Hour * 30 * 24)
				}

				if err := writeData(startFrom, d.ID); err != nil {
					return err
				}

				// every 2 iterations (10 minutes) gets speed
				if i%2 == 0 {
					if err := writeSpeed(startFrom.Add(time.Second), d.ID); err != nil {
						return err
					}
				}

				if err := stores.GetStore().SetDishyLastReceivedStats(ctx, d.ID, startFrom); err != nil {
					return err
				}

				startFrom = startFrom.Add(time.Duration(5) * time.Minute)
				i++
			}

			// and keep running to write more data
			for {
				time.Sleep(time.Minute * 5)

				if nextGeoCheck.Before(startFrom) {
					if err := writeGeoCheck(startFrom, d.ID); err != nil {
						return err
					}
					nextGeoCheck = startFrom.Add(time.Hour * 30 * 24)
				}

				if err := writeData(startFrom, d.ID); err != nil {
					return err
				}

				// every 2 iterations (10 minutes) gets speed
				if i%2 == 0 {
					if err := writeSpeed(startFrom.Add(time.Second), d.ID); err != nil {
						return err
					}
				}

				if err := stores.GetStore().SetDishyLastReceivedStats(ctx, d.ID, startFrom); err != nil {
					return err
				}

				i++
			}
		},
	}

	cmd.Flags().String("log-level", "info", "set the log level")
	cmd.Flags().String("id", "", "the id of the dishy to generate  data for")
	cmd.MarkFlagRequired("id")

	return cmd
}

func writeGeoCheck(when time.Time, id string) error {
	fmt.Printf("writing geocheck for %s at %s\n", id, when)

	metricsDB := persistence.MustGetMetricsDBSession()

	ipAddress := "98.97.57.166"

	query := `insert into dishy_geo
(time, id, ip_address, continent, country, region, city, org, latitude, longitude,
timezone_id, timezone_abbr, timezone_offset, timezone_utc)
values
($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`
	if _, err := metricsDB.Exec(context.TODO(), query, when, id,
		ipAddress,
		"North America",
		"United States",
		"California",
		"Los Angeles",
		"SpaceX Services, Inc.",
		34.0522342,
		-118.2436849,
		"America/Los_Angeles",
		"PDT",
		-25200,
		"-07:00"); err != nil {
		return fmt.Errorf("error inserting geo: %w", err)
	}

	query = `update dishy set last_geocheck_at = $1 where id = $2`
	if _, err := metricsDB.Exec(context.TODO(), query, when, id); err != nil {
		return fmt.Errorf("error updating last_geocheck_at: %w", err)
	}

	return nil
}

func writeData(when time.Time, id string) error {
	fmt.Printf("writing data for %s at %s\n", id, when)

	softwareVersion := "123456-fake.0"
	hardwareVersion := "version1.preproduction"
	ipAddress := "98.97.57.166"

	metricsDB := persistence.MustGetMetricsDBSession()

	// every iteration gets data
	query := `insert into dishy_data (
		time, dishy_id, ip_address, snr, downlink_throughput_bps, uplink_throughput_bps,
		pop_ping_latency_ms, pop_ping_drop_rate, percent_obstructed, seconds_obstructed,
		software_version, hardware_version, user_agent)
		values
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`

	latency := 17 + rand.Float64()*(45-17)
	// 5% chance of a higher latency
	if rand.Intn(20) == 1 {
		latency = 17 + rand.Float64()*(2233-17)
	}

	_, err := metricsDB.Exec(context.Background(), query, when, id, ipAddress, 0,
		rand.Float64()*266517760, // downlink_throughput_bps
		rand.Float64()*68631664,  // uplink_throughput_bps
		latency,                  // pop_ping_latency_ms
		0,                        // pop_ping_drop_rate
		0,                        // percent_obstructed
		0,                        // seconds_obstructed
		softwareVersion,
		hardwareVersion,
		"NiceDishy/Generate/1.0")
	if err != nil {
		return err
	}

	return nil
}

func writeSpeed(when time.Time, id string) error {
	fmt.Printf("writing speed for %s at %s\n", id, when)

	softwareVersion := "123456-fake.0"
	hardwareVersion := "version1.preproduction"
	ipAddress := "98.97.57.166"

	metricsDB := persistence.MustGetMetricsDBSession()

	downloadSpeed := rand.Float64() * 398587307.1445051
	uploadSpeed := rand.Float64() * 40589865.43219394

	// from 4-10 pm pacific, artificially lower the speeds
	if when.UTC().Hour() > 23 || when.UTC().Hour() < 5 {
		downloadSpeed *= .4
		uploadSpeed *= .3
	}

	// every iteration gets data
	query := `insert into dishy_speed (
		time, dishy_id, ip_address,
		download_speed, upload_speed, software_version, hardware_version, user_agent)
		values
		($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := metricsDB.Exec(context.Background(), query,
		when, id, ipAddress,
		downloadSpeed,
		uploadSpeed,
		softwareVersion,
		hardwareVersion,
		"NiceDishy/Generate/1.0")
	if err != nil {
		return err
	}

	return nil
}
