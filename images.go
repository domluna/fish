package main

import (
	"fmt"

	"github.com/Niessy/fisherman/api/v1"
	"github.com/codegangsta/cli"
)

func allImages(c *cli.Context) {
	auth(c)

	filter := "my_images"
	if c.Bool("global") {
		filter = "global"
	}

	images, err := v1.GetImages(filter)
	if err != nil {
		fatalf(err.Error())
	}

	println("Images:")
	for _, i := range images {
		fmt.Printf("%s (distribution: %s id: %d)\n",
			i.Name,
			i.Distribution,
			i.ID)
	}
}