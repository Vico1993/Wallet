package service

import (
	"Vico1993/Wallet/util"
	"fmt"
	"log"
	"os"
)

const (
	CRYPTO_PURCHASE = "crypto_purchase"
	CRYPTO_EXCHANGE = "crypto_exchange"
)


type CSVStruct struct {
	Timestamp string
	Description string
	Currency string
	Amount float64
	ToCurrency string
	ToAmount string
	NativeCurrency string
	NativeAmount float64
	NativeAmountUSD float64
	TransactionKind string
	TransactionHash string
}

func InitCryptoCom() {
	data := loadInformations()

	for k, d := range util.ReverseSlice(data) {
		fmt.Println(d.Timestamp, k)
	}

	log.Fatalln("STOP DEV PROGRESS")
}

func loadInformations() []CSVStruct {
	var data []CSVStruct
	// TODO: Remove the CSV link from the ENV once the dev is complete
	rows, err := util.ReadCsv(os.Getenv("CSV_CRYPTO_COM"))
	if (err != nil ) {
		log.Fatalln("Error reading the CSV", err.Error())
	}

	for _, row := range rows {
		data = append(data, CSVStruct{
			Timestamp: row[0],
			Description: row[1],
			Currency: row[2],
			Amount: util.TransformStringToFloat(row[3]),
			ToCurrency: row[4],
			ToAmount: row[5],
			NativeCurrency: row[6],
			NativeAmount: util.TransformStringToFloat(row[7]),
			NativeAmountUSD: util.TransformStringToFloat(row[8]),
			TransactionKind: row[9],
			TransactionHash: row[10],
		})
	}

	return data
}