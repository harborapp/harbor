package main

import (
	"os"
	"runtime"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/joho/godotenv"
	"github.com/umschlag/umschlag-api/config"
	"github.com/umschlag/umschlag-api/pkg/version"
	"gopkg.in/urfave/cli.v2"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if env := os.Getenv("UMSCHLAG_ENV_FILE"); env != "" {
		godotenv.Load(env)
	}

	app := &cli.App{
		Name:     "umschlag-api",
		Version:  version.Version.String(),
		Usage:    "A docker distribution management system",
		Compiled: time.Now(),

		Authors: []*cli.Author{
			{
				Name:  "Thomas Boerger",
				Email: "thomas@webhippie.de",
			},
		},

		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "debug",
				Value:       false,
				Usage:       "Activate debug information",
				EnvVars:     []string{"UMSCHLAG_DEBUG"},
				Destination: &config.Debug,
				Hidden:      true,
			},
		},

		Before: func(c *cli.Context) error {
			logrus.SetOutput(os.Stdout)

			if config.Debug {
				logrus.SetLevel(logrus.DebugLevel)
			} else {
				logrus.SetLevel(logrus.InfoLevel)
			}

			return nil
		},

		Commands: []*cli.Command{
			Server(),
		},
	}

	cli.HelpFlag = &cli.BoolFlag{
		Name:    "help",
		Aliases: []string{"h"},
		Usage:   "Show the help, so what you see now",
	}

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "Print the current version of that tool",
	}

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
