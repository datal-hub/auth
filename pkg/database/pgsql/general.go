package pgsql

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	_ "github.com/lib/pq"

	log "github.com/datal-hub/auth/pkg/logger"
	"github.com/datal-hub/auth/pkg/settings"
)

// IsEmpty is method checking empties PostgreSQL
func (db *PgSQL) IsEmpty() bool {
	rows, err := db.DB.Query(`
		SELECT EXISTS (
			SELECT 1 FROM pg_tables
			WHERE  tablename IN ('credentials')
		);`)
	if err != nil {
		log.ErrorF("Error checking table existance.", log.Fields{"message": err.Error()})
		return false
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Error(err.Error())
		}
	}()

	if !rows.Next() {
		log.Error("Error fetching exists query result.")
		return false
	}

	var notEmpty bool
	if err := rows.Scan(&notEmpty); err != nil {
		log.Error(err.Error())
	}
	return !notEmpty
}

// Clear is method clearing PostgreSQL
func (db *PgSQL) Clear() {
	if _, err := db.DB.Exec("DROP TABLE IF EXISTS credentials"); err != nil {
		log.Error(err.Error())
	}
}

func (db *PgSQL) createTables() error {
	createSQL := `
		CREATE TABLE credentials (
			ID SERIAL PRIMARY KEY,
			login VARCHAR(256),
			phone VARCHAR(256),
			email VARCHAR(256),
			password VARCHAR(256),
			create_dttm timestamp DEFAULT CURRENT_TIMESTAMP
		);
		CREATE UNIQUE INDEX credentials_login_idx ON credentials(login);
	`

	if _, err := db.DB.Exec(createSQL); err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

func isValidDBName(name string) bool {
	if name != "" {
		if strings.HasSuffix(settings.DB.Url, "/"+name) {
			return true
		}
		if strings.Contains(settings.DB.Url, "/"+name+"?") {
			return true
		}
	}
	return false
}

// Init is method init database PostgreSQL
func (db *PgSQL) Init(force bool) error {
	logDetails := log.Fields{"db": settings.DB.Url}
	log.InfoF("Init: Initializing database.", logDetails)
	if !db.IsEmpty() {
		if force {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print(
				"Database is not empty. All data will be lost.\n",
				"Please confirm database initialization, type database name: ")
			dbname, _ := reader.ReadString('\n')
			dbname = strings.ReplaceAll(dbname, "\n", "")
			if !isValidDBName(dbname) {
				logDetails["dburl"] = settings.DB.Url
				log.ErrorF("Init: database names doesn't match.", logDetails)
				return errors.New("Database force initializing confirmation failed.")
			}
		} else {
			log.ErrorF("Init: database is not empty.", logDetails)
			return errors.New("Initialization failed - the database is not empty.")
		}
		db.Clear()
	}
	err := db.createTables()
	if err != nil {
		return err
	}
	log.Info("Database initialized successful.")
	return nil
}
