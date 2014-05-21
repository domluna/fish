package main

import (
	"os"
	"path"
	"strings"
	"fmt"
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
