package config

import (
	"log"
	"os"
)

// Return homedire of wallet config project
func homedir() string {
	homedir, err := os.UserHomeDir()
	if err != nil {
		homedir = "./"
	}

	return homedir + "/.wallet"
}

// Create if not exist a json file
func initJsonFile(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		var file, err = os.Create(path)
        if err != nil {
            log.Fatal("Can't create config file at " + path)
        }
        defer file.Close()

		// initialisation of the JSON string
		_, err = file.WriteString("{}")
		if err != nil {
            log.Fatal("Can't initiate JSON config file " + err.Error())
        }
	}
}