package webgo

import (
	"github.com/spf13/cobra"
)

var (
	VERSION = "0.0.1"
)

func init() {
}

func NewWebgoCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "webgo",
		Version: VERSION,

		Short: ``,
		Long:  ``,

		PersistentPreRun: func(c *cobra.Command, args []string) {
		},

		PreRun: func(c *cobra.Command, args []string) {
		},

		Run: func(c *cobra.Command, args []string) {
			c.Help()
		},
	}
	cmd.AddCommand(NewCmdServer())
	cmd.AddCommand(NewCmdVersion())
	return cmd
}


