package cryptocom

import (
	"Vico1993/Wallet/domain/wallet"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildOperationEarnTransaction(t *testing.T) {
	result := buildOperations(cryptoCSV{
		Timestamp: "2022-01-01",
		Description: "Super Buy",
		Currency: "BTC",
		Amount: 100,
		ToCurrency: "",
		ToAmount: 0,
		NativeCurrency: "CAD",
		NativeAmount: 100,
		NativeAmountUSD: 90,
		TransactionKind: CRYPTO_EARN,
		TransactionHash: "123141414810418904182908409280810",
	})

	assert.Equal(
		t,
		&wallet.Operation{
			Date: "2022-01-01",
			Quantity: 100,
			Unit: "BTC",
			UnitPrice: 1,
			From: "CAD",
			Price: 100,
			FromQuantity: 100,
			Fiat: "CAD",
			OType: wallet.EARN,
			Tag: []string{"crypto.com", wallet.EARN},
		},
		result,
	)
}

func TestBuildOperationPurchaseTransaction(t *testing.T) {
	result := buildOperations(cryptoCSV{
		Timestamp: "2022-01-01",
		Description: "Super Buy",
		Currency: "BTC",
		Amount: 100,
		ToCurrency: "",
		ToAmount: 0,
		NativeCurrency: "CAD",
		NativeAmount: 100,
		NativeAmountUSD: 90,
		TransactionKind: CRYPTO_PURCHASE,
		TransactionHash: "123141414810418904182908409280810",
	})

	assert.Equal(
		t,
		&wallet.Operation{
			Date: "2022-01-01",
			Quantity: 100,
			Unit: "BTC",
			UnitPrice: 1,
			From: "CAD",
			Price: 100,
			FromQuantity: 100,
			Fiat: "CAD",
			OType: wallet.PURCHASE,
			Tag: []string{"crypto.com", wallet.PURCHASE},
		},
		result,
	)
}

func TestBuildOperationEventNotSupported(t *testing.T) {
	result := buildOperations(cryptoCSV{
		Timestamp: "2022-01-01",
		Description: "Super Buy",
		Currency: "BTC",
		Amount: 100,
		ToCurrency: "",
		ToAmount: 0,
		NativeCurrency: "CAD",
		NativeAmount: 100,
		NativeAmountUSD: 90,
		TransactionKind: "event_payment_crypto",
		TransactionHash: "123141414810418904182908409280810",
	})

	assert.Nil(t, result)
}