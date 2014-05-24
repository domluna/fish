package main

import (
	"fmt"
	"strconv"

	"github.com/Niessy/fisherman/api/v1"
	"github.com/codegangsta/cli"
)

func droplets(c *cli.Context) {
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

func create(c *cli.Context) {
	auth(c)
	args := c.Args()
	if len(args) < 1 {
		cli.ShowCommandHelp(c, c.Command.Name)
		return
	}

	name := args[0]

	iID := c.Int("image")
	sID := c.Int("size")
	rID := c.Int("region")
	keys := c.String("keys")

	droplet, err := v1.CreateDroplet(name, sID, iID, rID, keys)
	if err != nil {
		fatalf(err.Error())
	}
	fmt.Printf("Successfully queued %s for creation ... \n", name)
	fmt.Printf("%v\n", droplet)
}

func destroy(c *cli.Context) {
	auth(c)
	args := c.Args()
	if len(args) < 1 {
		cli.ShowCommandHelp(c, c.Command.Name)
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fatalf(err.Error())
	}

	err = v1.DestroyDroplet(id)
	if err != nil {
		fatalf(err.Error())
	}

	fmt.Printf("Queing droplet for deletion ... \n")
}

func resize(c *cli.Context) {
	auth(c)
	args := c.Args()
	if len(args) < 1 {
		fmt.Printf("Incorrect Usage\n")
		cli.ShowCommandHelp(c, c.Command.Name)
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fatalf(err.Error())
	}

	slug := c.String("size")

	err = v1.ResizeDroplet(id, slug)
	if err != nil {
		fatalf(err.Error())
	}
	fmt.Printf("resizing droplet\n")
}

func reboot(c *cli.Context) {
	auth(c)
	args := c.Args()
	if len(args) < 1 {
		fmt.Printf("Incorrect Usage\n")
		cli.ShowCommandHelp(c, c.Command.Name)
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fatalf(err.Error())
	}

	err = v1.RebootDroplet(id)
	if err != nil {
		fatalf(err.Error())
	}
	fmt.Printf("rebooting droplet\n")
}

func rebuild(c *cli.Context) {
	auth(c)
}

func stop(c *cli.Context) {
	auth(c)
}

func start(c *cli.Context) {
	auth(c)
}

func snapshot(c *cli.Context) {
	auth(c)
}

func info(c *cli.Context) {
}
