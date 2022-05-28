package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

var configFileName = "/config.json"

func InitConfig () {
	homedir, err := os.UserHomeDir()
	if err != nil {
		homedir = "./"
	}

	path := homedir + "/.wallet"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			log.Fatal("Can't create Folder at " + path)
		}
	}

	configFilePath := path + configFileName

	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		var file, err = os.Create(configFilePath)
        if err != nil {
            log.Fatal("Can't create config file at " + configFilePath)
        }
        defer file.Close()
	}

	viper.SetConfigFile(configFilePath)

	// err = viper.ReadInConfig()
	// if err != nil {
	// 	log.Fatal("Fatal error config file: %w \n", err)
	// }
}