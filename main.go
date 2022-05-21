package main

import (
	"Vico1993/Wallet/service"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/glamour"
	"github.com/joho/godotenv"
)

func formatFloat(numb float64) string {
	return strconv.FormatFloat(numb, 'g', -1, 64)
}

func formatStringForRendering(str string) string {
	return strings.ReplaceAll(str, "\t", "")
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

	render += fmt.Sprintf("We found %s number of transaction in your wallet, here is a small Summary:", strconv.Itoa(len(wallet.Transactions)))

	render += "\n"
	render += fmt.Sprintf(
		`| Asset |  Quantity  |  By at   | By for | Price today | Profit |
		| :---: | :--------: | :------: | :----: | :---------: | :----: |`,
	)
	render += "\n"

	for _, transaction := range wallet.Transactions {
		price, err := service.GetAssetPrice(transaction.Asset)
		if err != nil {
			log.Fatalln(transaction.Asset, "---" , err)
		}

		render += fmt.Sprintf(
			`| %s | %s | %s | %s | %s | %s |`,
			transaction.Asset,
			formatFloat(transaction.Quantity),
			formatFloat(transaction.AssetPrice),
			formatFloat(transaction.Price),
			formatFloat(price),
			"0%",
		)
		render += "\n"
	}

	out, err := glamour.Render(
		strings.ReplaceAll(render, "\t", ""),
		"dark",
	)

	if err != nil {
		log.Println(err)
	}
	fmt.Print(out)
}