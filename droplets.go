package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

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
	p := c.String("path")

	// Clean paths, get public keys
	kpaths := strings.Split(p, ",")
	for i, kp := range kpaths {
		buf, err := ioutil.ReadFile(cleanPath(kp, "/"))
		if err != nil {
			log.Println(err)
		}
		kpaths[i] = string(buf)
	}

	droplet, err := v1.CreateDroplet(name, sID, iID, rID, kpaths)
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
