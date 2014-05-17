package main

import (
    "fmt"
    "os"
    "github.com/codegangsta/cli"
    "github.com/Niessy/fisherman/api"
)

func main() {
    app := cli.NewApp()
    app.Name = "fisherman"
    app.Usage = "CLI for DigitalOcean"
    app.Flags = []cli.Flag {
	cli.StringFlag{"apikey", "", "Your api key for DigitalOcean"},
	cli.StringFlag{"clientid", "", "Your client id for DigitalOcean"},
    }

    app.Commands = []cli.Command{
	{
	    Name: "authorize",
	    Usage: "Authorize users DigitalOcean credentails.",
	},
    }
    app.Action = func(c *cli.Context) {
	fmt.Println("Going to be awesome!", api.BaseURL)
    }

    app.Run(os.Args)
}
