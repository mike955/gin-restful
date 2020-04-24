package server

import (
	//"apiServer/web"

	"github.com/spf13/cobra"
	"webgo/web"
)

// base package, always hold on to anti golang build warning
func preRun(c *cobra.Command, args []string) {
}

func run(c *cobra.Command, args []string) {
	zkCluster := vpFlag.GetString("zk")
	intstanceId := vpFlag.GetString("i")
	appName := vpFlag.GetString("appName")
	if err := web.Run(zkCluster, intstanceId, appName); err !=nil {
		// TODO add log
		panic(err)
	}
}
