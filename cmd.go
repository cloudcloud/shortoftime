// Package main provides the core binary for shortofti.me
package main

import (
	"log"
	"os"

	"github.com/cloudcloud/shortoftime/config"
	"github.com/cloudcloud/shortoftime/server"
	"github.com/gin-gonic/gin"
	"github.com/namsral/flag"
)

func main() {
	var err error
	c := &config.Config{}

	// load any flag detail
	c = loadFlags(os.Args, "SHORT", flag.PanicOnError, c)

	// bring in the standard config file too
	var conf config.Conf
	conf, err = c.LoadFromFile()
	if err != nil {
		log.Printf("An error occured during the config file processing: [%s]", err)
	}

	err = serve(conf, "")
	if err != nil {
		log.Fatalf("Server had issues: [%s]", err)
	}
}

func loadFlags(a []string, p string, f flag.ErrorHandling, c *config.Config) *config.Config {
	flags := flag.NewFlagSetWithEnvPrefix(a[0], p, f)
	flags.StringVar(&c.ConfigFile, "config", "/var/shortoftime/config/config.conf", "Configuration file")
	flags.StringVar(&c.Database.Host, "db-host", "localhost", "Database hostname")
	flags.StringVar(&c.Database.Name, "db-name", "shortoftime", "Database name")
	flags.StringVar(&c.Database.Pass, "db-pass", "", "Database password")
	flags.StringVar(&c.Database.User, "db-user", "root", "Database username")
	flags.BoolVar(&c.Debug, "debug", false, "If debug mode is enabled")
	flags.StringVar(&c.Listener, "listener", "0.0.0.0:7666", "Listener for HTTP Connections")
	flags.StringVar(&c.RootDir, "root-dir", "/var/shortoftime", "Root location for Shortoftime")

	err := flags.Parse(a[1:])
	if err != nil {
		log.Fatalf("Unable to parse command options: [%s]", err)
	}

	return c
}

func serve(c config.Conf, override string) error {
	m := gin.ReleaseMode
	if override != "" {
		m = override
	} else if c.IsDebug() {
		m = gin.DebugMode
	}

	s := server.Init(c, m)

	return s.Serve()
}
