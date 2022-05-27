package main

import (
	"Vico1993/Wallet/domain"
	"Vico1993/Wallet/service"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/glamour"
	"github.com/joho/godotenv"
)

var stats = domain.Statistic{}

func formatFloat(numb float64) string {
	return strconv.FormatFloat(numb, 'g', -1, 64)
}

func main() {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	wallet := GetData()
	today := time.Now()

	render := fmt.Sprintf(
		`# Wallet
		_At %s_ `,
		today.Local().Format("2006-01-02 15:04:05"),
	)

	render += fmt.Sprintf("We found %s number of transaction in your wallet, here is a small Summary: \n", strconv.Itoa(len(wallet.Transactions)))

	render += fmt.Sprintln(
		`| Asset |  Quantity  |  By at   | By for (CAD) | Price today | Profit |
		| :---: | :--------: | :------: | :----: | :---------: | :----: |`,
	)

	for _, transaction := range wallet.Transactions {
		price, err := service.GetAssetPrice(transaction.Asset)
		if err != nil {
			log.Fatalln(transaction.Asset, "---" , err)
		}

		stats.AddInvest(
			transaction.Asset,
			transaction.Price,
			transaction.Quantity * price,
		)

		render += fmt.Sprintf(
			`| %s | %s | %s | %s | %s | %s%% |` + "\n",
			transaction.Asset,
			formatFloat(transaction.Quantity),
			formatFloat(transaction.GetAssetPrice()),
			formatFloat(transaction.Price),
			formatFloat(price),
			formatFloat(transaction.GetProfit(price)),
		)
	}

	render += fmt.Sprintf(
		"\n ## You invest in total: %s and your total profit is: %s%%",
		formatFloat(stats.GetTotalInvest()),
		formatFloat(stats.GetTotalProfit()),
	)

	render += fmt.Sprintln("\n# Top Crypto")

	render += fmt.Sprintln(
		`| Symbol | Profit |
		| :---: | :--------: |`,
	)

	for _, value := range stats.GetDetails() {
		render += fmt.Sprintf(
			`| %s | %s%% |` + "\n",
			value.Symbol,
			formatFloat(value.Profit),
		)
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