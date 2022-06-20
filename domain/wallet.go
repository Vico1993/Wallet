package domain

type Wallet struct {
	// To not break everything
	Transactions 	[]Transaction

	operations 		[]Operation
	tag 			[]string
}

func NewWallet(o []Operation, t []string) Wallet {
	return Wallet{
		operations: o,
		tag: t,
	}
}

// Will be use for statistic later
// Act upon the Operation type
func (w *Wallet) AddOperation(ope Operation) {
	w.operations = append(w.operations, ope)
}

func (w Wallet) GetOperations() []Operation {
	return w.operations
}

// func (w Wallet) filterByAsset(asset string) ([]Transaction) {
// 	filtered := []Transaction{}

// 	for _, transaction := range w.Transactions {
// 		if transaction.Asset == asset {
// 			filtered = append(filtered, transaction)
// 		}
// 	}

// 	return filtered
// }
