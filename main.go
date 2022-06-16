package main

import (
	"Vico1993/Wallet/builder"
	"Vico1993/Wallet/domain"
	"Vico1993/Wallet/service"
	"Vico1993/Wallet/util"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/guptarohit/asciigraph"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// TODO: Deal with error in a nice way...
var stats = domain.Statistic{}
var today = time.Now().Local().Format("2006-01-02 15:04:05")
var v = viper.GetViper()

// TOOD: Move these method into.. Something pretty.
func saveResults(p float64) {
	v.Set(
		"previous_result",
		append(
			v.GetStringSlice("previous_result"),
			today + " - " + util.FormatFloat(p) + "%",
		),
	)

	err := v.WriteConfig()
	if err != nil {
		log.Fatalln("Error saving profit")
	}
}

func getHistoricData() []float64 {
	historicFloat := []float64{}
	for _, row := range v.GetStringSlice("previous_result") {
		row = strings.Replace(row, "%", "", -1)

		data := strings.Split(row, " - ")
		if len(data) > 1 {
			profitValue, _ := strconv.ParseFloat(data[1], 64)
			historicFloat = append(historicFloat, profitValue)
		}
	}

	return historicFloat
}

func getFirstDateOfHistoric() string {
	firstDate := today
	rows := v.GetStringSlice("previous_result")

	if len(rows) > 0 {
		firstRow := rows[0]
		tmp := strings.Split(firstRow, " - ")

		firstDate = tmp[0]
	}

	return firstDate
}

func main() {
	// Create a home directory to save some basic information
	InitConfig()

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	wallet := GetData()
	var transactions [][]string
	for _, transaction := range wallet.Transactions {
		price, err := service.GetAssetPrice(transaction.Asset)
		if err != nil {
			log.Fatalln(transaction.Asset, "---" , err)
		}

		stats.AddInvest(
			transaction.Asset,
			transaction.Price,
			transaction.Quantity * price,
			transaction.Quantity,
		)

		transactions = append(transactions, []string{
			transaction.Asset,
			util.FormatFloat(transaction.Quantity),
			util.FormatFloat(transaction.GetAssetPrice()),
			util.FormatFloat(transaction.Price),
			util.FormatFloat(price),
			util.FormatFloat(transaction.GetProfit(price)),
		})
	}

	var details [][]string
	for _, detail := range stats.GetDetails() {
		details = append(details, []string{
			detail.Symbol,
			util.FormatFloat(detail.Profit),
			util.FormatFloat(detail.Quantity),
		})
	}

	// Save result for later
	if os.Getenv("DEBUG") == "0" {
		saveResults(stats.GetTotalProfit())
	}

	historicData := getHistoricData()
	markdown := builder.NewMarkDown([]builder.MarkDownData{
		{
			Block: builder.NewMarkDowText("Wallet", "h1", nil),
		},
		{
			Block: builder.NewMarkDowText("At %s", "h2", util.TransformStringSliceIntoInterface([]string{today})),
		},
		{
			Block: builder.NewMarkDowText("We found %s number of transaction in your wallet, here is a small Summary:", "text", util.TransformStringSliceIntoInterface([]string{
					strconv.Itoa(len(wallet.Transactions)),
				}),
			),
		},
		{
			Block: builder.NewMarkDowTable(
				[]string{"Asset", "Quantity", "By at", "By for (CAD)", "Price today", "Profit"},
				transactions,
			),
		},
		{
			Block: builder.NewMarkDowText("Resume by Crypto", "h1", nil),
		},
		{
			Block: builder.NewMarkDowTable(
				[]string{"Symbol", "Profit", "Quantity"},
				details,
			),
		},
		{
			Block: builder.NewMarkDowText("You invest in total: %s and your total profit is: %s%%", "h3", util.TransformStringSliceIntoInterface([]string{
					util.FormatFloat(stats.GetTotalInvest()),
					util.FormatFloat(stats.GetTotalProfit()),
				}),
			),
		},
		{
			Block: builder.NewMarkDowText("Historic of Profit", "h1", nil),
		},
		{
			Block: builder.NewGraph(
				[]asciigraph.Option{
					asciigraph.SeriesColors(
						asciigraph.Red,
						asciigraph.White,
					),
				},
				historicData,
				make([]float64, len(historicData)),
			),
		},
		{
			Block: builder.NewMarkDowText("From %s to %s", "h2", util.TransformStringSliceIntoInterface([]string{
					getFirstDateOfHistoric(),
					today,
				}),
			),
		},
	})

	err = markdown.Render()
	if err != nil {
		log.Fatalln("Error building the Markdown", err.Error())
	}
}