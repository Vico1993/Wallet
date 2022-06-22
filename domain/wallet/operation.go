package wallet

var (
	PURCHASE = "purchase"
	EXCHANGE = "exchange"
	EARN = "earn"
)

type Operation struct {
	Date 			string 		`json:"date"`
	Quantity 		float64 	`json:"quantity"`
	Unit 			string 		`json:"unit"`
	UnitPrice		float64 	`json:"unitPrice"`
	From 			string  	`json:"from"`
	FromQuantity	float64		`json:"fromQuantity"`
	Price 			float64		`json:"price"`
	Fiat 			string		`json:"fiat"`
	OType			string		`json:"otype"`
	Tag 			[]string	`json:"tag"`
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
		Date: date,
		Quantity: quantity,
		Unit: unit,
		UnitPrice: unitPrice,
		From: from,
		FromQuantity: fromQuantity,
		Price: price,
		Fiat: fiat,
		OType: operationType,
		Tag: tag,
	}
}

func (o Operation) GetUnitPrice() float64 {
	return o.UnitPrice
}

