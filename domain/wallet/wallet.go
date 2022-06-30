package wallet

import (
	"errors"
	"strings"

	"github.com/spf13/viper"
)

var v = viper.GetViper()

type unitDetail struct {
	symbol  string
	// profit  float64 -- TODO: Implement Profit
	quantity float64
}

type Wallet struct {
	tag 				[]string
	units 				map[string]unitDetail

	operations 			[]Operation
	operationsProfit    [][]string
}

func NewWallet(o []Operation, t ...string) Wallet {
	wallet := Wallet{
		operations: o,
		tag: t,
		units: make(map[string]unitDetail),
	}

	for _, operation := range o {
		// Handle each operation one by one
		wallet.handleOperation(operation)
	}

	return wallet
}

func (w *Wallet) handleOperation(o Operation) {
	// In case NewWallet is not use
	if w.units == nil {
		w.units = make(map[string]unitDetail)
	}

	unit := strings.ToLower(o.Unit)
	from := strings.ToLower(o.From)

	if entry, ok := w.units[unit]; ok {
		entry.quantity += o.Quantity

		w.units[unit] = entry
	} else {
		w.units[unit] = unitDetail{
			quantity: o.Quantity,
			symbol: unit,
		}
	}

	if o.OType == EXCHANGE {
		w.units[from] = unitDetail{
			quantity: w.units[from].quantity - o.FromQuantity,
			symbol: w.units[from].symbol,
		}
	}

	// Add operations to our Profits constants
	w.operationsProfit = append(w.operationsProfit, o.WithProfit())
}

func (w *Wallet) AddOperation(ope Operation) {
	w.operations = append(w.operations, ope)

	// Update the unit key
	w.handleOperation(ope)
}

func (w Wallet) GetQuantityByUnit(unit string) (float64, error) {
	if entry, ok := w.units[strings.ToLower(unit)]; ok {
		return entry.quantity, nil
	} else {
		return 0, errors.New("Unit not found in the wallet")
	}
}

func (w Wallet) GetOperations() []Operation {
	return w.operations
}

func (w Wallet) GetProfitTable() ([]string, [][]string) {
	return []string{
		"Unit", "Quantity", "Buy at", "Buy for (CAD)", "Current Price", "Profit",
	}, w.operationsProfit
}

func (w Wallet) Save() error {
	// if no operations to save, do nothing
	if len(w.operations) == 0 {
		return nil
	}

	var oldOperations []Operation

	err := v.UnmarshalKey("operations", &oldOperations)
	if err != nil {
		return errors.New("Error saving operations: " + err.Error())
	}

	v.Set(
		"operations",
		append(
			w.operations,
			oldOperations...
		),
	)

	err = v.WriteConfig()
	if err != nil {
		return errors.New("Error saving operations: " + err.Error())
	}

	return nil
}