package cmd

import (
	"Vico1993/Wallet/domain/builder"
	"os"
)

var markdown builder.MarkDown
func Execute() {
	rootCmd := InitRootCommand(os.Stdout, os.Stderr)

	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}