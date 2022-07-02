package wallet

import (
	"Vico1993/Wallet/service"
	"Vico1993/Wallet/util"
	"errors"
	"strings"

	"github.com/spf13/viper"
)

var v = viper.GetViper()

type unitDetail struct {
	symbol  	string
	invest 		float64
	profit  	float64
	quantity 	float64
}

func (u unitDetail) GetProfitRow() []string {
	return []string{
		u.symbol,
		util.FormatFloat(u.quantity),
		util.FormatFloat(u.invest),
		util.FormatFloat(u.profit),
	}
}

type Wallet struct {
	tag 				[]string
	units 				map[string]unitDetail

	operations 			[]Operation
	operationsProfit    [][]string

	TotalInvest			float64
	TotalValue			float64
}

func NewWallet(o []Operation, t ...string) Wallet {
	wallet := Wallet{
		operations: o,
		tag: t,
		units: make(map[string]unitDetail),
		TotalInvest: float64(0),
		TotalValue: float64(0),
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
		entry.invest += o.Price

		w.units[unit] = entry
	} else {
		w.units[unit] = unitDetail{
			quantity: o.Quantity,
			invest: o.Price,
			symbol: unit,
		}
	}

	if o.OType == EXCHANGE {
		w.units[from] = unitDetail{
			quantity: w.units[from].quantity - o.FromQuantity,
			invest: w.units[from].invest - o.Price,
			symbol: w.units[from].symbol,
		}
	}

	// Add operations to our Profits constants
	w.operationsProfit = append(w.operationsProfit, o.WithProfit())

	// Total
	currentUnitPrice := o.getCurrentUnitPrice()

	w.TotalInvest += o.Price
	w.TotalValue += o.GetCurrentPrice(currentUnitPrice)
}

func (w *Wallet) AddOperation(ope Operation) {
	w.operations = append(w.operations, ope)

	// Update the unit key
	w.handleOperation(ope)
}

func (w *Wallet) GetProfitByUnit() map[string]unitDetail {
	for unit, unitDetail := range w.units {
		currentValue, err := service.GetAssetPrice(unitDetail.symbol)
		if err != nil {
			continue
		}

		unitDetail.profit = calculProfit(unitDetail.invest, currentValue * unitDetail.quantity)
		w.units[unit] = unitDetail
	}

	return w.units
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

func (w Wallet) GetTotalProfit() string {
	return util.FormatFloat(calculProfit(w.TotalInvest, w.TotalValue))
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