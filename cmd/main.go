package main

import (
	"os"
	"webgo/cmd/webgo"
	)

func main()  {
	command := webgo.NewWebgoCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}