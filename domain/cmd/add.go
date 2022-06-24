package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func addCommand(flags *flags) *cobra.Command {
	addCmd := &cobra.Command{
		Use: "add ",
		Short: "Add operations to your wallet",
		Long: "Add operations to your wallet base on your exchange. Use `-l` or `--list-exchange` to see the list of exchance suported",
		Run: func(cmd *cobra.Command, args []string) {
			if flags.addListExchange {
				fmt.Println("LIST")
			} else {
				fmt.Println("NOT LIST")
			}
		},
	}

	addCmd.Flags().BoolVarP(&flags.addListExchange, "list-exchange", "l", false, "display the list of exchange supported")

	return addCmd
}