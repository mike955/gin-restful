package webgo

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewCmdVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "version",
		Short: ``,
		Long:  ``,

		PersistentPreRun: func(c *cobra.Command, args []string) {
		},

		PreRun: func(c *cobra.Command, args []string) {
		},

		Run: versionRun,
	}
	return cmd
}

func versionRun(c *cobra.Command, args []string) {
	fmt.Printf("version: %s", VERSION)
}