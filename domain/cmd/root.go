package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

func InitRootCommand(stdout, stderr io.Writer) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:           "wallet",
		Short:         "Wallet",
		Long:          "Small cli tool to help see your crypto portfolio",
		Version:       "1.0",
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	fmt.Println("In Root Command")

	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	return rootCmd
}