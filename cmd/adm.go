package cmd

import (
	"github.com/urfave/cli"

	"github.com/datal-hub/auth/pkg/database"
	log "github.com/datal-hub/auth/pkg/logger"
)

//description run the application from the command line in administrator mode
var Adm = cli.Command{
	Name:        "adm",
	Usage:       "Administrative tools for auth service",
	Description: ``,
	Subcommands: []cli.Command{
		{
			Name:   "initdb",
			Usage:  "Initialize database",
			Action: initDB,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "force, f",
					Usage: "Drop auth database!"},
			},
		},
	},
}

func initDB(c *cli.Context) error {
	db, err := database.NewDB()
	if err != nil {
		log.ErrorF("Error database initialization", log.Fields{"message": err.Error()})
		return err
	}
	defer func() {
		err := db.Close()
		if err != nil {
			log.Error(err.Error())
		}
	}()
	return db.Init(c.Bool("force"))
}
