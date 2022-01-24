package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/grafana-tools/sdk"
	"github.com/marc-campbell/nicedishy/pkg/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func ExportDashboardCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "export-dashboard",
		Short: "Exports a dashboard to json",
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

			c, err := sdk.NewClient("http://grafana:3000", os.Getenv("GRAFANA_API_KEY"), sdk.DefaultHTTPClient)
			if err != nil {
				return err
			}
			board, boardProperties, err := c.GetDashboardBySlug(ctx, v.GetString("dashboard"))
			if err != nil {
				return err
			}

			fmt.Printf("%#v\n", board)
			fmt.Printf("%#v\n", boardProperties)
			return nil
		},
	}

	cmd.Flags().String("log-level", "info", "set the log level")
	cmd.Flags().String("dashboard", "", "the slug of the dashboard to export")

	return cmd
}
