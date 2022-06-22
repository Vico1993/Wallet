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
	homedir, err := os.UserHomeDir()
	if err != nil {
		homedir = "./"
	}

	path := homedir + "/.wallet"

	// Check if .wallet folder exist
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			log.Fatal("Can't create Folder at " + path)
		}
	}

	configFilePath := path +  "/data.json"

	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		var file, err = os.Create(configFilePath)
        if err != nil {
            log.Fatal("Can't create config file at " + configFilePath)
        }
        defer file.Close()

		// initialisation of the JSON string
		_, err = file.WriteString("{}")
		if err != nil {
            log.Fatal("Can't initiate JSON config file " + err.Error())
        }
	}

	viper.SetConfigFile(configFilePath)

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal("Fatal error config file: %w \n", err)
	}
}