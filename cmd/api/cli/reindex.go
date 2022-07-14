package cli

import (
	"log"

	"github.com/getsentry/sentry-go"
	"github.com/marc-campbell/nicedishy/pkg/logger"
	"github.com/marc-campbell/nicedishy/pkg/rollup"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func ReindexCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reindex",
		Short: "Reindex the rollup data",
		Long:  `Reindexes all rollup data`,
		PreRun: func(cmd *cobra.Command, args []string) {
			viper.BindPFlags(cmd.Flags())
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			v := viper.GetViper()

			if v.GetString("log-level") == "debug" {
				logger.Info("setting log level to debug")
				logger.SetDebug()
			}

			err := sentry.Init(sentry.ClientOptions{
				Dsn: "https://36e408e5af5647d2988935d9b9da584b@o242537.ingest.sentry.io/6247594",
			})
			if err != nil {
				log.Fatalf("sentry.Init: %s", err)
			}

			// this is horribly inefficient, but it's a one-off script
			if err := rollup.ReindexAll(); err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().String("log-level", "info", "set the log level")

	return cmd
}
