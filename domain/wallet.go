package domain

import (
	"errors"
	"strings"
)

type unitDetail struct {
	symbol  string
	// profit  float64 -- TODO: Implement Profit
	quantity float64
}

type Wallet struct {
	// To not break everything
	Transactions 	[]Transaction

	operations 		[]Operation
	tag 			[]string
	units 			map[string]unitDetail
}

func NewWallet(o []Operation, t ...string) Wallet {
	units := make(map[string]unitDetail)

	// Kind of Ugly
	for _, operation := range o {
		unit := strings.ToLower(operation.unit)

		switch operation.oType {
			case PURCHASE:
				if entry, ok := units[unit]; ok {
					entry.quantity += operation.quantity

					units[unit] = entry
				} else {
					units[unit] = unitDetail{
						quantity: operation.quantity,
						symbol: unit,
					}
				}
				break
			case EXCHANGE:
				// Remove from old
				if entry, ok := units[strings.ToLower(operation.from)]; ok {
					entry.quantity -= operation.fromQuantity

					units[strings.ToLower(operation.from)] = entry
				} else {
					// ???
				}

				// Add from other
				if entry, ok := units[unit]; ok {
					entry.quantity += operation.quantity

					units[unit] = entry
				} else {
					units[unit] = unitDetail{
						quantity: operation.quantity,
						symbol: unit,
					}
				}
				break
		}
	}

	return Wallet{
		operations: o,
		tag: t,
		units: units,
	}
}

// Will be use for statistic later
// Act upon the Operation type
func (w *Wallet) AddOperation(ope Operation) {
	w.operations = append(w.operations, ope)
}

func (w Wallet) GetQuantityByUnit(unit string) (float64, error) {
	if entry, ok := w.units[strings.ToLower(unit)]; ok {
		return entry.quantity, nil
	} else {
		return 0, errors.New("Unit not found in the wallet")
	}
}