package cli

import (
	"github.com/marc-campbell/nicedishy/pkg/apiserver"
	"github.com/marc-campbell/nicedishy/pkg/logger"
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

			apiserver.Start()
			return nil
		},
	}

	cmd.Flags().String("log-level", "info", "set the log level")

	return cmd
}
