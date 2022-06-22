package main

import (
	"Vico1993/Wallet/domain/config"
	"Vico1993/Wallet/service/cryptocom"
	"Vico1993/Wallet/util"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// TODO: Deal with error in a nice way...
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
	config.InitConfig()

	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Crypto.com load
	crypto := cryptocom.NewCryptoCom()
	cWallet, err := crypto.Load()
	if err != nil {
		log.Fatalln("Error with Crypto.com", err.Error())
	}

	fmt.Println("BEFORE")

	for _, r := range cWallet.GetOperations() {
		fmt.Println(r)
	}

	config.SaveOperations(cWallet.GetOperations()...)

	fmt.Println("AFTER")

	for _, ope := range config.LoadOperations() {
		fmt.Println(ope)
	}

	// fmt.Println(cWallet.GetQuantityByUnit("BTC"))
}