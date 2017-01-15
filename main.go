package main

import (
	"os"
	"runtime"

	"github.com/Sirupsen/logrus"
	"github.com/joho/godotenv"
	"github.com/umschlag/umschlag-api/cmd"
	"github.com/umschlag/umschlag-api/config"
	"github.com/urfave/cli"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if env := os.Getenv("UMSCHLAG_ENV_FILE"); env != "" {
		godotenv.Load(env)
	}

	app := cli.NewApp()
	app.Name = "umschlag"
	app.Version = config.Version
	app.Author = "Thomas Boerger <thomas@webhippie.de>"
	app.Usage = "A docker distribution management system"

	app.Flags = []cli.Flag{
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
		logrus.Error(err)
		os.Exit(1)
	}
}
