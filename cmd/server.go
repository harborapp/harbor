package cmd

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"path"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/umschlag/umschlag-api/config"
	"github.com/umschlag/umschlag-api/router"
	"github.com/umschlag/umschlag-api/shared/s3client"
	"github.com/urfave/cli"
	"golang.org/x/crypto/acme/autocert"
)

// Server provides the sub-command to start the API server.
func Server() cli.Command {
	return cli.Command{
		Name:  "server",
		Usage: "Start the Umschlag API",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "db-driver",
				Value:       "mysql",
				Usage:       "Database driver selection",
				EnvVar:      "UMSCHLAG_DB_DRIVER",
				Destination: &config.Database.Driver,
			},
			cli.StringFlag{
				Name:        "db-name",
				Value:       "umschlag",
				Usage:       "Name for database connection",
				EnvVar:      "UMSCHLAG_DB_NAME",
				Destination: &config.Database.Name,
			},
			cli.StringFlag{
				Name:        "db-username",
				Value:       "root",
				Usage:       "Username for database connection",
				EnvVar:      "UMSCHLAG_DB_USERNAME",
				Destination: &config.Database.Username,
			},
			cli.StringFlag{
				Name:        "db-password",
				Value:       "root",
				Usage:       "Password for database connection",
				EnvVar:      "UMSCHLAG_DB_PASSWORD",
				Destination: &config.Database.Password,
			},
			cli.StringFlag{
				Name:        "db-host",
				Value:       "localhost:3306",
				Usage:       "Host for database connection",
				EnvVar:      "UMSCHLAG_DB_HOST",
				Destination: &config.Database.Host,
			},
			cli.StringFlag{
				Name:        "host",
				Value:       "http://localhost:8080",
				Usage:       "External access to server",
				EnvVar:      "UMSCHLAG_SERVER_HOST",
				Destination: &config.Server.Host,
			},
			cli.StringFlag{
				Name:        "addr",
				Value:       ":8080",
				Usage:       "Address to bind the server",
				EnvVar:      "UMSCHLAG_SERVER_ADDR",
				Destination: &config.Server.Addr,
			},
			cli.StringFlag{
				Name:        "root",
				Value:       "/",
				Usage:       "Root folder of the app",
				EnvVar:      "UMSCHLAG_SERVER_ROOT",
				Destination: &config.Server.Root,
			},
			cli.StringFlag{
				Name:        "storage",
				Value:       "storage/",
				Usage:       "Folder for storing uploads",
				EnvVar:      "UMSCHLAG_SERVER_STORAGE",
				Destination: &config.Server.Storage,
			},
			cli.BoolFlag{
				Name:        "pprof",
				Usage:       "Enable pprof debugging server",
				EnvVar:      "UMSCHLAG_SERVER_PPROF",
				Destination: &config.Server.Pprof,
			},
			cli.StringFlag{
				Name:        "cert",
				Value:       "",
				Usage:       "Path to SSL cert",
				EnvVar:      "UMSCHLAG_SERVER_CERT",
				Destination: &config.Server.Cert,
			},
			cli.StringFlag{
				Name:        "key",
				Value:       "",
				Usage:       "Path to SSL key",
				EnvVar:      "UMSCHLAG_SERVER_KEY",
				Destination: &config.Server.Key,
			},
			cli.BoolFlag{
				Name:        "letsencrypt",
				Usage:       "Enable Let's Encrypt SSL",
				EnvVar:      "UMSCHLAG_SERVER_LETSENCRYPT",
				Destination: &config.Server.LetsEncrypt,
			},
			cli.DurationFlag{
				Name:        "expire",
				Value:       time.Hour * 24,
				Usage:       "Session expire duration",
				EnvVar:      "UMSCHLAG_SESSION_EXPIRE",
				Destination: &config.Session.Expire,
			},
			cli.StringSliceFlag{
				Name:   "admin-user",
				Value:  &cli.StringSlice{},
				Usage:  "Enforce user as an admin",
				EnvVar: "UMSCHLAG_ADMIN_USERS",
			},
			cli.BoolTFlag{
				Name:        "admin-create",
				Usage:       "Create an initial admin user",
				EnvVar:      "UMSCHLAG_ADMIN_CREATE",
				Destination: &config.Admin.Create,
			},
			cli.BoolFlag{
				Name:        "s3-enabled",
				Usage:       "Enable S3 uploads",
				EnvVar:      "UMSCHLAG_S3_ENABLED",
				Destination: &config.S3.Enabled,
			},
			cli.StringFlag{
				Name:        "s3-endpoint",
				Value:       "",
				Usage:       "S3 API endpoint",
				EnvVar:      "UMSCHLAG_S3_ENDPOINT",
				Destination: &config.S3.Endpoint,
			},
			cli.StringFlag{
				Name:        "s3-bucket",
				Value:       "umschlag",
				Usage:       "S3 bucket name",
				EnvVar:      "UMSCHLAG_S3_BUCKET",
				Destination: &config.S3.Bucket,
			},
			cli.StringFlag{
				Name:        "s3-region",
				Value:       "us-east-1",
				Usage:       "S3 region name",
				EnvVar:      "UMSCHLAG_S3_REGION",
				Destination: &config.S3.Region,
			},
			cli.StringFlag{
				Name:        "s3-access",
				Value:       "",
				Usage:       "S3 public key",
				EnvVar:      "UMSCHLAG_S3_ACCESS_KEY",
				Destination: &config.S3.Access,
			},
			cli.StringFlag{
				Name:        "s3-secret",
				Value:       "",
				Usage:       "S3 secret key",
				EnvVar:      "UMSCHLAG_S3_SECRET_KEY",
				Destination: &config.S3.Secret,
			},
			cli.BoolFlag{
				Name:        "s3-path-style",
				Usage:       "S3 path style",
				EnvVar:      "UMSCHLAG_S3_PATH_STYLE",
				Destination: &config.S3.PathStyle,
			},
		},
		Before: func(c *cli.Context) error {
			if len(c.StringSlice("admin-user")) > 0 {
				// StringSliceFlag doesn't support Destination
				config.Admin.Users = c.StringSlice("admin-user")
			}

			if config.S3.Enabled {
				_, err := s3client.New().List()

				if err != nil {
					return fmt.Errorf("Failed to connect to S3. %s", err)
				}
			}

			return nil
		},
		Action: func(c *cli.Context) {
			logrus.Infof("Starting API on %s", config.Server.Addr)

			var (
				server *http.Server
			)

			if config.Server.LetsEncrypt || (config.Server.Cert != "" && config.Server.Key != "") {
				curves := []tls.CurveID{
					tls.CurveP521,
					tls.CurveP384,
					tls.CurveP256,
				}

				ciphers := []uint16{
					tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
					tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
					tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
					tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				}

				cfg := &tls.Config{
					PreferServerCipherSuites: true,
					MinVersion:               tls.VersionTLS12,
					CurvePreferences:         curves,
					CipherSuites:             ciphers,
				}

				if config.Server.LetsEncrypt {
					certManager := autocert.Manager{
						Prompt: autocert.AcceptTOS,
						Cache:  autocert.DirCache(path.Join(config.Server.Storage, "certs")),
					}

					cfg.GetCertificate = certManager.GetCertificate
				} else {
					cert, err := tls.LoadX509KeyPair(
						config.Server.Cert,
						config.Server.Key,
					)

					if err != nil {
						logrus.Fatal("Failed to load SSL certificates. %s", err)
					}

					cfg.Certificates = []tls.Certificate{
						cert,
					}
				}

				server = &http.Server{
					Addr:         config.Server.Addr,
					Handler:      router.Load(),
					ReadTimeout:  5 * time.Second,
					WriteTimeout: 10 * time.Second,
					TLSConfig:    cfg,
				}
			} else {
				server = &http.Server{
					Addr:         config.Server.Addr,
					Handler:      router.Load(),
					ReadTimeout:  5 * time.Second,
					WriteTimeout: 10 * time.Second,
				}
			}

			if err := startServer(server); err != nil {
				logrus.Fatal(err)
			}
		},
	}
}
