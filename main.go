package main

import (
	"fmt"
	"strconv"
)

func main() {
	wallet := GetData()

	fmt.Printf("Number of Transactions: %s\n", strconv.Itoa(len(wallet.Transactions)))

	for _, transaction := range wallet.Transactions {
		fmt.Printf("%s: You bough %s, at %s\n", transaction.Date, transaction.Asset, strconv.FormatFloat(transaction.AssetPrice, 'g', -1, 64))
	}
}