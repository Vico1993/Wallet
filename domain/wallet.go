package domain

type Wallet struct {
	transactions []Transaction
}

func (w Wallet) filterByAsset(asset string) ([]Transaction) {
	filtered := []Transaction{}

	for _, transaction := range w.transactions {
		if transaction.asset == asset {
			filtered = append(filtered, transaction)
		}
	}

	return filtered
}
