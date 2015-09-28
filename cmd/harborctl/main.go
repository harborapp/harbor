package main

import (
	"os"

	"github.com/codegangsta/cli"
)

var (
	version string
)

func main() {
	app := cli.NewApp()
	app.Name = "harborctl"
	app.Version = version
	app.Author = "Thomas Boerger <thomas@webhippie.de>"
	app.Usage = "A simple docker distribution management web UI"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "api",
			Value:  "http://localhost:3000",
			Usage:  "Connect to specified API endpoint",
			EnvVar: "HARBORCTL_API",
		},
		cli.StringFlag{
			Name:   "username, u",
			Value:  "",
			Usage:  "Use specified username for the API",
			EnvVar: "HARBORCTL_USERNAME",
		},
		cli.StringFlag{
			Name:   "password, p",
			Value:  "",
			Usage:  "Use specified password for API",
			EnvVar: "HARBORCTL_PASSWORD",
		},
	}

	app.Before = func(c *cli.Context) error {
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:    "registry",
			Aliases: []string{"r"},
			Usage:   "Commands for registry management",
			Subcommands: []cli.Command{
				{
					Name:  "remove, r",
					Usage: "Remove a registry",
					Flags: []cli.Flag{},
					Action: func(c *cli.Context) {

					},
				},
				{
					Name:  "update, u",
					Usage: "Update a registry",
					Flags: []cli.Flag{},
					Action: func(c *cli.Context) {

					},
				},
				{
					Name:  "create, c",
					Usage: "Create a registry",
					Flags: []cli.Flag{},
					Action: func(c *cli.Context) {

					},
				},
			},
		},
		{
			Name:    "user",
			Aliases: []string{"u"},
			Usage:   "Commands for user management",
			Subcommands: []cli.Command{
				{
					Name:  "remove, r",
					Usage: "Remove a user",
					Flags: []cli.Flag{},
					Action: func(c *cli.Context) {

					},
				},
				{
					Name:  "update, u",
					Usage: "Update a user",
					Flags: []cli.Flag{},
					Action: func(c *cli.Context) {

					},
				},
				{
					Name:  "create, c",
					Usage: "Create a user",
					Flags: []cli.Flag{},
					Action: func(c *cli.Context) {

					},
				},
			},
		},
		{
			Name:    "team",
			Aliases: []string{"t"},
			Usage:   "Commands for team management",
			Subcommands: []cli.Command{
				{
					Name:  "remove, r",
					Usage: "Remove a team",
					Flags: []cli.Flag{},
					Action: func(c *cli.Context) {

					},
				},
				{
					Name:  "update, u",
					Usage: "Update a team",
					Flags: []cli.Flag{},
					Action: func(c *cli.Context) {

					},
				},
				{
					Name:  "create, c",
					Usage: "Create a team",
					Flags: []cli.Flag{},
					Action: func(c *cli.Context) {

					},
				},
				{
					Name:  "assign",
					Usage: "Assign a user",
					Flags: []cli.Flag{},
					Action: func(c *cli.Context) {

					},
				},
				{
					Name:  "revoke",
					Usage: "Revoke a user",
					Flags: []cli.Flag{},
					Action: func(c *cli.Context) {

					},
				},
			},
		},
	}

	app.Run(os.Args)
}
