package main

import (
	"Vico1993/Wallet/builder"
	"Vico1993/Wallet/domain"
	"Vico1993/Wallet/service"
	"Vico1993/Wallet/util"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/glamour"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var stats = domain.Statistic{}
var today = time.Now().Local().Format("2006-01-02 15:04:05")
var v = viper.GetViper()

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

	data := []builder.MasterBuilder{
		{
			String: builder.NewMarkDowText("Wallet", "h1"),
		},
		{
			String: builder.NewMarkDowText("At %s", "h2"),
			Data: util.TransformStringSliceIntoInterface([]string{today}),
		},
		{
			String: "We found %s number of transaction in your wallet, here is a small Summary: \n",
			Data: util.TransformStringSliceIntoInterface([]string{strconv.Itoa(len(wallet.Transactions))}),
		},
		{
			String: builder.NewMarkDowTable(
				[]string{"Asset", "Quantity", "By at", "By for (CAD)", "Price today", "Profit"},
				transactions,
			),
		},
		{
			String: builder.NewMarkDowText("You invest in total: %s and your total profit is: %s%%", "h3"),
			Data: util.TransformStringSliceIntoInterface([]string{
				util.FormatFloat(stats.GetTotalInvest()),
				util.FormatFloat(stats.GetTotalProfit()),
			}),
		},
		{
			String: builder.NewMarkDowText("Top Crypto", "h1"),
		},
		{
			String: builder.NewMarkDowTable(
				[]string{"Symbol", "Profit", "Quantity"},
				details,
			),
		},
	}

	render := ""
	// TODO: Clean main.go - find a better way to manage []]builder.MasterBuilder...
	for _, element := range data {
		var renderStr = ""

		if s, ok := element.String.(string); ok {
			renderStr = s
		}

		if s, ok := element.String.(builder.MarkDownBuilder); ok {
			renderStr, err = s.Render()
			if err != nil {
				log.Fatalln("Error Building MarkDownBuilder: ", err.Error())
			}
		}

		paramMatch := regexp.MustCompile("%s")
		fmt.Println(renderStr, paramMatch.MatchString(renderStr))
		if paramMatch.FindString("%s") != "" {
			render += fmt.Sprintf(
				renderStr,
				element.Data...,
			)
		} else {
			render += fmt.Sprint(
				renderStr,
			)
		}
	}

	r, _ := glamour.NewTermRenderer(
		// detect background color and pick either the default dark or light theme
		glamour.WithAutoStyle(),
		// wrap output at specific width
		glamour.WithWordWrap(100),
	)

	out, err := r.Render(
		strings.ReplaceAll(render, "\t", ""),
	)

	if err != nil {
		log.Println(err)
	}
	fmt.Print(out)
}