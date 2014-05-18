//
// CLI for DigitalOcean
//
package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "fisherman"
	app.Usage = "Command Line Interface for DigitalOcean"
	app.Commands = []cli.Command{
		// authorize
		{
			Name:  "auth",
			Usage: "Authenticate DigitalOcean credentials",
			Flags: []cli.Flag{
				cli.StringFlag{"configFile, conf", "", "credentials file"},
			},
			Action: auth,
		},
		// droplets
		{
			Name:   "droplets",
			Usage:  "Stuff with Droplets",
			Action: droplet,
		},
		// regions
		{
			Name:  "regions",
			Usage: "Stuff with Regions",
			Action: func(c *cli.Context) {
				fmt.Println("Nothing here yet!")
			},
		},
		// images
		{
			Name:  "images",
			Usage: "Stuff with Images",
			Action: func(c *cli.Context) {
				fmt.Println("Nothing here yet!")
			},
		},
	}
	app.Run(os.Args)
}
