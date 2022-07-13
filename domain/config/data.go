package config

import (
	"Vico1993/Wallet/domain/wallet"
	"log"

	"github.com/spf13/viper"
)

var v = viper.GetViper()

// Save Array of Operations with what we have in data.json
func SaveOperations(operations ...wallet.Operation) {
	// For now just erase what we have.
	v.Set("operations", operations)

	err := v.WriteConfig()
	if err != nil {
		log.Fatalln("Error saving operations: ", err.Error())
	}
}

// Get all operations from the data.json
func LoadOperations() []wallet.Operation {
	var operations []wallet.Operation

	err := v.UnmarshalKey("operations", &operations)
	if err != nil {
		log.Fatalln("Error loading operations: ", err.Error())
	}

	return operations
}