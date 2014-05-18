package main

import (
	"log"

	"github.com/Niessy/fisherman/api"
	"github.com/codegangsta/cli"
)

func auth(c *cli.Context) {
	path := c.String("configFile")
	err := api.LoadConfig(path)
	if err != nil {
		log.Fatal(err)
	}
}

func droplet(c *cli.Context) {
	auth(c)
	api.GetAllDroplets()
}
