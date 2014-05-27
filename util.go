package main

import (
	"fmt"
	"os"
	"path"
	"strings"
	"github.com/codegangsta/cli"
)

// Cleans a filepath and replaces enviroment variables
// with their mapping
func cleanPath(fp, sep string) string {
	cleaned := path.Clean(fp)
	spl := strings.Split(cleaned, sep)
	for i, s := range spl {
		spl[i] = os.ExpandEnv(s)
	}
	return strings.Join(spl, sep)
}

// For logging
func fatalf(format string, v ...interface{}) {
	println(fmt.Sprintf(format, v...))
	os.Exit(1)
}

// checkArgs makes sure there's a first argument. If used
// incorrectly will output help and exit.
func checkArgs(c *cli.Context) {
	if len(c.Args()) < 1 {
		fmt.Printf("Arguments missing\n")
		cli.ShowCommandHelp(c, c.Command.Name)
		os.Exit(1)
	}
}
