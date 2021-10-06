package cli

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/marc-campbell/nicedishy/pkg/generator"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate-data",
		Short: "generate-data",
		Long:  ``,
		PreRun: func(cmd *cobra.Command, args []string) {
			viper.BindPFlags(cmd.Flags())
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			v := viper.GetViper()

			token := v.GetString("token")

			// generate 4 weeks of data for the token
			end := time.Now()
			current := end.Add(-time.Duration(time.Hour * 24 * 7 * 4))

			for current.Before(end) {
				if err := generator.GenerateAndSendData(token, current); err != nil {
					return err
				}

				current = current.Add(time.Minute * 5)
			}
			return nil
		},
	}

	cobra.OnInitialize(initConfig)

	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	cmd.Flags().String("token", "", "the token")
	cmd.MarkFlagRequired("token")

	return cmd
}

func InitAndExecute() {
	if err := RootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() {
	viper.SetEnvPrefix("NICEDISHY")
	viper.AutomaticEnv()
}
