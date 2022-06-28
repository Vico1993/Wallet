package main

import (
	"Vico1993/Wallet/domain/cmd"
	"Vico1993/Wallet/domain/config"
)

func main() {
	config.InitConfig()

	cmd.Execute()
}