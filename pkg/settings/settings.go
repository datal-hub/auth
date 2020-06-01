package settings

import (
	"fmt"
	"log"

	"github.com/go-ini/ini"
)

//DBType is the structure which define parameters for database connection
type DBType struct {
	Url string
}

var (
	//ListenAddr is the address on which the server is running
	ListenAddr string
	// VerboseMode define logging mode
	// True - for verbose mode
	// False - for common mode
	VerboseMode bool
	//DB define parameters for database connection
	DB DBType
)

func init() {
	ListenAddr = "localhost:8181"
	VerboseMode = false
	DB.Url = "postgresql://xxx:xxxd@127.0.0.1:5432/auth?sslmode=disable"
}

//FromFile initialize settings from file
func FromFile(file string) error {
	log.Printf("Loading configuration from '%s'", file)

	cfg, err := ini.InsensitiveLoad(file)
	if err != nil {
		return fmt.Errorf("Error parsing file '%s'. Message: %s", file, err)
	}

	newListenAddr := cfg.Section("").Key("ListenAddr").String()

	var newDB DBType
	newDB.Url = cfg.Section("DB").Key("Url").String()

	ListenAddr = newListenAddr
	DB = newDB

	return nil
}
