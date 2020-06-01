package main

import (
	"os"

	"github.com/urfave/cli"

	"github.com/datal-hub/auth/cmd"
	log "github.com/datal-hub/auth/pkg/logger"
	"github.com/datal-hub/auth/pkg/settings"
)

func initApp(c *cli.Context) error {
	if c.IsSet("verbose") {
		settings.VerboseMode = true
	}

	defaultCfg := "/etc/auth/auth.conf"
	if _, err := os.Stat(defaultCfg); err == nil {
		err := settings.FromFile(defaultCfg)
		if err != nil {
			logDetails := log.Fields{
				"message":    err.Error(),
				"configPath": defaultCfg,
			}
			log.ErrorF("initApp: get settings from file error.", logDetails)
		}
	}

	if c.IsSet("config") {
		if len(c.String("config")) != 0 {
			err := settings.FromFile(c.String("config"))
			if err != nil {
				logDetails := log.Fields{
					"message":    err.Error(),
					"configPath": c.String("config"),
				}
				log.ErrorF("initApp: get settings from file error.", logDetails)
			}
		}
	}

	log.Init()
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "auth"
	app.Usage = "Auth service"
	app.Commands = []cli.Command{
		cmd.Srv,
	}
	app.Flags = append(app.Flags, cli.StringFlag{
		Name: "config, c", Value: "/etc/auth/auth.conf",
		Usage: "Load configuration `FILE`",
	})
	app.Flags = append(app.Flags, cli.BoolFlag{
		Name:  "verbose, vv",
		Usage: "Enable verbose mode",
	})
	app.Before = initApp
	app.Flags = append(app.Flags, []cli.Flag{}...)
	app.Run(os.Args)
}
