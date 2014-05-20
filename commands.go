// TODO: figure how to only auth at one point
package main

import (
	"log"

	"github.com/Niessy/fisherman/api/v1"
	"github.com/codegangsta/cli"
)

func auth(c *cli.Context) {
	path := c.String("configFile")
	err := v1.LoadConfig(path)
	if err != nil {
		log.Fatal(err)
	}
}

func droplets(c *cli.Context) {
	auth(c)
	v1.GetDroplets()
}

func regions(c *cli.Context) {
	auth(c)
	v1.GetRegions()
}

func images(c *cli.Context) {
	auth(c)
	v1.GetImages()
}

func sizes(c *cli.Context) {
	auth(c)
	v1.GetSizes()
}

func keys(c *cli.Context) {
	auth(c)
	v1.GetSSHKeys()
}
