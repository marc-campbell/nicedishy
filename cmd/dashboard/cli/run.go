package cli

import (
	"context"

	"github.com/marc-campbell/nicedishy/pkg/grafanaproxy"
	"github.com/marc-campbell/nicedishy/pkg/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func RunCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Starts the dashboard",
		Long:  `starts the nicedishy dashboard proxy`,
		PreRun: func(cmd *cobra.Command, args []string) {
			viper.BindPFlags(cmd.Flags())
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			v := viper.GetViper()

			if v.GetString("log-level") == "debug" {
				logger.Info("setting log level to debug")
				logger.SetDebug()
			}

			ctx := context.Background()
			grafanaproxy.Start(ctx)
			<-ctx.Done()

			return nil
		},
	}

	cmd.Flags().String("log-level", "info", "set the log level")

	return cmd
}
