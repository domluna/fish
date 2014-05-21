package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/Niessy/fisherman/api/v1"
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

// For authenticating the configuration variables,
// client id and api key mainly.
func auth(c *cli.Context) {
	path := c.String("configFile")
	err := v1.LoadConfig(path)
	if err != nil {
		fatalf(err.Error())
	}
}
