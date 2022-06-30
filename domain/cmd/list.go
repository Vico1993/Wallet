package cmd

import (
	"Vico1993/Wallet/domain/builder"
	"Vico1993/Wallet/domain/wallet"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var v = viper.GetViper()

func listCommand(flags *flags) *cobra.Command {
	listCmd := &cobra.Command{
		Use: "list",
		Short: "List all your operations",
		Long: "Will list all your operations in your wallet with some analytics",
		Run: func(cmd *cobra.Command, args []string) {
			w := loadWallet()

			if len(w.GetOperations()) == 0 {
				markdown.AddData(builder.Data{
					Block: builder.NewMarkDowText("We don't have any operations in your wallet. Please use the `add` command first. Feel free to use the `help` command if you need more information.", "italic", nil),
				})

				return;
			}

			markdown.AddData(builder.Data{
				Block: builder.NewMarkDowText("We found %d operations in your wallet, here is the data:", "text", []interface{}{len(w.GetOperations())}),
			})

			header, operations := w.GetProfitTable()
			markdown.AddData(builder.Data{
				Block: builder.NewMarkDowTable(
					header,
					operations,
				),
			})

			// markdown.AddData(builder.Data{
			// 	Block: builder.NewMarkDowText("You invest in total: %s and your total profit is: %s%%", "h3", util.TransformStringSliceIntoInterface([]string{
			// 			util.FormatFloat(stats.GetTotalInvest()),
			// 			util.FormatFloat(stats.GetTotalProfit()),
			// 		}),
			// 	),
			// })

			err := markdown.Render()
			if err != nil {
				log.Fatalln("Error Rendering the add command", err.Error())
			}
		},
	}

	listCmd.Flags().BoolVarP(&flags.listGraph, "display-graph", "g", false, "Display the graphique at the end")

	return listCmd
}

func loadWallet() wallet.Wallet {
	var operations []wallet.Operation

	err := v.UnmarshalKey("operations", &operations)
	if err != nil {
		log.Fatalln("Error loading operations: ", err.Error())
	}

	return wallet.NewWallet(operations)
}