package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/Sirupsen/logrus"
	"github.com/sanbornm/go-selfupdate/selfupdate"
	"github.com/umschlag/umschlag-api/cmd"
	"github.com/umschlag/umschlag-api/config"
	"github.com/urfave/cli"

	_ "github.com/joho/godotenv/autoload"
)

var (
	updates = "http://dl.webhippie.de/"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	app := cli.NewApp()
	app.Name = "umschlag"
	app.Version = config.Version
	app.Author = "Thomas Boerger <thomas@webhippie.de>"
	app.Usage = "A docker distribution management system"

	app.Flags = []cli.Flag{
		cli.BoolTFlag{
			Name:        "update, u",
			Usage:       "Enable auto update",
			EnvVar:      "UMSCHLAG_UPDATE",
			Destination: &config.Debug,
		},
		cli.BoolFlag{
			Name:        "debug",
			Usage:       "Activate debug information",
			EnvVar:      "UMSCHLAG_DEBUG",
			Destination: &config.Debug,
		},
	}

	app.Before = func(c *cli.Context) error {
		logrus.SetOutput(os.Stdout)

		if config.Debug {
			logrus.SetLevel(logrus.DebugLevel)
		} else {
			logrus.SetLevel(logrus.InfoLevel)
		}

		if c.BoolT("update") {
			Update()
		}

		return nil
	}

	app.Commands = []cli.Command{
		cmd.Server(),
	}

	cli.HelpFlag = cli.BoolFlag{
		Name:  "help, h",
		Usage: "Show the help, so what you see now",
	}

	cli.VersionFlag = cli.BoolFlag{
		Name:  "version, v",
		Usage: "Print the current version of that tool",
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Update handles automated binary updates in the background.
func Update() {
	if config.VersionDev == "dev" {
		fmt.Fprintf(os.Stderr, "Updates are disabled for development versions.\n")
	} else {
		updater := &selfupdate.Updater{
			CurrentVersion: fmt.Sprintf(
				"%d.%d.%d",
				config.VersionMajor,
				config.VersionMinor,
				config.VersionPatch,
			),
			ApiURL:  updates,
			BinURL:  updates,
			DiffURL: updates,
			Dir:     "updates/",
			CmdName: "umschlag-api",
		}

		go updater.BackgroundRun()
	}
}
