// TODO: figure how to only auth at one point
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

func droplets(c *cli.Context) {
	auth(c)
	api.GetDroplets()
}

func regions(c *cli.Context) {
	auth(c)
	api.GetRegions()
}

func images(c *cli.Context) {
	auth(c)
	api.GetImages()
}
