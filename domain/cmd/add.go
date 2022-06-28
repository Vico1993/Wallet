package cmd

import (
	"Vico1993/Wallet/domain/builder"
	"Vico1993/Wallet/domain/wallet"
	"Vico1993/Wallet/service/cryptocom"
	"errors"
	"log"

	"github.com/spf13/cobra"
)

func addCommand(flags *flags) *cobra.Command {
	addCmd := &cobra.Command{
		Use: "add",
		Short: "Add operations to your wallet",
		Long: "Add operations to your wallet base on your exchange.",
		Run: func(cmd *cobra.Command, args []string) {
			var markdown builder.MarkDown

			if flags.addFilePath == "" {
				markdown.AddData(builder.Data{
					Block: builder.NewMarkDowText("We are missing the csv file path", "h1", nil),
				})
			} else {
				w, err := handleCryptoCom(flags.addFilePath)
				if err != nil {
					log.Fatalln("Error Rendering the add command", err.Error())
				}

				if len(w.GetOperations()) == 0 {
					markdown.AddData(builder.Data{
						Block: builder.NewMarkDowText("No new operation found","italic", nil),
					})
				} else {
					markdown.AddData(builder.Data{
						Block: builder.NewMarkDowText("We load and save %d","italic", []interface{}{len(w.GetOperations())}),
					})
				}
			}

			err := markdown.Render()
			if err != nil {
				log.Fatalln("Error Rendering the add command", err.Error())
			}
		},
	}

	// For now only support one exchange
	// addCmd.Flags().BoolVarP(&flags.addCryptoCom, "crypto-com", "c", false, "Add operations coming from Crypto.com")
	addCmd.Flags().StringVarP(&flags.addFilePath, "csv-file-path", "p", "", "Path to the CSV file you want to import")

	return addCmd
}

func handleCryptoCom(path string) (*wallet.Wallet, error) {
	service := cryptocom.NewCryptoCom(path)

	w, err := service.Load()
	if err != nil {
		return nil, errors.New("Error loading your csv file: " + err.Error())
	}

	err = w.Save()
	if err != nil {
		return nil, errors.New("Error loading your csv file: " + err.Error())
	}

	return &w, nil
}