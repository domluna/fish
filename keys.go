// TODO: figure how to only auth at one point
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"github.com/codegangsta/cli"
	"github.com/Niessy/dogo"
)

func keys(c *cli.Context) {
	auth(c)

	args := c.Args()
	// show all keys then exit
	if len(args) == 0 {
		allKeys()
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fatalf(err.Error())
	}
	showKey(id)

}

func allKeys() {
	keys, err := dogo.GetSSHKeys()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("SSHKeys: ")
	for _, k := range keys {
		fmt.Printf("%s (id: %d)\n", k.Name, k.ID)
	}
	fmt.Println()
}

func showKey(id int) {
	key, err := dogo.GetSSHKey(id)
	if err != nil {
		fatalf(err.Error())
	}
	fmt.Printf("%s\n", key.SSHPublicKey)
}

func addKey(c *cli.Context) {
	auth(c)
	args := c.Args()

	if len(args) < 1 {
		cli.ShowCommandHelp(c, c.Command.Name)
		return
	}

	name := args[0]
	kp := c.String("path")

	if kp == "" {
		fatalf("Public key path not specified, use --path=$YOURKEYPATH")
	}

	f := cleanPath(kp, "/")
	buf, err := ioutil.ReadFile(f)
	if err != nil {
		fatalf(err.Error())
	}

	key, err := dogo.AddSSHKey(name, string(buf))
	if err != nil {
		fatalf(err.Error())
	}
	fmt.Printf("Added key %s to remote ssh keys\n", key.Name)
}

func rmKey(c *cli.Context) {
	auth(c)
	args := c.Args()

	if len(args) < 1 {
		cli.ShowCommandHelp(c, c.Command.Name)
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fatalf(err.Error())
	}

	println(id)

	err = dogo.DestroySSHKey(id)
	if err != nil {
		fatalf(err.Error())
	}
	fmt.Printf("Queued removal of ssh key (id: %d)\n", id)
}
