package cmd

import (
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

	flags := &flags{}

	rootCmd.AddCommand(addCommand(flags))
	rootCmd.AddCommand(listCommand(flags))

	rootCmd.SetOut(stdout)
	rootCmd.SetErr(stderr)

	return rootCmd
}