package domain

var (
	PURCHASE = "purchase"
	EXCHANGE = "exchange"
	EARN = "earn"
)

type Operation struct {
	price 		float64
	date 		string
	quantity 	float64
	asset 		string
	assetPrice 	float64
	oType		string
	tag 		[]string
}

func NewOperation(
	price float64,
	date string,
	quantity float64,
	asset string,
	assetPrice float64,
	operationType string,
	tag ...string,
) Operation {
	if assetPrice == 0 {
		assetPrice = 1 * price / quantity
	}

	return Operation{
		price: price,
		date: date,
		quantity: quantity,
		asset: asset,
		assetPrice: assetPrice,
		oType: operationType,
		tag: tag,
	}
}

func (o Operation) GetAssetPrice() float64 {
	return o.assetPrice
}