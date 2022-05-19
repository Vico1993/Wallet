package main

import (
	"Vico1993/Wallet/domain"
	"encoding/json"
	"io/ioutil"
)

// For now the data will come from a json in the root of the project.
// Could be move later in small SQLite DB?
// Or even a JSON but setup by the CLI and not the user.

func GetData() (domain.Wallet) {
	file, _ := ioutil.ReadFile("data.json")
	transactions := []domain.Transaction{}

	_ = json.Unmarshal([]byte(file), &transactions)

	return domain.Wallet{
		Transactions: transactions,
	}
}