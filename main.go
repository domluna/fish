//
// CLI for DigitalOcean
//
package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/Niessy/dogo"
)

var docli *dogo.Client

func init() {
	// Load the configuration in $HOME/.fish
	loadConfig()
	auth := dogo.Auth{config.Conf.ClientID, config.Conf.APIKey}
	docli = dogo.NewClient(auth)

}

func main() {
	app := cli.NewApp()
	app.Name = "fish"
	app.Author = "Dominique Luna"
	app.Usage = "Command Line Interface for DigitalOcean"
	app.Commands = []cli.Command{
		// droplets
		{
			Name:   "droplets",
			Usage:  "List user all droplets",
			Action: droplets,
		},
		// regions
		{
			Name:   "regions",
			Usage:  "List all regions",
			Action: regions,
		},
		// images
		{
			Name:  "images",
			Usage: "Lists all global or user images",
			Flags: []cli.Flag{
				cli.BoolFlag{"global, g", "show all DigitalOcean images"},
			},
			Action: images,
		},
		// sizes
		{
			Name:   "sizes",
			Usage:  "Lists all available droplet sizes",
			Action: sizes,
		},
		// ssh keys
		{
			Name:   "keys",
			Usage:  "Lists all user ssh keys",
			Action: keys,
		},
		// add ssh key
		{
			Name:        "addkey",
			Usage:       "Add an ssh key",
			Description: "First arg is the key name",
			Flags: []cli.Flag{
				cli.StringFlag{"path, p", "", "path to the public key"},
			},
			Action: addKey,
		},
		// remove ssh key
		{
			Name:        "rmkey",
			Usage:       "Remove an ssh key",
			Description: "First arg is the key id",
			Action:      rmKey,
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
			Name:        "create",
			Usage:       "Create a new droplet",
			Description: "First arg is the name of the droplet",
			Flags: []cli.Flag{
				cli.IntFlag{"image, i", 3668014, "image id, default: Docker Image"},
				cli.StringFlag{"size, s", "512MB", "size slug, ex: 512MB, 1GB"},
				cli.StringFlag{"region, r", "nyc2", "region slug, ex: nyc2, ams1"},
				cli.BoolFlag{"network, n", "enable private networking"},
				cli.BoolFlag{"backups, b", "enable backups"},
				cli.IntSliceFlag{"keys, k", &cli.IntSlice{}, "ssh key ids"},
			},
			Action: create,
		},
		// destroy droplet
		{
			Name:        "destroy",
			Description: "First arg is the droplet id",
			Usage:       "Destroy a droplet",
			Action:      destroy,
		},
		// resize droplet
		{
			Name:        "resize",
			Usage:       "Resize a droplet",
			Description: "First arg is the droplet id",
			Flags: []cli.Flag{
				cli.StringFlag{"size, s", "", "size slug, Ex. 512MB, 1GB"},
			},
			Action: resize,
		},
		// reboot droplet
		{
			Name:   "reboot",
			Usage:  "Reboot a droplet",
			Action: reboot,
		},
		// rebuild droplet
		{
			Name:  "rebuild",
			Usage: "Rebuilds a droplet with a default image, keeps ip address",
			Flags: []cli.Flag{
				cli.IntFlag{"image, i", 0, "Image id to rebuild"},
			},
			Action: rebuild,
		},
		// stop droplet
		{
			Name:   "off",
			Usage:  "Power off a droplet",
			Action: off,
		},
		// start droplet
		{
			Name:   "on",
			Usage:  "Power on a droplet",
			Action: on,
		},
		// snapshot droplet
		{
			Name:  "snapshot",
			Usage: "Take a snapshot of a current droplet",
			Flags: []cli.Flag{
				cli.StringFlag{"name, n", "", "Name of the snapshot image"},
			},
			Action: snapshot,
		},
		// restore droplet
		{
			Name:  "restore",
			Usage: "Restores a droplet with a previous image/snapshot",
			Flags: []cli.Flag{
				cli.IntFlag{"image, i", 0, "Image id to restore"},
			},
			Action: restore,
		},
		// info droplet
		{
			Name:   "info",
			Usage:  "Detailed info about a droplet",
			Action: info,
		},
	}
	app.Run(os.Args)
}
