package main

import (
	"webgo/pkg/net/http/kama"
)

func main()  {
	r := kama.Default()
	r.GET("/ping", func(c *kama.Context) {
		c.JSON(200, kama.H{
			"message": "pong",
		})
	})
	r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}