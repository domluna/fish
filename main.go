//
// CLI for DigitalOcean
//
package main

import (
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
			Action: droplets,
		},
		// regions
		{
			Name:   "regions",
			Usage:  "Stuff with Regions",
			Action: regions,
		},
		// images
		{
			Name:   "images",
			Usage:  "Stuff with Images",
			Action: images,
		},
		// sizes
		{
			Name:   "sizes",
			Usage:  "Stuff with Sizes",
			Action: sizes,
		},
		// sshkeys
		{
			Name:   "keys",
			Usage:  "Stuff with SSHKeys",
			Action: keys,
		},

	}
	app.Run(os.Args)
}
