package webgo

import (
	"github.com/spf13/viper"
)

var vpFlag *viper.Viper = viper.New()

func init() {
	cmd.PersistentFlags().StringP(
		"zk", "z", "", "zk connstr",
	)
	cmd.PersistentFlags().StringP(
		"i", "i", "0", "intstance id, default 0",
	)
	cmd.PersistentFlags().StringP(
		"appName", "a", "apiServer", "application name, like apiServer",
	)
}
