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
			Usage:  "List user all droplets",
			Action: allDroplets,
		},
		// regions
		{
			Name:   "regions",
			Usage:  "List all regions",
			Action: allRegions,
		},
		// images
		{
			Name:  "images",
			Usage: "Lists all global or user images",
			Flags: []cli.Flag{
				cli.BoolFlag{"global, g", "show all DigitalOcean images"},
			},
			Action: allImages,
		},
		// sizes
		{
			Name:   "sizes",
			Usage:  "Lists all available droplet sizes",
			Action: allSizes,
		},
		// ssh keys
		{
			Name:   "keys",
			Usage:  "Lists all user ssh keys",
			Action: keys,
		},
		// add ssh key
		{
			Name:  "addkey",
			Usage: "Add an ssh key",
			Flags: []cli.Flag{
				cli.StringFlag{"path, p", "", "path to the public key"},
			},
			Action: addKey,
		},
		// remove ssh key
		{
			Name:   "rmkey",
			Usage:  "Remove an ssh key",
			Action: rmKey,
		},
		// ssh access to digitalocean servers
		{
			Name:  "ssh",
			Usage: "Access DigitalOcean servers via ssh",
			Action: func(c *cli.Context) {
				println("Not implemented yet!")
			},
		},
		// create droplet
		{
			Name:  "create",
			Usage: "Create a new droplet",
			Action: func(c *cli.Context) {
				println("Not implemented yet!")
			},
		},
		// destroy droplet
		{
			Name:  "destroy",
			Usage: "Destroy a droplet",
			Action: func(c *cli.Context) {
				println("Not implemented yet!")
			},
		},
		// resize droplet
		{
			Name:  "resize",
			Usage: "Resize a droplet",
			Action: func(c *cli.Context) {
				println("Not implemented yet!")
			},
		},
		// reboot droplet
		{
			Name:  "reboot",
			Usage: "Reboot a droplet",
			Action: func(c *cli.Context) {
				println("Not implemented yet!")
			},
		},
		// rebuild droplet
		{
			Name:  "rebuild",
			Usage: "Rebuilds a droplet with a default image, keeps ip address",
			Action: func(c *cli.Context) {
				println("Not implemented yet!")
			},
		},
		// shutdown droplet
		{
			Name:  "stop",
			Usage: "Shutdown a droplet",
			Action: func(c *cli.Context) {
				println("Not implemented yet!")
			},
		},
		// start droplet
		{
			Name:  "start",
			Usage: "Start up a droplet",
			Action: func(c *cli.Context) {
				println("Not implemented yet!")
			},
		},
		// snapshot droplet
		{
			Name:  "snapshot",
			Usage: "Snapshot a droplet",
			Action: func(c *cli.Context) {
				println("Not implemented yet!")
			},
		},
		// info droplet
		{
			Name:  "info",
			Usage: "Gives detailed droplet info",
			Action: func(c *cli.Context) {
				println("Not implemented yet!")
			},
		},
	}
	app.Run(os.Args)
}
