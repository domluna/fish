// TODO: figure how to only auth at one point
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/codegangsta/cli"
)

func keys(c *cli.Context) {
	keys, err := docli.GetSSHKeys()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("SSHKeys: ")
	for _, k := range keys {
		fmt.Printf("%s (id: %d)\n", k.Name, k.ID)
	}
}

func showKey(c *cli.Context) {
	checkArgs(c)

	id, err := strconv.Atoi(c.Args().First())
	if err != nil {
		fatalf(err.Error())
	}

	key, err := docli.GetSSHKey(id)
	if err != nil {
		fatalf(err.Error())
	}
	fmt.Printf("%s\n", key.SSHPublicKey)
}

func addKey(c *cli.Context) {
	checkArgs(c)

	name := c.Args().First()
	kp := c.String("path")

	if kp == "" {
		fatalf("Public key path not specified, use --path=$YOURKEYPATH")
	}

	f := cleanPath(kp, "/")
	buf, err := ioutil.ReadFile(f)
	if err != nil {
		fatalf(err.Error())
	}

	key, err := docli.AddSSHKey(name, buf)
	if err != nil {
		fatalf(err.Error())
	}
	fmt.Printf("Added key %s to remote ssh keys\n", key.Name)
}

func rmKey(c *cli.Context) {
	checkArgs(c)

	id, err := strconv.Atoi(c.Args().First())
	if err != nil {
		fatalf(err.Error())
	}

	println(id)

	err = docli.DestroySSHKey(id)
	if err != nil {
		fatalf(err.Error())
	}
	fmt.Printf("Queued removal of ssh key\n")
}
