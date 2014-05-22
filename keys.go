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

func showKey(id int) {
	key, err := v1.GetSSHKey(id)
	if err != nil {
		fatalf(err.Error())
	}
	fmt.Printf("%s\n", key.SSHPublicKey)
}

func addKey(c *cli.Context) {
	auth(c)
	args := c.Args()

	if len(args) < 1 {
		fatalf("first argument is the key name")
	}

	name := args[0]
	kp := c.String("path")

	if kp == "" {
		fatalf("public key path not specified, use --path=$YOURKEYPATH")
	}

	f := cleanPath(kp, "/")
	buf, err := ioutil.ReadFile(f)
	if err != nil {
		fatalf(err.Error())
	}

	key, err := v1.AddSSHKey(name, string(buf))
	if err != nil {
		fatalf(err.Error())
	}
	fmt.Printf("added key %s to remote ssh keys\n", key.Name)
}

func rmKey(c *cli.Context) {
	auth(c)
	args := c.Args()

	if len(args) < 1 {
		fatalf("first argument is the key name")
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fatalf(err.Error())
	}

	println(id)

	err = v1.DestroySSHKey(id)
	if err != nil {
		fatalf(err.Error())
	}
	fmt.Printf("queued removal of ssh key (id: %d)\n", id)
}
