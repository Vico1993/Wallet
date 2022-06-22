package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

// For now the data will come from a json in the root of the project.
// Could be move later in small SQLite DB?
// Or even a JSON but setup by the CLI and not the user.

// Temp for now and see how it goes.

func InitConfig () {
	path := homedir()

	// Check if .wallet folder exist
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			log.Fatal("Can't create Folder at " + path)
		}
	}

	configFilePath := path +  "/data.json"

	initJsonFile(configFilePath)

	viper.SetConfigFile(configFilePath)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Fatal error config file: %w \n", err)
	}
}