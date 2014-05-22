// TODO: figure how to only auth at one point
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/Niessy/fisherman/api/v1"
	"github.com/codegangsta/cli"
)

func keys(c *cli.Context) {
	auth(c)
	args := c.Args()

	if len(args) == 0 {
		allKeys()
		return
	}

	// sub command of sshkeys
	command := args[0]
	switch command {
	case "show":
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fatalf(err.Error())
		}
		showKey(id)
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

	case "rm":
		if c.Int("id") == -1 {
			fatalf("--id flag is unspecified")
		}
		removeKey(c.Int("id"))

	default:
		fatalf("Invalid keys subcommand: try show, add or rm")
	}
}

func allKeys() {
	keys, err := v1.GetSSHKeys()
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
	f := cleanPath(keypath, "/")
	buf, err := ioutil.ReadFile(f)
	if err != nil {
		log.Fatal(err)
	}

	key, err := v1.AddSSHKey(name, string(buf))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Succesfully added key %s to remote sshkeys\n", key.Name)
}

func removeKey(id int) {
	err := v1.DestroySSHKey(id)
	if err != nil {
		fatalf(err.Error())
	}
	fmt.Printf("Succesfully destroyed ssh key (ID: %d)\n", id)
}

func showKey(id int) {
	key, err := v1.GetSSHKey(id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", key.SSHPublicKey)
}