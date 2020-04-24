package server

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var vpFlag *viper.Viper

var cmd = &cobra.Command{
	Use:    "server",
	Short:  ``,
	Long:   ``,
	PreRun: preRun,
	Run:    run,
}

func init() {
}

func RegisterTo(father *cobra.Command, flag *viper.Viper) {
	father.AddCommand(cmd)
	vpFlag = flag
}

func Called(c *cobra.Command) bool {
	return c == cmd
}
