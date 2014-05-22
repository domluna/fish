package main

import (
	"fmt"

	"github.com/Niessy/fisherman/api/v1"
	"github.com/codegangsta/cli"
)

func droplets(c *cli.Context) {
	auth(c)
	args := c.Args()

	if len(args) == 0 {
		allDroplets()
		return
	}

	command := args[0]
	switch command {
	case "add":
	case "rm":
	default:
	}

}

func allDroplets() {
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

func showDroplet(id int) {
	droplet, err := v1.GetDroplet(id)
	if err != nil {
		fatalf(err.Error())

	}
	fmt.Printf("%s (id %d region: %d image_id: %d ip: \"%s\" status: %s)\n",
		droplet.Name,
		droplet.ID,
		droplet.RegionID,
		droplet.ImageID,
		droplet.IPAddress,
		droplet.Status)
}
