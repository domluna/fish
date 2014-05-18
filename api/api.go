package api

import (
	"fmt"
)

var version = "v1"
var hostname = "https://api.digitalocean.com"

func formatURL(resource string) string {
	return fmt.Sprintf("%s/%s/%s", hostname, version, resource)
}
