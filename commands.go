// TODO: figure how to only auth at one point
package main

import (
	"log"
	"fmt"

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

func allDroplets(c *cli.Context) {
	auth(c)
	v1.GetDroplets()
}

func allRegions(c *cli.Context) {
	auth(c)
	v1.GetRegions()
}

func allImages(c *cli.Context) {
	auth(c)
	v1.GetImages()
}

func allSizes(c *cli.Context) {
	auth(c)
	v1.GetSizes()
}

func allKeys(c *cli.Context) {
	auth(c)
	keys, err := v1.AllSSHKeys()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("SSHKeys: ")
	for _, k := range keys {
		fmt.Printf("%v\n", k)
	}
	fmt.Println()
}
