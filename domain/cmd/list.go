package cmd

import (
	"Vico1993/Wallet/domain/builder"
	"Vico1993/Wallet/domain/wallet"
	"Vico1993/Wallet/util"
	"log"
	"strings"

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
			w := loadWallet(strings.ToUpper(flags.filterByUnit))

			if len(w.GetOperations()) == 0 {
				markdown.AddData(builder.Data{
					Block: builder.NewMarkDowText("We don't have any operations in your wallet. Please use the `add` command first. Feel free to use the `help` command if you need more information.", "italic", nil),
				})

				return;
			}

			markdown.AddData(builder.Data{
				Block: builder.NewMarkDowText("We found %d operations in your wallet, here is the data:", "text", []interface{}{len(w.GetOperations())}),
			})

			header, operations := wallet.GetOperationsProfitTableFromWallet(w)
			markdown.AddData(builder.Data{
				Block: builder.NewMarkDowTable(
					header,
					operations,
				),
			})

			if flags.listByCrypto {
				header, rows := wallet.GetUnitProfitTableFromWallet(w)
				markdown.AddData(builder.Data{
					Block: builder.NewMarkDowTable(
						header,
						rows,
					),
				})
			}

			markdown.AddData(builder.Data{
				Block: builder.NewMarkDowText("You invest in total: %s and your total profit is: %s%%", "h3", util.TransformStringSliceIntoInterface([]string{
						util.FormatFloat(w.TotalInvest),
						w.GetTotalProfit(),
					}),
				),
			})

			err := markdown.Render()
			if err != nil {
				log.Fatalln("Error Rendering the add command", err.Error())
			}
		},
	}

	listCmd.Flags().BoolVarP(&flags.listGraph, "display-graph", "g", false, "Display the graphique at the end")
	listCmd.Flags().BoolVarP(&flags.listByCrypto, "by-crypto", "c", false, "Display the profit by crypto")
	listCmd.Flags().StringVarP(&flags.filterByUnit, "filter-by-unit", "u", "", "Filter all the list with one unit")

	return listCmd
}

func loadWallet(unit string) wallet.Wallet {
	var data []wallet.Operation
	var operations []wallet.Operation

	err := v.UnmarshalKey("operations", &data)
	if err != nil {
		log.Fatalln("Error loading operations: ", err.Error())
	}

	// Filtering for one UNIT
	if unit != "" {
		for _, operation := range data {
			if operation.Unit == unit || operation.From == unit {
				operations = append(operations, operation)
			}
		}
	} else {
		operations = data
	}

	return wallet.NewWallet(operations)
}