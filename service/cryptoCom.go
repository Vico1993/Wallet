package service

import (
	"Vico1993/Wallet/domain/wallet"
	"Vico1993/Wallet/util"
	"os"
)

const (
	CRYPTO_PURCHASE = "crypto_purchase"
	CRYPTO_EXCHANGE = "crypto_exchange"
	CRYPTO_EARN = "crypto_earn_interest_paid"
)

type cryptoCSV struct {
	Timestamp string
	Description string
	Currency string
	Amount float64
	ToCurrency string
	ToAmount float64
	NativeCurrency string
	NativeAmount float64
	NativeAmountUSD float64
	TransactionKind string
	TransactionHash string
}

type CryptoCom struct {
	CsvPath string
}

func NewCryptoCom() Exchange {
	return CryptoCom{
		CsvPath: os.Getenv("CSV_CRYPTO_COM"),
	}
}

func (c CryptoCom) Load() (wallet.Wallet, error) {
	csvData, err := c.readCryptoComCSV()
	if err != nil {
		return wallet.Wallet{}, err
	}

	var operations []wallet.Operation
	for _, d := range util.ReverseSlice(csvData) {
		var tpe string

		// Hard code for now
		from := "CAD"
		unit := d.Currency
		quantity := d.Amount
		fromQuantity := d.NativeAmount

		switch d.TransactionKind {
			case CRYPTO_EARN:
				tpe = wallet.EARN
			case CRYPTO_PURCHASE:
				tpe = wallet.PURCHASE
			case CRYPTO_EXCHANGE:
				unit = d.ToCurrency
				from = d.Currency
				fromQuantity = d.Amount
				quantity = d.ToAmount
				tpe = wallet.EXCHANGE
			default:
				// If not supported for the moment, skip
				continue;
		}

		operations = append(
			operations,
			wallet.NewOperation(
				d.Timestamp,
				quantity,
				unit,
				0,
				from,
				fromQuantity,
				d.NativeAmount,
				"CAD",
				tpe,
				"crypto.com",
				tpe,
			),
		)
	}

	return wallet.NewWallet(operations, "Crypto.com"), nil
}

func (c CryptoCom) readCryptoComCSV() ([]cryptoCSV, error) {
	var data []cryptoCSV

	// TODO: Remove the CSV link from the ENV once the dev is complete
	rows, err := util.ReadCsv(c.CsvPath)
	if (err != nil ) {
		return nil, err
	}

	for _, row := range rows {
		data = append(data, cryptoCSV{
			Timestamp: row[0],
			Description: row[1],
			Currency: row[2],
			Amount: util.TransformStringToFloat(row[3]),
			ToCurrency: row[4],
			ToAmount: util.TransformStringToFloat(row[5]),
			NativeCurrency: row[6],
			NativeAmount: util.TransformStringToFloat(row[7]),
			NativeAmountUSD: util.TransformStringToFloat(row[8]),
			TransactionKind: row[9],
			TransactionHash: row[10],
		})
	}

	return data, nil
}