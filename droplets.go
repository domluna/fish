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
		fmt.Printf("%s (id: %d region: %d image_id: %d ip: %q status: %s)\n",
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
	checkArgs(c)

	name := c.Args().First()

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
	checkArgs(c)

	id, err := strconv.Atoi(c.Args().First())
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
	checkArgs(c)

	id, err := strconv.Atoi(c.Args().First())
	if err != nil {
		fatalf(err.Error())
	}

	slug := c.String("size")

	err = v1.ResizeDroplet(id, slug)
	if err != nil {
		fatalf(err.Error())
	}
	fmt.Printf("Resizing droplet\n")
}

func reboot(c *cli.Context) {
	auth(c)
	checkArgs(c)

	id, err := strconv.Atoi(c.Args().First())
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
	checkArgs(c)

	id, err := strconv.Atoi(c.Args().First())
	if err != nil {
		fatalf(err.Error())
	}

	image := c.Int("image")

	err = v1.RebuildDroplet(id, image)
	if err != nil {
		fatalf(err.Error())
	}
	fmt.Printf("Rebuilding droplet\n")
}

func stop(c *cli.Context) {
	auth(c)
	checkArgs(c)

	id, err := strconv.Atoi(c.Args().First())
	if err != nil {
		fatalf(err.Error())
	}

	err = v1.StopDroplet(id)
	if err != nil {
		fatalf(err.Error())
	}
	fmt.Println("Stopped Droplet")
}

func start(c *cli.Context) {
	auth(c)
	checkArgs(c)

	id, err := strconv.Atoi(c.Args().First())
	if err != nil {
		fatalf(err.Error())
	}

	err = v1.StartDroplet(id)
	if err != nil {
		fatalf(err.Error())
	}
	fmt.Println("Started Droplet")
}

func snapshot(c *cli.Context) {
	auth(c)
	checkArgs(c)

	id, err := strconv.Atoi(c.Args().First())
	if err != nil {
		fatalf(err.Error())
	}

	name := c.String("name")

	err = v1.SnapshotDroplet(id, name)
	if err != nil {
		fatalf(err.Error())
	}
	fmt.Printf("Created a snapshot, name = %s\n", name)
}

func restore(c *cli.Context) {
	auth(c)
	checkArgs(c)

	id, err := strconv.Atoi(c.Args().First())
	if err != nil {
		fatalf(err.Error())
	}

	image := c.Int("image")

	err = v1.RestoreDroplet(id, image)
	if err != nil {
		fatalf(err.Error())
	}
	fmt.Println("Restored Droplet with image id = %d", image)
}

func info(c *cli.Context) {
	auth(c)
	checkArgs(c)

	id, err := strconv.Atoi(c.Args().First())
	if err != nil {
		fatalf(err.Error())
	}

	droplet, err := v1.GetDroplet(id)
	if err != nil {
		fatalf(err.Error())
	}
	fmt.Printf("Droplet:\t %s\n", droplet.Name)
	fmt.Printf("ID:\t %d\n", droplet.ID)
	fmt.Printf("Image ID:\t %d\n", droplet.ImageID)
	fmt.Printf("Size ID:\t %d\n", droplet.SizeID)
	fmt.Printf("Region ID:\t %d\n", droplet.RegionID)
	fmt.Printf("Backups Active:\t %t\n", droplet.BackupsActive)
	fmt.Printf("IP Address:\t %q\n", droplet.IPAddress)
	fmt.Printf("Private IP Address:\t %q\n", droplet.PrivateIPAddress)
	fmt.Printf("Locked:\t %t\n", droplet.Locked)
	fmt.Printf("Status:\t %s\n", droplet.Status)
	fmt.Printf("Created At:\t %v\n", droplet.CreatedAt)
}
