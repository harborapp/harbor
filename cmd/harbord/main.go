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

	"github.com/webhippie/harbor/pkg/server"
	"github.com/webhippie/harbor/pkg/store"
)

var (
	version string
)

var conf = struct {
	debug bool

	database struct {
		driver string
		config string
	}

	server struct {
		listen string
		port   string
		root   string
		cert   string
		key    string
	}
}{}

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
		cli.BoolFlag{
			Name:   "debug",
			Usage:  "Run the commands with debug mode",
			EnvVar: "HARBOR_DEBUG",
		},
	}

	app.Before = func(c *cli.Context) error {
		conf.debug = c.GlobalBool("debug")

		conf.database.driver = c.GlobalString("db-driver")
		conf.database.config = c.GlobalString("db-driver")

		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "Start web UI server",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "listen, l",
					Value:  "0.0.0.0",
					Usage:  "Listen to the specified IP",
					EnvVar: "HARBOR_SERVER_PORT",
				},
				cli.StringFlag{
					Name:   "port, p",
					Value:  "3000",
					Usage:  "Run on the specified port",
					EnvVar: "HARBOR_SERVER_PORT",
				},
				cli.StringFlag{
					Name:   "root, r",
					Value:  "/",
					Usage:  "Run on the specified root",
					EnvVar: "HARBOR_SERVER_ROOT",
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
			Before: func(c *cli.Context) error {
				conf.server.listen = c.String("listen")
				conf.server.port = c.String("port")
				conf.server.root = c.String("root")
				conf.server.cert = c.String("cert-path")
				conf.server.key = c.String("key-path")

				return nil
			},
			Action: func(c *cli.Context) {
				var err error

				listen := strings.Join(
					[]string{
						conf.server.listen,
						conf.server.port,
					},
					":")

				http.Handle(
					staticPath(),
					static())

				http.Handle(
					rootPath(),
					root())

				if len(conf.server.cert) > 0 && len(conf.server.key) > 0 {
					err = http.ListenAndServeTLS(
						listen,
						conf.server.cert,
						conf.server.key,
						nil)
				} else {
					err = http.ListenAndServe(
						listen,
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

func staticPath() string {
	return strings.Join(
		[]string{
			conf.server.root,
			"static/",
		},
		"")
}

func static() http.Handler {
	var handler = http.FileServer(&assetfs.AssetFS{
		Asset:    Asset,
		AssetDir: AssetDir,
		Prefix:   "cmd/harbord/static",
	})

	if conf.debug {
		handler = http.FileServer(
			http.Dir("cmd/harbord/static"))
	}

	return http.StripPrefix(
		staticPath(),
		handler)
}

func rootPath() string {
	return conf.server.root
}

func root() http.Handler {
	store, err := store.New(
		conf.database.driver,
		conf.database.config)

	if err != nil {
		logrus.Fatal(err)
	}

	r := gin.Default()

	api := r.Group("/api")
	api.Use(server.SetStore(store))
	api.Use(server.SetUser())
	api.Use(server.SetHeaders())
	api.OPTIONS("/*path", func(c *gin.Context) {})

	profile := api.Group("/profile")
	{
		profile.Use(server.MustUser())
		profile.GET("", server.GetProfile)
	}

	registries := api.Group("/registries")
	{
		registries.Use(server.MustAdmin())
		registries.GET("", server.GetRegistries)
	}

	users := api.Group("/users")
	{
		users.Use(server.MustAdmin())
		users.GET("", server.GetUsers)
	}

	teams := api.Group("/teams")
	{
		teams.Use(server.MustAdmin())
		teams.GET("", server.GetTeams)
	}

	r.SetHTMLTemplate(
		index())

	r.NoRoute(func(c *gin.Context) {
		c.HTML(
			200,
			"index.html",
			gin.H{
				"root": conf.server.root,
			})
	})

	return r
}

func index() *template.Template {
	file := string(
		MustAsset("cmd/harbord/static/index.html"))

	return template.Must(
		template.New("index.html").Parse(file))
}
