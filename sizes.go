package main

import (
	"fmt"
	"github.com/codegangsta/cli"
)

func sizes(c *cli.Context) {
	sizes, err := docli.GetSizes()
	if err != nil {
		fatalf(err.Error())
	}
	fmt.Println("Sizes:")
	for _, s := range sizes {
		fmt.Printf("%s (id: %d)\n", s.Name, s.ID)
	}
}
