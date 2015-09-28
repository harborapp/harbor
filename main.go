package main

import (
	"html/template"
	"net/http"
	"os"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/elazarl/go-bindata-assetfs"
	"github.com/gin-gonic/gin"
)

var (
	version string
)

func main() {
	app := cli.NewApp()
	app.Name = "harbor"
	app.Version = version
	app.Author = "Thomas Boerger <thomas@webhippie.de>"
	app.Usage = "A simple docker distribution management web UI"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "db-driver",
			Value:  "sqlite3",
			Usage:  "Driver for the database connection",
			EnvVar: "HARBOR_DB_DRIVER",
		},
		cli.StringFlag{
			Name:   "db-config",
			Value:  "file://harbor.sqlite3",
			Usage:  "Config for the database connection",
			EnvVar: "HARBOR_DB_CONFIG",
		},
	}

	app.Before = func(c *cli.Context) error {
		// TODO(must): Add store stuff
		store, err := store.New(
			c.GlobalString("db-driver"),
			c.GlobalString("db-config"))

		if err != nil {
			return err
		}

		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "Start web UI server",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "port, p",
					Value:  "3000",
					Usage:  "Run on the specified port",
					EnvVar: "HARBOR_SERVER_PORT",
				},
				cli.StringFlag{
					Name:   "listen, l",
					Value:  "0.0.0.0",
					Usage:  "Listen to the specified IP",
					EnvVar: "HARBOR_SERVER_PORT",
				},
				cli.StringFlag{
					Name:   "root, r",
					Value:  "/",
					Usage:  "Run on the specified root",
					EnvVar: "HARBOR_SERVER_ROOT",
				},
				cli.StringFlag{
					Name:   "ca-path",
					Value:  "",
					Usage:  "Path to CA for HTTPS connection",
					EnvVar: "HARBOR_SERVER_CA",
				},
				cli.StringFlag{
					Name:   "cert-path",
					Value:  "",
					Usage:  "Path to cert for HTTPS connection",
					EnvVar: "HARBOR_SERVER_CERT",
				},
				cli.StringFlag{
					Name:   "key-path",
					Value:  "",
					Usage:  "Path to key for HTTPS connection",
					EnvVar: "HARBOR_SERVER_KEY",
				},
			},
			Action: func(c *cli.Context) {
				var err error

				http.Handle(
					staticPath(c.String("root")),
					static())

				http.Handle(
					rootPath(c.String("root")),
					root())

				if len(c.String("cert-path")) == 0 {
					err = http.ListenAndServe(
						strings.Join(
							[]string{
								c.String("listen"),
								c.String("port"),
							},
							":"),
						nil)
				} else {
					// TODO(must): Add CA stuff

					err = http.ListenAndServeTLS(
						strings.Join(
							[]string{
								c.String("listen"),
								c.String("port"),
							},
							":"),
						c.String("cert-path"),
						c.String("key-path"),
						nil)
				}

				if err != nil {
					logrus.Error("Cannot start server: ", err)
				}
			},
		},
	}

	app.Run(os.Args)
}

func staticPath(root string) string {
	// TODO(must): Add root
	return "/static/"
}

func static() http.Handler {
	var handler = http.FileServer(&assetfs.AssetFS{
		Asset:    Asset,
		AssetDir: AssetDir,
		Prefix:   "static",
	})

	return http.StripPrefix("/static/", handler)
}

func rootPath(root string) string {
	// TODO(must): Add root
	return "/"
}

func root() http.Handler {
	r := gin.Default()

	r.SetHTMLTemplate(index())

	r.NoRoute(func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	return r
}

func index() *template.Template {
	file := string(
		MustAsset("static/templates/index.html"))

	return template.Must(
		template.New("index.html").Parse(file))
}
