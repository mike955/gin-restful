package webgo

import (
	"github.com/spf13/cobra"
	"webgo/build"

	cmdVersion "webgo/cmd/webgo/version"
	cmdServer "webgo/cmd/webgo/server"
)

func init() {
	cmdVersion.RegisterTo(cmd, vpFlag)
	cmdServer.RegisterTo(cmd, vpFlag)
}

var cmd = &cobra.Command{
	Use:     build.Appname(),
	Version: build.Version(),

	Short: ``,
	Long:  ``,

	PersistentPreRun: func(c *cobra.Command, args []string) {
		vpFlag.BindPFlags(c.Flags())
		if cmdVersion.Called(c) {
			return
		}
	},

	PreRun: func(c *cobra.Command, args []string) {
	},

	Run: func(c *cobra.Command, args []string) {
		c.Help()
	},
}
func NewWebgoCommand() *cobra.Command {
	return cmd
}

