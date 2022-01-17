package cli

import (
	"github.com/marc-campbell/nicedishy/pkg/logger"
	"github.com/marc-campbell/nicedishy/pkg/mailer"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func MailerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mailer",
		Short: "Starts the mailer daemon",
		Long:  `starts the nicedishy mailer daemon`,
		PreRun: func(cmd *cobra.Command, args []string) {
			viper.BindPFlags(cmd.Flags())
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			v := viper.GetViper()

			if v.GetString("log-level") == "debug" {
				logger.Info("setting log level to debug")
				logger.SetDebug()
			}

			mailer.StartDaemon()
			return nil
		},
	}

	cmd.Flags().String("log-level", "info", "set the log level")

	return cmd
}
