package wallet

func GetOperationsProfitTableFromWallet(w Wallet) ([]string, [][]string) {
	return []string{
		"Unit", "Quantity", "Buy at", "Buy for (CAD)", "Current Price", "Profit",
	}, w.operationsProfit
}

func GetUnitProfitTableFromWallet(w Wallet) ([]string, [][]string) {
	var data [][]string
	for _, u := range w.GetProfitByUnit() {
		data = append(data, u.GetProfitRow())
	}

	return []string{"Unit", "Quantity", "Invest", "Profit"}, data
}