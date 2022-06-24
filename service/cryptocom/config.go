package cryptocom

import (
	"log"

	"github.com/spf13/viper"
)

var v = viper.GetViper()

type cryptoComConfig struct {
	Operations_hash []string `json:"Operations_hash"`
}

func newCryptoComConfig() cryptoComConfig {
	var config cryptoComConfig

	err := v.UnmarshalKey("crypto_com_config", &config)
	if err != nil {
		log.Fatalln("Error loading the config:", err.Error())
	}

	return config
}

func (c *cryptoComConfig) addHash(h string) {
	c.Operations_hash = append(c.Operations_hash, h)
}

func (c cryptoComConfig) save() {
	v.Set("crypto_com_config", c)

	err := v.WriteConfig()
	if err != nil {
		log.Fatalln("Error saving Hash:", err.Error())
	}
}