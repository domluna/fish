// TODO: figure how to only auth at one point
package main

import (
	"fmt"
	"log"
	"io/ioutil"

	"github.com/Niessy/fisherman/api/v1"
	"github.com/codegangsta/cli"
)

func auth(c *cli.Context) {
	path := c.String("configFile")
	err := v1.LoadConfig(path)
	if err != nil {
		log.Fatal(err)
	}
}

func allDroplets(c *cli.Context) {
	auth(c)
	v1.GetDroplets()
}

func allRegions(c *cli.Context) {
	auth(c)
	v1.GetRegions()
}

func allImages(c *cli.Context) {
	auth(c)
	v1.GetImages()
}

func allSizes(c *cli.Context) {
	auth(c)
	v1.GetSizes()
}

func keys(c *cli.Context) {
	auth(c)
	args := c.Args()

	if len(args) != 1 {
		fatalf("Subcommands include add, show, and destroy")
	}

	// sub command of sshkeys
	command := args[0]
	switch command {
	case "show":
		if c.Int("id") == -1 {
			allKeys()
			return
		}
		showKey(c.Int("id"))

	case "add":
		keypath := c.String("keypath")
		name := c.String("name")
		if keypath == "" {
			fatalf("--keypath flag is unspecified")
		}

		if name == "" {
			fatalf("--name flag is unspecified")
		}

		addKey(name, keypath)

	case "destroy":
		if c.Int("id") == -1 {
			fatalf("--id flag is unspecified")
		}
		destroyKey(c.Int("id"))

	default:
		fatalf("Invalid subcommand")
	}
}

func allKeys() {
	keys, err := v1.AllSSHKeys()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("SSHKeys: ")
	for _, k := range keys {
		fmt.Printf("%s (id: %d)\n", k.Name, k.ID)
	}
	fmt.Println()
}

func addKey(name, keypath string) {
	file := cleanPath(keypath, "/")
	println(keypath)
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	key, err := v1.AddSSHKey(name, string(buf))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Succesfully added key %s to remote sshkeys\n", key.Name)
}

func destroyKey(id int) {
	err := v1.DestroySSHKey(id)
	if err != nil {
		fatalf(err.Error())
	}
	fmt.Printf("Succesfully destroyed ssh key (ID: %d)\n", id)
}

func showKey(id int) {
	key, err := v1.ShowSSHKey(id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", key.SSHPublicKey)
}
