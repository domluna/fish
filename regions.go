package main

import (
	"fmt"

	"github.com/Niessy/dogo"
	"github.com/codegangsta/cli"
)

func regions(c *cli.Context) {
	auth(c)

	args := c.Args()
	if len(args) != 0 {
		fatalf("%s takes no arguments", "regions")
	}

	regions, err := v1.GetRegions()
	if err != nil {
		fatalf(err.Error())
	}

	println("Regions:")
	for _, r := range regions {
		fmt.Printf("%s (id: %d)\n", r.Name, r.ID)
	}
}
