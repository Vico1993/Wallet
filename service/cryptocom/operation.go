package cryptocom

import (
	"Vico1993/Wallet/domain/wallet"
)

func buildOperations(row cryptoCSV) *wallet.Operation {
	// Hard code for now
	var tpe string
	from := "CAD"
	unit := row.Currency
	quantity := row.Amount
	fromQuantity := row.NativeAmount

	switch row.TransactionKind {
		case CRYPTO_EARN:
			tpe = wallet.EARN
		case CRYPTO_PURCHASE:
			tpe = wallet.PURCHASE
		case CRYPTO_EXCHANGE:
			unit = row.ToCurrency
			from = row.Currency
			fromQuantity = row.Amount
			quantity = row.ToAmount
			tpe = wallet.EXCHANGE
		default:
			// If not supported for the moment, skip
			return nil
	}

	ope := wallet.NewOperation(
		row.Timestamp,
		quantity,
		unit,
		0,
		from,
		fromQuantity,
		row.NativeAmount,
		"CAD",
		tpe,
		"crypto.com",
		tpe,
	)

	return &ope
}