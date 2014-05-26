package main

import (
	"fmt"
	"github.com/codegangsta/cli"
)

func images(c *cli.Context) {
	filter := "my_images"
	if c.Bool("global") {
		filter = "global"
	}

	images, err := docli.GetImages(filter)
	if err != nil {
		fatalf(err.Error())
	}

	switch filter {
	case "global":
		fmt.Println("Global Images:")
	default:
		fmt.Println("My Images:")
	}

	for _, i := range images {
		fmt.Printf("%s (distribution: %s id: %d)\n",
			i.Name,
			i.Slug,
			i.ID)
	}
}
