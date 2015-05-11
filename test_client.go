package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/nesv/go-dynect/dynect"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "backup_dyn_zones"
	app.Usage = "Dyn zone file backup tool"
	app.Action = testClient
	app.HideVersion = true

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "client, c",
			Usage: "Your Dyn client name",
		},
		cli.StringFlag{
			Name:  "username, u",
			Usage: "Your Dyn username",
		},
		cli.StringFlag{
			Name:  "password, p",
			Usage: "Your Dyn password",
		},
		cli.BoolFlag{
			Name:  "verbose, v",
			Usage: "Verbose Dyn client",
		},
	}

	app.Run(os.Args)
}

func testClient(c *cli.Context) {
	client := dynect.NewClient(c.String("client"))

	client.Verbose(c.Bool("verbose"))

	err := client.Login(c.String("username"), c.String("password"))
	if err != nil {
		panic(err)
	}

	defer func() {
		err := client.Logout()
		if err != nil {
			panic(err)
		}
	}()

	var response dynect.ZonesResponse

	if err := client.Do("GET", "Zone", nil, &response); err != nil {
		panic(err)
	}

	fmt.Println(response.Data)
}
