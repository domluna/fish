package main

import (
	"fmt"
	"github.com/codegangsta/cli"
)

func regions(c *cli.Context) {
	regions, err := docli.GetRegions()
	if err != nil {
		fatalf(err.Error())
	}

	println("Regions:")
	for _, r := range regions {
		fmt.Printf("%s (slug: %s id: %d)\n", r.Name, r.Slug, r.ID)
	}
}
