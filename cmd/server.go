package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/harborapp/harbor-api/config"
	"github.com/harborapp/harbor-api/model"
	"github.com/harborapp/harbor-api/router"
	"github.com/harborapp/harbor-api/router/middleware/context"
	"github.com/harborapp/harbor-api/server"
)

// Server provides the sub-command to start the API server.
func Server(cfg *config.Config) cli.Command {
	return cli.Command{
		Name:  "server",
		Usage: "Start the Harbor server",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "db-driver",
				Value:       "mysql",
				Usage:       "Database driver selection",
				EnvVar:      "HARBOR_DB_DRIVER",
				Destination: &cfg.Database.Driver,
			},
			cli.StringFlag{
				Name:        "db-name",
				Value:       "harbor",
				Usage:       "Name for database connection",
				EnvVar:      "HARBOR_DB_NAME",
				Destination: &cfg.Database.Name,
			},
			cli.StringFlag{
				Name:        "db-username",
				Value:       "root",
				Usage:       "Username for database connection",
				EnvVar:      "HARBOR_DB_USERNAME",
				Destination: &cfg.Database.Username,
			},
			cli.StringFlag{
				Name:        "db-password",
				Value:       "root",
				Usage:       "Password for database connection",
				EnvVar:      "HARBOR_DB_PASSWORD",
				Destination: &cfg.Database.Password,
			},
			cli.StringFlag{
				Name:        "db-host",
				Value:       "localhost:3306",
				Usage:       "Host for database connection",
				EnvVar:      "HARBOR_DB_HOST",
				Destination: &cfg.Database.Host,
			},
			cli.StringFlag{
				Name:        "addr",
				Value:       ":8080",
				Usage:       "Address to bind the server",
				EnvVar:      "HARBOR_SERVER_ADDR",
				Destination: &cfg.Server.Addr,
			},
			cli.StringFlag{
				Name:        "cert",
				Value:       "",
				Usage:       "Path to SSL cert",
				EnvVar:      "HARBOR_SERVER_CERT",
				Destination: &cfg.Server.Cert,
			},
			cli.StringFlag{
				Name:        "key",
				Value:       "",
				Usage:       "Path to SSL key",
				EnvVar:      "HARBOR_SERVER_KEY",
				Destination: &cfg.Server.Key,
			},
			cli.StringFlag{
				Name:        "root",
				Value:       "/",
				Usage:       "Root folder of the app",
				EnvVar:      "HARBOR_SERVER_ROOT",
				Destination: &cfg.Server.Root,
			},
			cli.BoolFlag{
				Name:        "debug",
				Usage:       "Activate debug information",
				EnvVar:      "HARBOR_SERVER_DEBUG",
				Destination: &cfg.Debug,
			},
		},
		Action: func(c *cli.Context) {
			dat := model.Load(
				cfg,
			)

			srv := server.Load(
				cfg,
			)

			srv.Run(
				router.Load(
					cfg,
					context.SetConfig(
						*cfg,
					),
					context.SetStore(
						*dat,
					),
				),
			)
		},
	}
}
