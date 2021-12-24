package cli

import (
	"fmt"
	"math/rand"
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
			rand.Seed(time.Now().Unix())

			v := viper.GetViper()

			token := v.GetString("token")

			current := time.Now().UTC().Add(-1 * viper.GetDuration("duration"))

			if v.GetBool("fill") {
				// get the most recent time and set current the the interval after that
				// TODO
			}

			uptimeSeconds := 1500

			for {
				if err := generator.GenerateAndSendData(v.GetString("endpoint"), token, current, uptimeSeconds); err != nil {
					return err
				}

				current = current.Add(time.Minute * 5)
				uptimeSeconds += 5

				if current.After(time.Now()) {
					return nil
				}

				// time.Sleep(time.Second * 5)
			}
		},
	}

	cobra.OnInitialize(initConfig)

	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	cmd.Flags().String("token", "", "the token")
	cmd.MarkFlagRequired("token")
	cmd.Flags().String("endpoint", "https://api.nicedishy.com", "endpoint of the api server")
	cmd.Flags().Duration("duration", time.Hour*7*24, "number of days to ingest")
	cmd.Flags().Bool("fill", false, "set to true to fill the most recent gap of data")

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
