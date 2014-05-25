package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/Niessy/dogo"
)

func sizes(c *cli.Context) {
	auth(c)
	sizes, err := dogo.GetSizes()
	if err != nil {
		fatalf(err.Error())
	}
	fmt.Println("Sizes:")
	for _, s := range sizes {
		fmt.Printf("%s (id: %d)\n", s.Name, s.ID)
	}
}
