package domain

type Wallet struct {
	Transactions []Transaction
}

func (w Wallet) filterByAsset(asset string) ([]Transaction) {
	filtered := []Transaction{}

	for _, transaction := range w.Transactions {
		if transaction.Asset == asset {
			filtered = append(filtered, transaction)
		}
	}

	return filtered
}
