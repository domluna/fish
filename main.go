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
	app.Author = "Dominique Luna"
	app.Usage = "Command Line Interface for DigitalOcean"
	app.Commands = []cli.Command{
		// authorize
		{
			Name:      "auth",
			ShortName: "a",
			Usage:     "Authenticate DigitalOcean credentials",
			Flags: []cli.Flag{
				cli.StringFlag{"configFile, conf", "", "credentials file"},
			},
			Action: auth,
		},
		// droplets
		{
			Name:      "droplets",
			ShortName: "d",
			Usage:     "things with digitalocean droplets",
			Action: func(c *cli.Context) {

			},
		},
		// regions
		{
			Name:      "regions",
			ShortName: "r",
			Usage:     "things with digitalocean regions",
			Action:    allRegions,
		},
		// images
		{
			Name:      "images",
			ShortName: "i",
			Usage:     "things with digitalocean images",
			Action: func(c *cli.Context) {

			},
		},
		// sizes
		{
			Name:      "sizes",
			ShortName: "s",
			Usage:     "Stuff with Sizes",
			Action: allSizes,
		},
		// sshkeys
		{
			Name:      "keys",
			ShortName: "k",
			Flags: []cli.Flag{
				cli.IntFlag{"id", -1, "ID of remote ssh key."},
				cli.StringFlag{"name", "", "Name of ssh key."},
				cli.StringFlag{
					"keypath, path",
					"",
					"Path to ssh key on local machine.",
				},
			},
			Usage:  "Stuff with SSHKeys",
			Action: keys,
		},
	}
	app.Run(os.Args)
}
