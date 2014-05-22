package main

import (
	"fmt"

	"github.com/Niessy/fisherman/api/v1"
	"github.com/codegangsta/cli"
)

func allDroplets(c *cli.Context) {
	auth(c)
	droplets, err := v1.GetDroplets()
	if err != nil {
		fatalf(err.Error())
	}
	println("Droplets:")
	for _, d := range droplets {
		fmt.Printf("%s (id %d region: %d image_id: %d ip: \"%s\" status: %s)\n",
			d.Name,
			d.ID,
			d.RegionID,
			d.ImageID,
			d.IPAddress,
			d.Status)
	}
}
