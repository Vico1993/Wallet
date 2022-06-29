package cryptocom

import (
	"Vico1993/Wallet/domain/wallet"
	"Vico1993/Wallet/util"
	"log"
	"strconv"

	"github.com/mitchellh/hashstructure/v2"
)

// TODO: Clean this file because its TOO MESSY

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

func NewCryptoCom(path string) CryptoCom {
	return CryptoCom{
		CsvPath: path,
	}
}

func (c CryptoCom) Load() (wallet.Wallet, error) {
	csvData, err := c.readCryptoComCSV()
	if err != nil {
		return wallet.Wallet{}, err
	}

	config := newCryptoComConfig()
	var operations []wallet.Operation
	for _, d := range util.ReverseSlice(csvData) {
		rowHash, err := hashstructure.Hash(d, hashstructure.FormatV2, nil)
		if err != nil {
			log.Println("Couldn't generate hash for CSV row: ", err.Error())
			log.Println(d)
		}

		// if already in the json, continue
		if util.IsInStringSlice(strconv.FormatUint(rowHash, 10), config.Operations_hash) {
			continue
		} else {
			config.addHash(strconv.FormatUint(rowHash, 10))
		}

		operations = append(
			operations,
			*buildOperations(d),
		)
	}

	config.save()

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