package main

import (
	"fmt"

	"github.com/Niessy/fisherman/api/v1"
	"github.com/codegangsta/cli"
)

func sizes(c *cli.Context) {
	auth(c)
	sizes, err := v1.GetSizes()
	if err != nil {
		fatalf(err.Error())
	}
	fmt.Println("Sizes:")
	for _, s := range sizes {
		fmt.Printf("%s (id: %d)\n", s.Name, s.ID)
	}
}
