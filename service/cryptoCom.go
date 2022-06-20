package service

import (
	"Vico1993/Wallet/domain"
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
	ToAmount string
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

func (c CryptoCom) Load() (domain.Wallet, error) {
	var wallet domain.Wallet
	csvData, err := c.readCryptoComCSV()
	if err != nil {
		return domain.Wallet{}, err
	}

	for _, d := range util.ReverseSlice(csvData) {
		var tpe string

		switch d.TransactionKind {
			case CRYPTO_PURCHASE:
				tpe = domain.PURCHASE
			case CRYPTO_EXCHANGE:
				tpe = domain.EXCHANGE
			case CRYPTO_EARN:
				tpe = domain.EARN
		}

		wallet.AddOperation(
			domain.NewOperation(
				d.NativeAmount,
				d.Timestamp,
				d.Amount,
				d.Currency,
				0.0,
				tpe,
				"crypto.com",
				tpe,
			),
		)
	}

	return wallet, nil
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
			ToAmount: row[5],
			NativeCurrency: row[6],
			NativeAmount: util.TransformStringToFloat(row[7]),
			NativeAmountUSD: util.TransformStringToFloat(row[8]),
			TransactionKind: row[9],
			TransactionHash: row[10],
		})
	}

	return data, nil
}