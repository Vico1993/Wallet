package domain

var (
	PURCHASE = "purchase"
	EXCHANGE = "exchange"
	EARN = "earn"
)

type Operation struct {
	date 			string
	quantity 		float64
	unit 			string
	unitPrice		float64
	from 			string
	fromQuantity	float64
	price 			float64
	fiat 			string
	oType			string
	tag 			[]string
}

func NewOperation(
	date string,
	quantity float64,
	unit string,
	unitPrice float64,
	from string,
	fromQuantity float64,
	price float64,
	fiat string,
	operationType string,
	tag ...string,
) Operation {
	if unitPrice == 0 {
		unitPrice = 1 * price / quantity
	}

	return Operation{
		date: date,
		quantity: quantity,
		unit: unit,
		unitPrice: unitPrice,
		from: from,
		fromQuantity: fromQuantity,
		price: price,
		fiat: fiat,
		oType: operationType,
		tag: tag,
	}
}

func (o Operation) GetUnitPrice() float64 {
	return o.unitPrice
}

