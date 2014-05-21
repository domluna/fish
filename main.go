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
			Usage:     "Stuff with Droplets",
			Action:    allDroplets,
		},
		// regions
		{
			Name:      "regions",
			ShortName: "r",
			Usage:     "Stuff with Regions",
			Action:    allRegions,
		},
		// images
		{
			Name:      "images",
			ShortName: "i",
			Usage:     "Stuff with Images",
			Action:    allImages,
		},
		// sizes
		{
			Name:      "sizes",
			ShortName: "s",
			Usage:     "Stuff with Sizes",
			Action:    allSizes,
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
