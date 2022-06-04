package cli

import (
	"log"

	"github.com/getsentry/sentry-go"
	"github.com/marc-campbell/nicedishy/pkg/apiserver"
	"github.com/marc-campbell/nicedishy/pkg/logger"
	"github.com/marc-campbell/nicedishy/pkg/mailer"
	"github.com/marc-campbell/nicedishy/pkg/reports"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func APICmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "api",
		Short: "Starts the API server",
		Long:  `starts the nicedishy rest api server`,
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
				Dsn: "https://fe53e0b20cbc4565b2deb5b577dbb8a8@o242537.ingest.sentry.io/6247592",
			})
			if err != nil {
				log.Fatalf("sentry.Init: %s", err)
			}

			go mailer.QueueDisconnectedDeviceEmails()
			go reports.GenerateReports()

			apiserver.Start()
			return nil
		},
	}

	cmd.Flags().String("log-level", "info", "set the log level")

	return cmd
}
